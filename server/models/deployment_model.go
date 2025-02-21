package models

import "time"

type DeploymentStatus string

const (
	DeploymentStatusPending DeploymentStatus = "PENDING"
	DeploymentStatusRunning DeploymentStatus = "RUNNING"
	DeploymentStatusSuccess DeploymentStatus = "SUCCESS"
	DeploymentStatusFailure DeploymentStatus = "FAILURE"
)

type Deployment struct {
	Id        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	ProjectId     string           `json:"project_id" gorm:"index"`
	Project       Project          `json:"project" gorm:"foreignKey:ProjectId;references:Id"`
	EnvironmentId string           `json:"environment_id" gorm:"index"`
	Environment   Environment      `json:"environment" gorm:"foreignKey:EnvironmentId;references:Id"`
	Status        DeploymentStatus `json:"status" gorm:"default:'PENDING'"`
	Message       string           `json:"message"`
}

type DeploymentLog struct {
	Id        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	DeploymentId string     `json:"deployment_id" gorm:"index"`
	Deployment   Deployment `json:"deployment" gorm:"foreignKey:DeploymentId;references:Id"`
	Message      string     `json:"message"`
}
