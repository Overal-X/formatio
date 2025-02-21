import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type { CreateAppQueryResponse } from '../types/CreateApp.ts';
import type { QueryKey, CreateBaseQueryOptions, CreateQueryResult } from '@tanstack/svelte-query';
import { queryOptions, createQuery } from '@tanstack/svelte-query';

export const createAppQueryKey = () => [{ url: '/api/github' }] as const;

export type CreateAppQueryKey = ReturnType<typeof createAppQueryKey>;

/**
 * {@link /api/github}
 */
export async function createApp(config: Partial<RequestConfig> & { client?: typeof client } = {}) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<CreateAppQueryResponse, ResponseErrorConfig<Error>, unknown>({
		method: 'GET',
		url: `/api/github`,
		...requestConfig
	});
	return res.data;
}

export function createAppQueryOptions(
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const queryKey = createAppQueryKey();
	return queryOptions<
		CreateAppQueryResponse,
		ResponseErrorConfig<Error>,
		CreateAppQueryResponse,
		typeof queryKey
	>({
		queryKey,
		queryFn: async ({ signal }) => {
			config.signal = signal;
			return createApp(config);
		}
	});
}

/**
 * {@link /api/github}
 */
export function createCreateApp<
	TData = CreateAppQueryResponse,
	TQueryData = CreateAppQueryResponse,
	TQueryKey extends QueryKey = CreateAppQueryKey
>(
	options: {
		query?: Partial<
			CreateBaseQueryOptions<
				CreateAppQueryResponse,
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
	const queryKey = queryOptions?.queryKey ?? createAppQueryKey();

	const query = createQuery({
		...(createAppQueryOptions(config) as unknown as CreateBaseQueryOptions),
		queryKey,
		...(queryOptions as unknown as Omit<CreateBaseQueryOptions, 'queryKey'>)
	}) as CreateQueryResult<TData, ResponseErrorConfig<Error>> & { queryKey: TQueryKey };

	query.queryKey = queryKey as TQueryKey;

	return query;
}
