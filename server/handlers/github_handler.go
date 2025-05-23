package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v69/github"
	"github.com/labstack/echo/v4"
	"github.com/overal-x/formatio/models"
	"github.com/overal-x/formatio/services"
	"github.com/overal-x/formatio/types"
	"github.com/overal-x/formatio/utils"
	"github.com/samber/do"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type IGithubHandler interface {
	CreateApp(c echo.Context) error
	ListApps(c echo.Context) error
	DeployRepo(c echo.Context) error
	ListInstallations(c echo.Context) error
	ListRepo(c echo.Context) error
}

type GithubHandler struct {
	db             *gorm.DB
	githubService  services.IGithubService
	projectService services.IProjectService
}

// @ID create-app
// @Success 201 {object} models.GithubApp
// @Router /api/github [get]
func (g *GithubHandler) CreateApp(c echo.Context) error {
	// authId := c.QueryParam("authId")
	next := c.QueryParam("next")
	code := c.QueryParam("code")

	client := github.NewClient(nil)
	ghApp, resp, err := client.Apps.CompleteAppManifest(context.Background(), code)
	if err != nil {
		return c.JSON(resp.StatusCode, echo.Map{"message": err})
	}

	app := models.GithubApp{
		AppId:         ghApp.GetID(),
		AppName:       ghApp.GetName(),
		OwnerId:       ghApp.GetOwner().GetID(),
		OwnerUsername: ghApp.GetOwner().GetLogin(),
		OwnerType:     ghApp.GetOwner().GetType(),
		ClientId:      ghApp.GetClientID(),
		ClientSecret:  ghApp.GetClientSecret(),
		WebhookSecret: ghApp.GetWebhookSecret(),
		PrivateKey:    ghApp.GetPEM(),
	}
	err = g.db.Create(&app).Error
	if err != nil {
		return utils.HandleGormError(c, err)
	}

	if next != "" {
		return c.Redirect(302, next)
	}

	return c.JSON(201, app)
}

// @ID list-apps
// @Success 200 {array} models.GithubApp
// @Router /api/github/apps [get]
func (g *GithubHandler) ListApps(c echo.Context) error {
	apps := []models.GithubApp{}
	err := g.db.Find(&apps).Error
	if err != nil {
		return utils.HandleGormError(c, err)
	}

	return c.JSON(200, apps)
}

func (g *GithubHandler) DeployRepo(c echo.Context) error {
	args := types.WebhhookPayload{}
	err := c.Bind(&args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	project := models.Project{}
	err = g.db.Find(&project,
		"installation_id = ? AND repo_id = ?",
		fmt.Sprintf("%d", args.Installation.ID), fmt.Sprintf("%d", args.Repository.ID),
	).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	err = g.projectService.Deploy(types.DeployArgs{
		ProjectId: project.Id,
		CommitSha: args.HeadCommit.ID,
		Message:   args.HeadCommit.Message,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	return c.JSON(200, nil)
}

// @ID list-installations
// @Success 200 {array} types.Installation
// @Param	app_id path string	true "App Id"
// @Router /api/github/installations/{app_id} [get]
func (g *GithubHandler) ListInstallations(c echo.Context) error {
	app := models.GithubApp{}
	err := g.db.First(&app).Where("id = ?", c.Param("app_id")).Error
	if err != nil {
		return utils.HandleGormError(c, err)
	}

	appToken, err := g.githubService.GetAppToken(services.GetAppTokenArgs{
		ClientId:   app.ClientId,
		PrivateKey: app.PrivateKey,
	})
	if err != nil {
		return err
	}

	client := github.NewClient(nil).WithAuthToken(*appToken)
	installations, _, err := client.Apps.ListInstallations(context.Background(), &github.ListOptions{})
	if err != nil {
		return err
	}

	return c.JSON(200, lo.Map(installations, func(installation *github.Installation, _ int) types.Installation {
		return types.Installation{
			Id:            installation.ID,
			AppId:         *installation.AppID,
			AppName:       app.AppName,
			AppSlug:       *installation.AppSlug,
			OwnerId:       *installation.Account.ID,
			OwnerUsername: *installation.Account.Login,
			OwnerType:     *installation.Account.Type,
			Events:        installation.Events,
		}
	}))
}

// @ID list-repo
// @Success 200 {array} types.Repo
// @Param	app_id path string	true "App Id"
// @Param	installation_id path string	true "Installation Id"
// @Router /api/github/repos/{app_id}/{installation_id} [get]
func (g *GithubHandler) ListRepo(c echo.Context) error {
	app := models.GithubApp{}
	err := g.db.First(&app).Where("id = ?", c.Param("app_id")).Error
	if err != nil {
		return utils.HandleGormError(c, err)
	}

	installationId := int64(lo.Must(strconv.Atoi(c.Param("installation_id"))))

	appToken, err := g.githubService.GetInstallationToken(services.GetInstallationTokenArgs{
		InstallationId: installationId,
		ClientId:       app.ClientId,
		PrivateKey:     app.PrivateKey,
	})
	if err != nil {
		return err
	}

	client := github.NewClient(nil).WithAuthToken(*appToken)
	repos, _, err := client.Apps.ListRepos(context.Background(), &github.ListOptions{})
	if err != nil {
		return err
	}

	return c.JSON(200, lo.Map(repos.Repositories, func(repo *github.Repository, _ int) types.Repo {
		return types.Repo{
			Id:            *repo.ID,
			Name:          *repo.Name,
			OwnerId:       *repo.Owner.ID,
			OwnerUsername: *repo.Owner.Login,
			OwnerType:     *repo.Owner.Type,
		}
	}))
}

func NewGithubHandler(i *do.Injector) (IGithubHandler, error) {
	db := do.MustInvoke[*gorm.DB](i)
	githubService := do.MustInvoke[services.IGithubService](i)
	projectService := do.MustInvoke[services.IProjectService](i)

	return &GithubHandler{db: db, githubService: githubService, projectService: projectService}, nil
}
