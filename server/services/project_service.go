package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v69/github"
	"github.com/overal-x/formatio/models"
	"github.com/overal-x/formatio/types"
	"github.com/samber/do"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	GITHUB_DEPLOYMENT_QUEUE     = "GITHUB_DEPLOYMENT_QUEUE"
	GITHUB_DEPLOYMENT_LOG_QUEUE = "GITHUB_DEPLOYMENT_LOG_QUEUE"
)

type IProjectService interface {
	List(types.ListProjectArgs) ([]models.Project, error)
	Create(types.CreateProjectArgs) (*models.Project, error)
	Get(string) (*models.Project, error)
	Update(types.UpdateProjectArgs) (*models.Project, error)
	Delete(string) error
	Deploy(types.DeployArgs) error       // publisher
	HandleDeploy(types.DeployArgs) error // subscriber
	GetNework(string) (*models.Network, error)
}

type ProjectService struct {
	db *gorm.DB

	nixpacksService INixpacksService
	execService     IExecService
	githubServices  IGithubService
	rabbitmqService IRabbitMQService
}

func (p *ProjectService) List(args types.ListProjectArgs) (projects []models.Project, err error) {
	err = p.db.Find(&projects).Error
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (p *ProjectService) Create(args types.CreateProjectArgs) (*models.Project, error) {
	project := args.ToModel()
	err := p.db.Create(&project).Error
	if err != nil {
		return nil, err
	}

	environment := models.Environment{
		Name:      "Production",
		ProjectId: project.Id,
	}
	if len(args.Variables) > 0 {
		environment.Variables = args.Variables
	}

	err = p.db.Create(&environment).Error
	if err != nil {
		return nil, err
	}

	return &project, err
}

func (p *ProjectService) Get(id string) (*models.Project, error) {
	project := models.Project{}
	err := p.db.Where("id = ?", id).First(&project).Error
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (p *ProjectService) Update(args types.UpdateProjectArgs) (*models.Project, error) {
	project := args.ToModel()
	err := p.db.Where("id = ?", args.Id).Clauses(clause.Returning{}).Updates(&project).Error
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (p *ProjectService) Delete(id string) error {
	err := p.db.Where("project_id = ?", id).Delete(&models.Environment{}).Error
	if err != nil {
		return err
	}

	err = p.db.Where("id = ?", id).Delete(&models.Project{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectService) Deploy(args types.DeployArgs) error {
	payload, err := json.Marshal(&args)
	if err != nil {
		return err
	}

	return p.rabbitmqService.Publish(PublishArgs{
		Queue:   GITHUB_DEPLOYMENT_QUEUE,
		Content: string(payload),
	})
}

func (p *ProjectService) HandleDeploy(args types.DeployArgs) error {
	project := models.Project{}
	err := p.db.Where("id = ?", args.ProjectId).First(&project).Error
	if err != nil {
		return err
	}

	environment := models.Environment{}
	err = p.db.Where("project_id = ?", args.ProjectId).First(&environment).Error
	if err != nil {
		return err
	}

	app := models.GithubApp{}
	err = p.db.First(&app, "id = ?", project.AppId).Error
	if err != nil {
		return err
	}

	commitMessage := args.Message

	if commitMessage == "" {
		token, err := p.githubServices.GetInstallationToken(GetInstallationTokenArgs{
			ClientId:       app.ClientId,
			PrivateKey:     app.PrivateKey,
			InstallationId: int64(lo.Must(strconv.Atoi(project.InstallationId))),
		})
		if err != nil {
			return err
		}

		client := github.NewClient(nil).WithAuthToken(*token)
		commit, _, err := client.Repositories.GetCommit(
			context.Background(),
			strings.Split(project.RepoFullname, "/")[0],
			strings.Split(project.RepoFullname, "/")[1],
			args.CommitSha,
			nil,
		)
		if err != nil {
			return err
		}

		commitMessage = *commit.Commit.Message
	}

	deployment := models.Deployment{
		ProjectId:     args.ProjectId,
		EnvironmentId: environment.Id,
		Message:       commitMessage,
		Status:        models.DeploymentStatusPending,
	}
	err = p.db.Create(&deployment).Error
	if err != nil {
		return err
	}

	cloneUrl, err := p.githubServices.GetRepoCloneUrl(GetRepoCloneUrlArgs{
		RepoId:         project.RepoId,
		ClientId:       app.ClientId,
		PrivateKey:     app.PrivateKey,
		InstallationId: int64(lo.Must(strconv.Atoi(project.InstallationId))),
	})
	if err != nil {
		return err
	}

	project.Name = strings.ToLower(strings.ReplaceAll(project.Name, " ", "-"))
	projectDir := fmt.Sprintf("_tmp/%s-%s", project.Name, lo.RandomString(6, lo.LettersCharset))

	err = p.execService.Execute(ExecuteArgs{
		Command: fmt.Sprintf("git clone %s %s", *cloneUrl, projectDir),
		OutputCallback: func(s string) {
			fmt.Println("> ", s)
			deployment_log := models.DeploymentLog{
				DeploymentId: deployment.Id,
				Message:      s,
			}
			p.db.Create(&deployment_log)
		},
		ErrorCallback: func(s string) {
			fmt.Println("x ", s)
			deployment_log := models.DeploymentLog{
				DeploymentId: deployment.Id,
				Message:      s,
			}
			p.db.Create(&deployment_log)
		},
	})
	if err != nil {
		return err
	}
	defer os.RemoveAll(projectDir)

	if args.CommitSha != "" {
		err = p.execService.Execute(ExecuteArgs{
			Directory: projectDir,
			Command:   fmt.Sprintf("git checkout %s", args.CommitSha),
			OutputCallback: func(s string) {
				fmt.Println("> ", s)
				deployment_log := models.DeploymentLog{
					DeploymentId: deployment.Id,
					Message:      s,
				}
				p.db.Create(&deployment_log)
			},
			ErrorCallback: func(s string) {
				fmt.Println("x ", s)
				deployment_log := models.DeploymentLog{
					DeploymentId: deployment.Id,
					Message:      s,
				}
				p.db.Create(&deployment_log)
			},
		})
		if err != nil {
			return err
		}
	}

	err = p.nixpacksService.Build(BuildArgs{
		AppName:      project.Name,
		AppDirectory: projectDir,
		Env:          &map[string]string{},
		Callback: func(out *string, err error) {
			if err != nil {
				fmt.Println("x ", err)
			} else {
				fmt.Println("> ", *out)
			}

			deployment_log := models.DeploymentLog{
				DeploymentId: deployment.Id,
				Message:      *out,
			}
			p.db.Create(&deployment_log)
		},
	})
	if err != nil {
		return err
	}

	err = p.db.
		Model(&models.Deployment{}).
		Where("id = ?", deployment.Id).
		Update("status", models.DeploymentStatusSuccess).Error
	if err != nil {
		return err
	}

	err = p.execService.Execute(ExecuteArgs{
		Command: fmt.Sprintf("docker container rm -f %s", project.Name),
		OutputCallback: func(s string) {
			fmt.Println("> ", s)
		},
		ErrorCallback: func(s string) {
			fmt.Println("x ", s)
		},
	})
	if err != nil {
		return err
	}

	hostName := fmt.Sprintf("%s.localhost", project.Name)
	err = p.execService.Execute(ExecuteArgs{
		Command: fmt.Sprintf(
			`docker run -d -l "traefik.http.routers.%s.rule=Host(\"%s\")" -l "traefik.http.services.%s.loadbalancer.server.port=8000" --name %s %s`,
			project.Name, hostName, project.Name, project.Name, project.Name,
		),
		OutputCallback: func(s string) {
			fmt.Println("> ", s)
			network := models.Network{
				ProjectId: project.Id,
				ProcessId: s,
				HostName:  hostName,
			}
			err := p.db.Create(&network).Error
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				p.db.Updates(&network)
			}
		},
		ErrorCallback: func(s string) {
			fmt.Println("x ", s)
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *ProjectService) GetNework(id string) (network *models.Network, err error) {
	err = p.db.First(&network, "project_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return network, nil
}

func NewProjectService(i *do.Injector) (IProjectService, error) {
	db := do.MustInvoke[*gorm.DB](i)
	execService := do.MustInvoke[IExecService](i)
	nixpacksService := do.MustInvoke[INixpacksService](i)
	githubService := do.MustInvoke[IGithubService](i)
	rabbitmqService := do.MustInvoke[IRabbitMQService](i)

	return &ProjectService{
		db:              db,
		execService:     execService,
		nixpacksService: nixpacksService,
		githubServices:  githubService,
		rabbitmqService: rabbitmqService,
	}, nil
}
