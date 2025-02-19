package types

import (
	"strconv"

	"github.com/overal-x/formatio/models"
	"github.com/samber/lo"
)

type ListProjectArgs struct{}

type CreateProjectArgs struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	AppId           string `json:"app_id"`
	InstallationId  string `json:"installation_id"`
	RepoId          string `json:"repo_id"`
	AutoDeploy      bool   `json:"auto_deploy"`
	RequireApproval bool   `json:"require_approval"`
	Variables       string `json:"variables"`
}

func (p *CreateProjectArgs) ToModel() models.Project {
	project := models.Project{
		Name:            p.Name,
		Description:     p.Description,
		AppId:           p.AppId,
		InstallationId:  p.InstallationId,
		AutoDeploy:      p.AutoDeploy,
		RequireApproval: p.RequireApproval,
	}
	if p.RepoId != "" {
		project.RepoId = int64(lo.Must(strconv.Atoi(p.RepoId)))
	}
	return project
}

type UpdateProjectArgs struct {
	Id              string `json:"id" param:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	AppId           string `json:"app_id"`
	InstallationId  string `json:"installation_id"`
	RepoId          string `json:"repo_id"`
	AutoDeploy      bool   `json:"auto_deploy"`
	RequireApproval bool   `json:"require_approval"`
}

func (p *UpdateProjectArgs) ToModel() models.Project {
	project := models.Project{
		Name:            p.Name,
		Description:     p.Description,
		AppId:           p.AppId,
		InstallationId:  p.InstallationId,
		AutoDeploy:      p.AutoDeploy,
		RequireApproval: p.RequireApproval,
	}

	if p.RepoId != "" {
		project.RepoId = int64(lo.Must(strconv.Atoi(p.RepoId)))
	}

	return project
}

type DeployArgs struct {
	ProjectId string `json:"id" param:"id"`
	CommitSha string `json:"commit_sha" body:"commit_sha"`
}
