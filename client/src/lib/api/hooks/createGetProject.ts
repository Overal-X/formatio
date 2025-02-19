import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type { GetProjectQueryResponse, GetProjectPathParams } from '../types/GetProject.ts';
import type { QueryKey, CreateBaseQueryOptions, CreateQueryResult } from '@tanstack/svelte-query';
import { queryOptions, createQuery } from '@tanstack/svelte-query';

export const getProjectQueryKey = (id: GetProjectPathParams['id']) =>
	[{ url: '/api/projects/:id', params: { id: id } }] as const;

export type GetProjectQueryKey = ReturnType<typeof getProjectQueryKey>;

/**
 * {@link /api/projects/:id}
 */
export async function getProject(
	id: GetProjectPathParams['id'],
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<GetProjectQueryResponse, ResponseErrorConfig<Error>, unknown>({
		method: 'GET',
		url: `/api/projects/${id}`,
		...requestConfig
	});
	return res.data;
}

export function getProjectQueryOptions(
	id: GetProjectPathParams['id'],
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const queryKey = getProjectQueryKey(id);
	return queryOptions<
		GetProjectQueryResponse,
		ResponseErrorConfig<Error>,
		GetProjectQueryResponse,
		typeof queryKey
	>({
		enabled: !!id,
		queryKey,
		queryFn: async ({ signal }) => {
			config.signal = signal;
			return getProject(id, config);
		}
	});
}

/**
 * {@link /api/projects/:id}
 */
export function createGetProject<
	TData = GetProjectQueryResponse,
	TQueryData = GetProjectQueryResponse,
	TQueryKey extends QueryKey = GetProjectQueryKey
>(
	id: GetProjectPathParams['id'],
	options: {
		query?: Partial<
			CreateBaseQueryOptions<
				GetProjectQueryResponse,
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
	const queryKey = queryOptions?.queryKey ?? getProjectQueryKey(id);

	const query = createQuery({
		...(getProjectQueryOptions(id, config) as unknown as CreateBaseQueryOptions),
		queryKey,
		...(queryOptions as unknown as Omit<CreateBaseQueryOptions, 'queryKey'>)
	}) as CreateQueryResult<TData, ResponseErrorConfig<Error>> & { queryKey: TQueryKey };

	query.queryKey = queryKey as TQueryKey;

	return query;
}
