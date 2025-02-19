import type { TypesRepo } from './types/Repo.ts';

export type ListRepoPathParams = {
	/**
	 * @description App Id
	 * @type string
	 */
	appId: string;
	/**
	 * @description Installation Id
	 * @type string
	 */
	installationId: string;
};

/**
 * @description OK
 */
export type ListRepo200 = TypesRepo[];

export type ListRepoQueryResponse = ListRepo200;

export type ListRepoQuery = {
	Response: ListRepo200;
	PathParams: ListRepoPathParams;
	Errors: any;
};
