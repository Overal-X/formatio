import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type { ListAppsQueryResponse } from '../types/ListApps.ts';
import type { QueryKey, CreateBaseQueryOptions, CreateQueryResult } from '@tanstack/svelte-query';
import { queryOptions, createQuery } from '@tanstack/svelte-query';

export const listAppsQueryKey = () => [{ url: '/api/providers/github/apps' }] as const;

export type ListAppsQueryKey = ReturnType<typeof listAppsQueryKey>;

/**
 * {@link /api/providers/github/apps}
 */
export async function listApps(config: Partial<RequestConfig> & { client?: typeof client } = {}) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<ListAppsQueryResponse, ResponseErrorConfig<Error>, unknown>({
		method: 'GET',
		url: `/api/providers/github/apps`,
		...requestConfig
	});
	return res.data;
}

export function listAppsQueryOptions(
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const queryKey = listAppsQueryKey();
	return queryOptions<
		ListAppsQueryResponse,
		ResponseErrorConfig<Error>,
		ListAppsQueryResponse,
		typeof queryKey
	>({
		queryKey,
		queryFn: async ({ signal }) => {
			config.signal = signal;
			return listApps(config);
		}
	});
}

/**
 * {@link /api/providers/github/apps}
 */
export function createListApps<
	TData = ListAppsQueryResponse,
	TQueryData = ListAppsQueryResponse,
	TQueryKey extends QueryKey = ListAppsQueryKey
>(
	options: {
		query?: Partial<
			CreateBaseQueryOptions<
				ListAppsQueryResponse,
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
	const queryKey = queryOptions?.queryKey ?? listAppsQueryKey();

	const query = createQuery({
		...(listAppsQueryOptions(config) as unknown as CreateBaseQueryOptions),
		queryKey,
		...(queryOptions as unknown as Omit<CreateBaseQueryOptions, 'queryKey'>)
	}) as CreateQueryResult<TData, ResponseErrorConfig<Error>> & { queryKey: TQueryKey };

	query.queryKey = queryKey as TQueryKey;

	return query;
}
