import type { ModelsDeploymentLog } from './models/DeploymentLog.ts';
import type { TypesListDeploymentLogsArgs } from './types/ListDeploymentLogsArgs.ts';

export type ListDeploymentLogsPathParams = {
	/**
	 * @description Deployment Id
	 * @type string
	 */
	deployment_id: string;
};

/**
 * @description OK
 */
export type ListDeploymentLogs200 = ModelsDeploymentLog[];

/**
 * @description List Deployments Logs Args
 */
export type ListDeploymentLogsQueryRequest = TypesListDeploymentLogsArgs;

export type ListDeploymentLogsQueryResponse = ListDeploymentLogs200;

export type ListDeploymentLogsQuery = {
	Response: ListDeploymentLogs200;
	Request: ListDeploymentLogsQueryRequest;
	PathParams: ListDeploymentLogsPathParams;
	Errors: any;
};
