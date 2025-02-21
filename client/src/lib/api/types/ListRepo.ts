import type { TypesRepo } from './types/Repo.ts';

export type ListRepoPathParams = {
	/**
	 * @description App Id
	 * @type string
	 */
	app_id: string;
	/**
	 * @description Installation Id
	 * @type string
	 */
	installation_id: string;
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
