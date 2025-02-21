import type { ModelsDeployment } from './models/Deployment.ts';
import type { TypesListDeploymentsArgs } from './types/ListDeploymentsArgs.ts';

export type ListDeploymentsPathParams = {
	/**
	 * @description Project Id
	 * @type string
	 */
	project_id: string;
};

/**
 * @description OK
 */
export type ListDeployments200 = ModelsDeployment[];

/**
 * @description List Deployments Args
 */
export type ListDeploymentsQueryRequest = TypesListDeploymentsArgs;

export type ListDeploymentsQueryResponse = ListDeployments200;

export type ListDeploymentsQuery = {
	Response: ListDeployments200;
	Request: ListDeploymentsQueryRequest;
	PathParams: ListDeploymentsPathParams;
	Errors: any;
};
