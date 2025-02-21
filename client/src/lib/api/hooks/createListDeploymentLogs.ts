import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type {
	ListDeploymentLogsQueryRequest,
	ListDeploymentLogsQueryResponse,
	ListDeploymentLogsPathParams
} from '../types/ListDeploymentLogs.ts';
import type { QueryKey, CreateBaseQueryOptions, CreateQueryResult } from '@tanstack/svelte-query';
import { queryOptions, createQuery } from '@tanstack/svelte-query';

export const listDeploymentLogsQueryKey = (
	deployment_id: ListDeploymentLogsPathParams['deployment_id'],
	data?: ListDeploymentLogsQueryRequest
) =>
	[
		{ url: '/api/deployments/:deployment_id/logs/', params: { deployment_id: deployment_id } },
		...(data ? [data] : [])
	] as const;

export type ListDeploymentLogsQueryKey = ReturnType<typeof listDeploymentLogsQueryKey>;

/**
 * {@link /api/deployments/:deployment_id/logs/}
 */
export async function listDeploymentLogs(
	deployment_id: ListDeploymentLogsPathParams['deployment_id'],
	data?: ListDeploymentLogsQueryRequest,
	config: Partial<RequestConfig<ListDeploymentLogsQueryRequest>> & { client?: typeof client } = {}
) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<
		ListDeploymentLogsQueryResponse,
		ResponseErrorConfig<Error>,
		ListDeploymentLogsQueryRequest
	>({
		method: 'GET',
		url: `/api/deployments/${deployment_id}/logs/`,
		data,
		...requestConfig
	});
	return res.data;
}

export function listDeploymentLogsQueryOptions(
	deployment_id: ListDeploymentLogsPathParams['deployment_id'],
	data?: ListDeploymentLogsQueryRequest,
	config: Partial<RequestConfig<ListDeploymentLogsQueryRequest>> & { client?: typeof client } = {}
) {
	const queryKey = listDeploymentLogsQueryKey(deployment_id, data);
	return queryOptions<
		ListDeploymentLogsQueryResponse,
		ResponseErrorConfig<Error>,
		ListDeploymentLogsQueryResponse,
		typeof queryKey
	>({
		enabled: !!deployment_id,
		queryKey,
		queryFn: async ({ signal }) => {
			config.signal = signal;
			return listDeploymentLogs(deployment_id, data, config);
		}
	});
}

/**
 * {@link /api/deployments/:deployment_id/logs/}
 */
export function createListDeploymentLogs<
	TData = ListDeploymentLogsQueryResponse,
	TQueryData = ListDeploymentLogsQueryResponse,
	TQueryKey extends QueryKey = ListDeploymentLogsQueryKey
>(
	deployment_id: ListDeploymentLogsPathParams['deployment_id'],
	data?: ListDeploymentLogsQueryRequest,
	options: {
		query?: Partial<
			CreateBaseQueryOptions<
				ListDeploymentLogsQueryResponse,
				ResponseErrorConfig<Error>,
				TData,
				TQueryData,
				TQueryKey
			>
		>;
		client?: Partial<RequestConfig<ListDeploymentLogsQueryRequest>> & { client?: typeof client };
	} = {}
) {
	const { query: queryOptions, client: config = {} } = options ?? {};
	const queryKey = queryOptions?.queryKey ?? listDeploymentLogsQueryKey(deployment_id, data);

	const query = createQuery({
		...(listDeploymentLogsQueryOptions(
			deployment_id,
			data,
			config
		) as unknown as CreateBaseQueryOptions),
		queryKey,
		...(queryOptions as unknown as Omit<CreateBaseQueryOptions, 'queryKey'>)
	}) as CreateQueryResult<TData, ResponseErrorConfig<Error>> & { queryKey: TQueryKey };

	query.queryKey = queryKey as TQueryKey;

	return query;
}
