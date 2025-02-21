import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type { ListRepoQueryResponse, ListRepoPathParams } from '../types/ListRepo.ts';
import type { QueryKey, CreateBaseQueryOptions, CreateQueryResult } from '@tanstack/svelte-query';
import { queryOptions, createQuery } from '@tanstack/svelte-query';

export const listRepoQueryKey = (
	app_id: ListRepoPathParams['app_id'],
	installation_id: ListRepoPathParams['installation_id']
) =>
	[
		{
			url: '/api/github/repos/:app_id/:installation_id',
			params: { app_id: app_id, installation_id: installation_id }
		}
	] as const;

export type ListRepoQueryKey = ReturnType<typeof listRepoQueryKey>;

/**
 * {@link /api/github/repos/:app_id/:installation_id}
 */
export async function listRepo(
	app_id: ListRepoPathParams['app_id'],
	installation_id: ListRepoPathParams['installation_id'],
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<ListRepoQueryResponse, ResponseErrorConfig<Error>, unknown>({
		method: 'GET',
		url: `/api/github/repos/${app_id}/${installation_id}`,
		...requestConfig
	});
	return res.data;
}

export function listRepoQueryOptions(
	app_id: ListRepoPathParams['app_id'],
	installation_id: ListRepoPathParams['installation_id'],
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const queryKey = listRepoQueryKey(app_id, installation_id);
	return queryOptions<
		ListRepoQueryResponse,
		ResponseErrorConfig<Error>,
		ListRepoQueryResponse,
		typeof queryKey
	>({
		enabled: !!(app_id && installation_id),
		queryKey,
		queryFn: async ({ signal }) => {
			config.signal = signal;
			return listRepo(app_id, installation_id, config);
		}
	});
}

/**
 * {@link /api/github/repos/:app_id/:installation_id}
 */
export function createListRepo<
	TData = ListRepoQueryResponse,
	TQueryData = ListRepoQueryResponse,
	TQueryKey extends QueryKey = ListRepoQueryKey
>(
	app_id: ListRepoPathParams['app_id'],
	installation_id: ListRepoPathParams['installation_id'],
	options: {
		query?: Partial<
			CreateBaseQueryOptions<
				ListRepoQueryResponse,
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
	const queryKey = queryOptions?.queryKey ?? listRepoQueryKey(app_id, installation_id);

	const query = createQuery({
		...(listRepoQueryOptions(app_id, installation_id, config) as unknown as CreateBaseQueryOptions),
		queryKey,
		...(queryOptions as unknown as Omit<CreateBaseQueryOptions, 'queryKey'>)
	}) as CreateQueryResult<TData, ResponseErrorConfig<Error>> & { queryKey: TQueryKey };

	query.queryKey = queryKey as TQueryKey;

	return query;
}
