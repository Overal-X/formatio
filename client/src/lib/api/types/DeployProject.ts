import type { ModelsProject } from './models/Project.ts';
import type { TypesDeployArgs } from './types/DeployArgs.ts';

export type DeployProjectPathParams = {
	/**
	 * @description Project ID
	 * @type string
	 */
	id: string;
};

/**
 * @description OK
 */
export type DeployProject200 = ModelsProject;

/**
 * @description Deploy Args
 */
export type DeployProjectMutationRequest = TypesDeployArgs;

export type DeployProjectMutationResponse = DeployProject200;

export type DeployProjectMutation = {
	Response: DeployProject200;
	Request: DeployProjectMutationRequest;
	PathParams: DeployProjectPathParams;
	Errors: any;
};
