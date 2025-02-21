import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type {
	ListInstallationsQueryResponse,
	ListInstallationsPathParams
} from '../types/ListInstallations.ts';
import type { QueryKey, CreateBaseQueryOptions, CreateQueryResult } from '@tanstack/svelte-query';
import { queryOptions, createQuery } from '@tanstack/svelte-query';

export const listInstallationsQueryKey = (app_id: ListInstallationsPathParams['app_id']) =>
	[{ url: '/api/github/installations/:app_id', params: { app_id: app_id } }] as const;

export type ListInstallationsQueryKey = ReturnType<typeof listInstallationsQueryKey>;

/**
 * {@link /api/github/installations/:app_id}
 */
export async function listInstallations(
	app_id: ListInstallationsPathParams['app_id'],
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<ListInstallationsQueryResponse, ResponseErrorConfig<Error>, unknown>({
		method: 'GET',
		url: `/api/github/installations/${app_id}`,
		...requestConfig
	});
	return res.data;
}

export function listInstallationsQueryOptions(
	app_id: ListInstallationsPathParams['app_id'],
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const queryKey = listInstallationsQueryKey(app_id);
	return queryOptions<
		ListInstallationsQueryResponse,
		ResponseErrorConfig<Error>,
		ListInstallationsQueryResponse,
		typeof queryKey
	>({
		enabled: !!app_id,
		queryKey,
		queryFn: async ({ signal }) => {
			config.signal = signal;
			return listInstallations(app_id, config);
		}
	});
}

/**
 * {@link /api/github/installations/:app_id}
 */
export function createListInstallations<
	TData = ListInstallationsQueryResponse,
	TQueryData = ListInstallationsQueryResponse,
	TQueryKey extends QueryKey = ListInstallationsQueryKey
>(
	app_id: ListInstallationsPathParams['app_id'],
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
	const queryKey = queryOptions?.queryKey ?? listInstallationsQueryKey(app_id);

	const query = createQuery({
		...(listInstallationsQueryOptions(app_id, config) as unknown as CreateBaseQueryOptions),
		queryKey,
		...(queryOptions as unknown as Omit<CreateBaseQueryOptions, 'queryKey'>)
	}) as CreateQueryResult<TData, ResponseErrorConfig<Error>> & { queryKey: TQueryKey };

	query.queryKey = queryKey as TQueryKey;

	return query;
}
