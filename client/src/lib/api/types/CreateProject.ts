import type { ModelsProject } from './models/Project.ts';
import type { TypesCreateProjectArgs } from './types/CreateProjectArgs.ts';

/**
 * @description Created
 */
export type CreateProject201 = ModelsProject;

/**
 * @description Project
 */
export type CreateProjectMutationRequest = TypesCreateProjectArgs;

export type CreateProjectMutationResponse = CreateProject201;

export type CreateProjectMutation = {
	Response: CreateProject201;
	Request: CreateProjectMutationRequest;
	Errors: any;
};
