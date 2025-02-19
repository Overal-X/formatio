import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type {
	CreateProjectMutationRequest,
	CreateProjectMutationResponse
} from '../types/CreateProject.ts';
import type { CreateMutationOptions } from '@tanstack/svelte-query';
import { createMutation } from '@tanstack/svelte-query';

export const createProjectMutationKey = () => [{ url: '/api/projects' }] as const;

export type CreateProjectMutationKey = ReturnType<typeof createProjectMutationKey>;

/**
 * {@link /api/projects}
 */
export async function createProject(
	data: CreateProjectMutationRequest,
	config: Partial<RequestConfig<CreateProjectMutationRequest>> & { client?: typeof client } = {}
) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<
		CreateProjectMutationResponse,
		ResponseErrorConfig<Error>,
		CreateProjectMutationRequest
	>({
		method: 'POST',
		url: `/api/projects`,
		data,
		...requestConfig
	});
	return res.data;
}

/**
 * {@link /api/projects}
 */
export function createCreateProject(
	options: {
		mutation?: CreateMutationOptions<
			CreateProjectMutationResponse,
			ResponseErrorConfig<Error>,
			{ data: CreateProjectMutationRequest }
		>;
		client?: Partial<RequestConfig<CreateProjectMutationRequest>> & { client?: typeof client };
	} = {}
) {
	const { mutation: mutationOptions, client: config = {} } = options ?? {};
	const mutationKey = mutationOptions?.mutationKey ?? createProjectMutationKey();

	return createMutation<
		CreateProjectMutationResponse,
		ResponseErrorConfig<Error>,
		{ data: CreateProjectMutationRequest }
	>({
		mutationFn: async ({ data }) => {
			return createProject(data, config);
		},
		mutationKey,
		...mutationOptions
	});
}
