import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type { ListProjectsQueryResponse } from '../types/ListProjects.ts';
import type { QueryKey, CreateBaseQueryOptions, CreateQueryResult } from '@tanstack/svelte-query';
import { queryOptions, createQuery } from '@tanstack/svelte-query';

export const listProjectsQueryKey = () => [{ url: '/api/projects' }] as const;

export type ListProjectsQueryKey = ReturnType<typeof listProjectsQueryKey>;

/**
 * {@link /api/projects}
 */
export async function listProjects(
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<ListProjectsQueryResponse, ResponseErrorConfig<Error>, unknown>({
		method: 'GET',
		url: `/api/projects`,
		...requestConfig
	});
	return res.data;
}

export function listProjectsQueryOptions(
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const queryKey = listProjectsQueryKey();
	return queryOptions<
		ListProjectsQueryResponse,
		ResponseErrorConfig<Error>,
		ListProjectsQueryResponse,
		typeof queryKey
	>({
		queryKey,
		queryFn: async ({ signal }) => {
			config.signal = signal;
			return listProjects(config);
		}
	});
}

/**
 * {@link /api/projects}
 */
export function createListProjects<
	TData = ListProjectsQueryResponse,
	TQueryData = ListProjectsQueryResponse,
	TQueryKey extends QueryKey = ListProjectsQueryKey
>(
	options: {
		query?: Partial<
			CreateBaseQueryOptions<
				ListProjectsQueryResponse,
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
	const queryKey = queryOptions?.queryKey ?? listProjectsQueryKey();

	const query = createQuery({
		...(listProjectsQueryOptions(config) as unknown as CreateBaseQueryOptions),
		queryKey,
		...(queryOptions as unknown as Omit<CreateBaseQueryOptions, 'queryKey'>)
	}) as CreateQueryResult<TData, ResponseErrorConfig<Error>> & { queryKey: TQueryKey };

	query.queryKey = queryKey as TQueryKey;

	return query;
}
