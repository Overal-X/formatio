import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type { GetNetworkQueryResponse, GetNetworkPathParams } from '../types/GetNetwork.ts';
import type { QueryKey, CreateBaseQueryOptions, CreateQueryResult } from '@tanstack/svelte-query';
import { queryOptions, createQuery } from '@tanstack/svelte-query';

export const getNetworkQueryKey = (id: GetNetworkPathParams['id']) =>
	[{ url: '/api/projects/:id/network', params: { id: id } }] as const;

export type GetNetworkQueryKey = ReturnType<typeof getNetworkQueryKey>;

/**
 * {@link /api/projects/:id/network}
 */
export async function getNetwork(
	id: GetNetworkPathParams['id'],
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<GetNetworkQueryResponse, ResponseErrorConfig<Error>, unknown>({
		method: 'GET',
		url: `/api/projects/${id}/network`,
		...requestConfig
	});
	return res.data;
}

export function getNetworkQueryOptions(
	id: GetNetworkPathParams['id'],
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const queryKey = getNetworkQueryKey(id);
	return queryOptions<
		GetNetworkQueryResponse,
		ResponseErrorConfig<Error>,
		GetNetworkQueryResponse,
		typeof queryKey
	>({
		enabled: !!id,
		queryKey,
		queryFn: async ({ signal }) => {
			config.signal = signal;
			return getNetwork(id, config);
		}
	});
}

/**
 * {@link /api/projects/:id/network}
 */
export function createGetNetwork<
	TData = GetNetworkQueryResponse,
	TQueryData = GetNetworkQueryResponse,
	TQueryKey extends QueryKey = GetNetworkQueryKey
>(
	id: GetNetworkPathParams['id'],
	options: {
		query?: Partial<
			CreateBaseQueryOptions<
				GetNetworkQueryResponse,
				ResponseErrorConfig<Error>,
				TData,
				TQueryData,
				TQueryKey
			>
		>;
		client?: Partial<RequestConfig> & { client?: typeof client };
	} = {}
) {
	const { query: queryOptions, client: config = {} } = options ?? {};
	const queryKey = queryOptions?.queryKey ?? getNetworkQueryKey(id);

	const query = createQuery({
		...(getNetworkQueryOptions(id, config) as unknown as CreateBaseQueryOptions),
		queryKey,
		...(queryOptions as unknown as Omit<CreateBaseQueryOptions, 'queryKey'>)
	}) as CreateQueryResult<TData, ResponseErrorConfig<Error>> & { queryKey: TQueryKey };

	query.queryKey = queryKey as TQueryKey;

	return query;
}
