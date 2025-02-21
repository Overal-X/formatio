package types

type ListDeploymentsArgs struct {
	ProjectId string `json:"-" param:"project_id"`
}

type ListDeploymentLogsArgs struct {
	DeploymentId string `json:"-" param:"deployment_id"`
}
