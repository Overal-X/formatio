import type { TypesInstallation } from './types/Installation.ts';

export type ListInstallationsPathParams = {
	/**
	 * @description App Id
	 * @type string
	 */
	app_id: string;
};

/**
 * @description OK
 */
export type ListInstallations200 = TypesInstallation[];

export type ListInstallationsQueryResponse = ListInstallations200;

export type ListInstallationsQuery = {
	Response: ListInstallations200;
	PathParams: ListInstallationsPathParams;
	Errors: any;
};
