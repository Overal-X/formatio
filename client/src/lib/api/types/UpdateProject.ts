import type { ModelsProject } from './models/Project.ts';
import type { TypesUpdateProjectArgs } from './types/UpdateProjectArgs.ts';

export type UpdateProjectPathParams = {
	/**
	 * @description Project ID
	 * @type string
	 */
	id: string;
};

/**
 * @description OK
 */
export type UpdateProject200 = ModelsProject;

/**
 * @description Project
 */
export type UpdateProjectMutationRequest = TypesUpdateProjectArgs;

export type UpdateProjectMutationResponse = UpdateProject200;

export type UpdateProjectMutation = {
	Response: UpdateProject200;
	Request: UpdateProjectMutationRequest;
	PathParams: UpdateProjectPathParams;
	Errors: any;
};
