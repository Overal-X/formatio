import type { ModelsProject } from './models/Project.ts';

/**
 * @description Accepted
 */
export type ListProjects202 = ModelsProject[];

export type ListProjectsQueryResponse = ListProjects202;

export type ListProjectsQuery = {
	Response: ListProjects202;
	Errors: any;
};
