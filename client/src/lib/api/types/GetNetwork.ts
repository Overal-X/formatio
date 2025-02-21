import type { ModelsNetwork } from './models/Network.ts';

export type GetNetworkPathParams = {
	/**
	 * @description Project ID
	 * @type string
	 */
	id: string;
};

/**
 * @description OK
 */
export type GetNetwork200 = ModelsNetwork;

export type GetNetworkQueryResponse = GetNetwork200;

export type GetNetworkQuery = {
	Response: GetNetwork200;
	PathParams: GetNetworkPathParams;
	Errors: any;
};
