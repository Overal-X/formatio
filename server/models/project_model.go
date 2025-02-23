package models

import "time"

type Project struct {
	Id        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name            string `json:"name"`
	Description     string `json:"description"`
	AppId           string `json:"app_id"`
	InstallationId  string `json:"installation_id"`
	RepoId          int64  `json:"repo_id"`
	RepoFullname    string `json:"repo_fullname"`
	AutoDeploy      bool   `json:"auto_deploy" gorm:"default:true"`
	RequireApproval bool   `json:"require_approval" gorm:"default:false"`
}

type Environment struct {
	Id        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name      string  `json:"name"`
	ProjectId string  `json:"project_id" gorm:"index"`
	Project   Project `json:"project" gorm:"foreignKey:ProjectId;references:Id"`
	IsActive  bool    `json:"is_active" gorm:"default:true"`
	Variables string  `json:"variables"`
}

type Network struct {
	Id        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	ProjectId string  `json:"project_id" gorm:"index"`
	Project   Project `json:"project" gorm:"foreignKey:ProjectId;references:Id"`
	ProcessId string  `json:"target_id"`
	Port      int     `json:"port"`
	HostName  string  `json:"host_name"`
}
