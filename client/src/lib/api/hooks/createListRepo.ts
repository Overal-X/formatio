import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type { ListRepoQueryResponse, ListRepoPathParams } from '../types/ListRepo.ts';
import type { QueryKey, CreateBaseQueryOptions, CreateQueryResult } from '@tanstack/svelte-query';
import { queryOptions, createQuery } from '@tanstack/svelte-query';

export const listRepoQueryKey = (
	appId: ListRepoPathParams['appId'],
	installationId: ListRepoPathParams['installationId']
) =>
	[
		{
			url: '/api/github/repos/:appId/:installationId',
			params: { appId: appId, installationId: installationId }
		}
	] as const;

export type ListRepoQueryKey = ReturnType<typeof listRepoQueryKey>;

/**
 * {@link /api/github/repos/:appId/:installationId}
 */
export async function listRepo(
	appId: ListRepoPathParams['appId'],
	installationId: ListRepoPathParams['installationId'],
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<ListRepoQueryResponse, ResponseErrorConfig<Error>, unknown>({
		method: 'GET',
		url: `/api/github/repos/${appId}/${installationId}`,
		...requestConfig
	});
	return res.data;
}

export function listRepoQueryOptions(
	appId: ListRepoPathParams['appId'],
	installationId: ListRepoPathParams['installationId'],
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const queryKey = listRepoQueryKey(appId, installationId);
	return queryOptions<
		ListRepoQueryResponse,
		ResponseErrorConfig<Error>,
		ListRepoQueryResponse,
		typeof queryKey
	>({
		enabled: !!(appId && installationId),
		queryKey,
		queryFn: async ({ signal }) => {
			config.signal = signal;
			return listRepo(appId, installationId, config);
		}
	});
}

/**
 * {@link /api/github/repos/:appId/:installationId}
 */
export function createListRepo<
	TData = ListRepoQueryResponse,
	TQueryData = ListRepoQueryResponse,
	TQueryKey extends QueryKey = ListRepoQueryKey
>(
	appId: ListRepoPathParams['appId'],
	installationId: ListRepoPathParams['installationId'],
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
	const queryKey = queryOptions?.queryKey ?? listRepoQueryKey(appId, installationId);

	const query = createQuery({
		...(listRepoQueryOptions(appId, installationId, config) as unknown as CreateBaseQueryOptions),
		queryKey,
		...(queryOptions as unknown as Omit<CreateBaseQueryOptions, 'queryKey'>)
	}) as CreateQueryResult<TData, ResponseErrorConfig<Error>> & { queryKey: TQueryKey };

	query.queryKey = queryKey as TQueryKey;

	return query;
}
