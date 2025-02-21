package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/overal-x/formatio/services"
	"github.com/overal-x/formatio/types"
)

type IProjectHandler interface {
	List(c echo.Context) error
	Create(c echo.Context) error
	Get(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	Deploy(c echo.Context) error
	GetNetwork(c echo.Context) error
}

type ProjectHandler struct {
	projectService services.IProjectService
}

// @ID list-projects
// @Success 202 {array} models.Project
// @Router /api/projects [get]
func (p *ProjectHandler) List(c echo.Context) error {
	args := types.ListProjectArgs{}
	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	projects, err := p.projectService.List(args)
	if err != nil {
		// TODO: handle error to return proper status code
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, projects)
}

// @ID create-project
// @Success 201 {object} models.Project
// @Param project body types.CreateProjectArgs true "Project"
// @Router /api/projects [post]
func (p *ProjectHandler) Create(c echo.Context) error {
	args := types.CreateProjectArgs{}
	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	project, err := p.projectService.Create(args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, project)
}

// @ID get-project
// @Success 200 {object} models.Project
// @Param id path string true "Project ID"
// @Router /api/projects/{id} [get]
func (p *ProjectHandler) Get(c echo.Context) error {
	id := c.Param("id")
	project, err := p.projectService.Get(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, project)
}

// @ID update-project
// @Success 200 {object} models.Project
// @Param id path string true "Project ID"
// @Param project body types.UpdateProjectArgs true "Project"
// @Router /api/projects/{id} [put]
func (p *ProjectHandler) Update(c echo.Context) error {
	args := types.UpdateProjectArgs{}
	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	project, err := p.projectService.Update(args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, project)
}

// @ID delete-project
// @Success 204
// @Param id path string true "Project ID"
// @Router /api/projects/{id} [delete]
func (p *ProjectHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := p.projectService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

// @ID deploy-project
// @Success 200 {object} models.Project
// @Param id path string true "Project ID"
// @Param args body types.DeployArgs true "Deploy Args"
// @Router /api/projects/{id}/deploy [post]
func (p *ProjectHandler) Deploy(c echo.Context) error {
	args := types.DeployArgs{}
	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	err := p.projectService.Deploy(args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, nil)
}

// @ID get-network
// @Success 200 {object} models.Network
// @Param id path string true "Project ID"
// @Router /api/projects/{id}/network [get]
func (p *ProjectHandler) GetNetwork(c echo.Context) error {
	id := c.Param("id")
	project, err := p.projectService.GetNework(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, project)
}

func NewProjectHandler(projectService services.IProjectService) IProjectHandler {
	return &ProjectHandler{projectService: projectService}
}
