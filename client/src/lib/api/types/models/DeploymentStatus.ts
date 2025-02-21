export const modelsDeploymentStatus = {
	DeploymentStatusPending: 'PENDING',
	DeploymentStatusRunning: 'RUNNING',
	DeploymentStatusSuccess: 'SUCCESS',
	DeploymentStatusFailure: 'FAILURE'
} as const;

export type ModelsDeploymentStatusEnum =
	(typeof modelsDeploymentStatus)[keyof typeof modelsDeploymentStatus];

export type ModelsDeploymentStatus = ModelsDeploymentStatusEnum;
