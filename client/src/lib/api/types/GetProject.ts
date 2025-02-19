import type { ModelsProject } from './models/Project.ts';

export type GetProjectPathParams = {
	/**
	 * @description Project ID
	 * @type string
	 */
	id: string;
};

/**
 * @description OK
 */
export type GetProject200 = ModelsProject;

export type GetProjectQueryResponse = GetProject200;

export type GetProjectQuery = {
	Response: GetProject200;
	PathParams: GetProjectPathParams;
	Errors: any;
};
