import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type {
	UpdateProjectMutationRequest,
	UpdateProjectMutationResponse,
	UpdateProjectPathParams
} from '../types/UpdateProject.ts';
import type { CreateMutationOptions } from '@tanstack/svelte-query';
import { createMutation } from '@tanstack/svelte-query';

export const updateProjectMutationKey = () => [{ url: '/api/projects/{id}' }] as const;

export type UpdateProjectMutationKey = ReturnType<typeof updateProjectMutationKey>;

/**
 * {@link /api/projects/:id}
 */
export async function updateProject(
	id: UpdateProjectPathParams['id'],
	data: UpdateProjectMutationRequest,
	config: Partial<RequestConfig<UpdateProjectMutationRequest>> & { client?: typeof client } = {}
) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<
		UpdateProjectMutationResponse,
		ResponseErrorConfig<Error>,
		UpdateProjectMutationRequest
	>({
		method: 'PUT',
		url: `/api/projects/${id}`,
		data,
		...requestConfig
	});
	return res.data;
}

/**
 * {@link /api/projects/:id}
 */
export function createUpdateProject(
	options: {
		mutation?: CreateMutationOptions<
			UpdateProjectMutationResponse,
			ResponseErrorConfig<Error>,
			{ id: UpdateProjectPathParams['id']; data: UpdateProjectMutationRequest }
		>;
		client?: Partial<RequestConfig<UpdateProjectMutationRequest>> & { client?: typeof client };
	} = {}
) {
	const { mutation: mutationOptions, client: config = {} } = options ?? {};
	const mutationKey = mutationOptions?.mutationKey ?? updateProjectMutationKey();

	return createMutation<
		UpdateProjectMutationResponse,
		ResponseErrorConfig<Error>,
		{ id: UpdateProjectPathParams['id']; data: UpdateProjectMutationRequest }
	>({
		mutationFn: async ({ id, data }) => {
			return updateProject(id, data, config);
		},
		mutationKey,
		...mutationOptions
	});
}
