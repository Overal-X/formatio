import type { ModelsGithubApp } from './models/GithubApp.ts';

/**
 * @description OK
 */
export type ListApps200 = ModelsGithubApp[];

export type ListAppsQueryResponse = ListApps200;

export type ListAppsQuery = {
	Response: ListApps200;
	Errors: any;
};
