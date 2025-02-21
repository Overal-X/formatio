package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/overal-x/formatio/services"
	"github.com/overal-x/formatio/types"
	"github.com/samber/do"
)

type IDeploymentHandler interface {
	ListDeployments(c echo.Context) error
	ListDeploymentLogs(c echo.Context) error
}

type DeploymentHandler struct {
	deploymentService services.IDeploymentService
}

// @ID list-deployments
// @Success 200 {array} models.Deployment
// @Param project_id path string true "Project Id"
// @Param args body types.ListDeploymentsArgs true "List Deployments Args"
// @Router /api/deployments/{project_id} [get]
func (d *DeploymentHandler) ListDeployments(c echo.Context) error {
	args := types.ListDeploymentsArgs{}
	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	deployments, err := d.deploymentService.ListDeployments(args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, deployments)
}

// @ID list-deployment-logs
// @Success 200 {array} models.DeploymentLog
// @Param deployment_id path string true "Deployment Id"
// @Param args body types.ListDeploymentLogsArgs true "List Deployments Logs Args"
// @Router /api/deployments/{deployment_id}/logs/ [get]
func (d *DeploymentHandler) ListDeploymentLogs(c echo.Context) error {
	args := types.ListDeploymentLogsArgs{}
	if err := c.Bind(&args); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	deployments, err := d.deploymentService.ListDeploymentLogs(args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, deployments)
}

func NewDeploymentHandler(i *do.Injector) (IDeploymentHandler, error) {
	deploymentService := do.MustInvoke[services.IDeploymentService](i)

	return &DeploymentHandler{deploymentService: deploymentService}, nil
}
