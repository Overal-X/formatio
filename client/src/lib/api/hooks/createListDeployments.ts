import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type {
	ListDeploymentsQueryRequest,
	ListDeploymentsQueryResponse,
	ListDeploymentsPathParams
} from '../types/ListDeployments.ts';
import type { QueryKey, CreateBaseQueryOptions, CreateQueryResult } from '@tanstack/svelte-query';
import { queryOptions, createQuery } from '@tanstack/svelte-query';

export const listDeploymentsQueryKey = (
	project_id: ListDeploymentsPathParams['project_id'],
	data?: ListDeploymentsQueryRequest
) =>
	[
		{ url: '/api/deployments/:project_id', params: { project_id: project_id } },
		...(data ? [data] : [])
	] as const;

export type ListDeploymentsQueryKey = ReturnType<typeof listDeploymentsQueryKey>;

/**
 * {@link /api/deployments/:project_id}
 */
export async function listDeployments(
	project_id: ListDeploymentsPathParams['project_id'],
	data?: ListDeploymentsQueryRequest,
	config: Partial<RequestConfig<ListDeploymentsQueryRequest>> & { client?: typeof client } = {}
) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<
		ListDeploymentsQueryResponse,
		ResponseErrorConfig<Error>,
		ListDeploymentsQueryRequest
	>({
		method: 'GET',
		url: `/api/deployments/${project_id}`,
		data,
		...requestConfig
	});
	return res.data;
}

export function listDeploymentsQueryOptions(
	project_id: ListDeploymentsPathParams['project_id'],
	data?: ListDeploymentsQueryRequest,
	config: Partial<RequestConfig<ListDeploymentsQueryRequest>> & { client?: typeof client } = {}
) {
	const queryKey = listDeploymentsQueryKey(project_id, data);
	return queryOptions<
		ListDeploymentsQueryResponse,
		ResponseErrorConfig<Error>,
		ListDeploymentsQueryResponse,
		typeof queryKey
	>({
		enabled: !!project_id,
		queryKey,
		queryFn: async ({ signal }) => {
			config.signal = signal;
			return listDeployments(project_id, data, config);
		}
	});
}

/**
 * {@link /api/deployments/:project_id}
 */
export function createListDeployments<
	TData = ListDeploymentsQueryResponse,
	TQueryData = ListDeploymentsQueryResponse,
	TQueryKey extends QueryKey = ListDeploymentsQueryKey
>(
	project_id: ListDeploymentsPathParams['project_id'],
	data?: ListDeploymentsQueryRequest,
	options: {
		query?: Partial<
			CreateBaseQueryOptions<
				ListDeploymentsQueryResponse,
				ResponseErrorConfig<Error>,
				TData,
				TQueryData,
				TQueryKey
			>
		>;
		client?: Partial<RequestConfig<ListDeploymentsQueryRequest>> & { client?: typeof client };
	} = {}
) {
	const { query: queryOptions, client: config = {} } = options ?? {};
	const queryKey = queryOptions?.queryKey ?? listDeploymentsQueryKey(project_id, data);

	const query = createQuery({
		...(listDeploymentsQueryOptions(project_id, data, config) as unknown as CreateBaseQueryOptions),
		queryKey,
		...(queryOptions as unknown as Omit<CreateBaseQueryOptions, 'queryKey'>)
	}) as CreateQueryResult<TData, ResponseErrorConfig<Error>> & { queryKey: TQueryKey };

	query.queryKey = queryKey as TQueryKey;

	return query;
}
