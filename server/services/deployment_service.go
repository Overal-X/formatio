package services

import (
	"github.com/overal-x/formatio/models"
	"github.com/overal-x/formatio/types"
	"gorm.io/gorm"
)

type IDeploymentService interface {
	ListDeployments(types.ListDeploymentsArgs) ([]models.Deployment, error)
	ListDeploymentLogs(types.ListDeploymentLogsArgs) ([]models.DeploymentLog, error)
}

type DeploymentService struct {
	db *gorm.DB
}

func (d *DeploymentService) ListDeployments(args types.ListDeploymentsArgs) (deployments []models.Deployment, err error) {
	err = d.db.Order("created_at desc").Find(&deployments, "project_id = ?", args.ProjectId).Error
	if err != nil {
		return nil, err
	}

	return deployments, nil
}

func (d *DeploymentService) ListDeploymentLogs(args types.ListDeploymentLogsArgs) (deployment_logs []models.DeploymentLog, err error) {
	err = d.db.Find(&deployment_logs, "deployment_id = ?", args.DeploymentId).Error
	if err != nil {
		return nil, err
	}

	return deployment_logs, nil
}

func NewDeploymentService(db *gorm.DB) IDeploymentService {
	return &DeploymentService{db: db}
}
