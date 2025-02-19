import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type {
	ListInstallationsQueryResponse,
	ListInstallationsPathParams
} from '../types/ListInstallations.ts';
import type { QueryKey, CreateBaseQueryOptions, CreateQueryResult } from '@tanstack/svelte-query';
import { queryOptions, createQuery } from '@tanstack/svelte-query';

export const listInstallationsQueryKey = (appId: ListInstallationsPathParams['appId']) =>
	[{ url: '/api/github/installations/:appId', params: { appId: appId } }] as const;

export type ListInstallationsQueryKey = ReturnType<typeof listInstallationsQueryKey>;

/**
 * {@link /api/github/installations/:appId}
 */
export async function listInstallations(
	appId: ListInstallationsPathParams['appId'],
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<ListInstallationsQueryResponse, ResponseErrorConfig<Error>, unknown>({
		method: 'GET',
		url: `/api/github/installations/${appId}`,
		...requestConfig
	});
	return res.data;
}

export function listInstallationsQueryOptions(
	appId: ListInstallationsPathParams['appId'],
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const queryKey = listInstallationsQueryKey(appId);
	return queryOptions<
		ListInstallationsQueryResponse,
		ResponseErrorConfig<Error>,
		ListInstallationsQueryResponse,
		typeof queryKey
	>({
		enabled: !!appId,
		queryKey,
		queryFn: async ({ signal }) => {
			config.signal = signal;
			return listInstallations(appId, config);
		}
	});
}

/**
 * {@link /api/github/installations/:appId}
 */
export function createListInstallations<
	TData = ListInstallationsQueryResponse,
	TQueryData = ListInstallationsQueryResponse,
	TQueryKey extends QueryKey = ListInstallationsQueryKey
>(
	appId: ListInstallationsPathParams['appId'],
	options: {
		query?: Partial<
			CreateBaseQueryOptions<
				ListInstallationsQueryResponse,
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
	const queryKey = queryOptions?.queryKey ?? listInstallationsQueryKey(appId);

	const query = createQuery({
		...(listInstallationsQueryOptions(appId, config) as unknown as CreateBaseQueryOptions),
		queryKey,
		...(queryOptions as unknown as Omit<CreateBaseQueryOptions, 'queryKey'>)
	}) as CreateQueryResult<TData, ResponseErrorConfig<Error>> & { queryKey: TQueryKey };

	query.queryKey = queryKey as TQueryKey;

	return query;
}
