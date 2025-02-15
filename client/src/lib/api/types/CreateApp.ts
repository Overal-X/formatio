import type { ModelsGithubApp } from './models/GithubApp.ts';

/**
 * @description Created
 */
export type CreateApp201 = ModelsGithubApp;

export type CreateAppQueryResponse = CreateApp201;

export type CreateAppQuery = {
	Response: CreateApp201;
	Errors: any;
};
