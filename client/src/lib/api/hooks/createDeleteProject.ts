import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type {
	DeleteProjectMutationResponse,
	DeleteProjectPathParams
} from '../types/DeleteProject.ts';
import type { CreateMutationOptions } from '@tanstack/svelte-query';
import { createMutation } from '@tanstack/svelte-query';

export const deleteProjectMutationKey = () => [{ url: '/api/projects/{id}' }] as const;

export type DeleteProjectMutationKey = ReturnType<typeof deleteProjectMutationKey>;

/**
 * {@link /api/projects/:id}
 */
export async function deleteProject(
	id: DeleteProjectPathParams['id'],
	config: Partial<RequestConfig> & { client?: typeof client } = {}
) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<DeleteProjectMutationResponse, ResponseErrorConfig<Error>, unknown>({
		method: 'DELETE',
		url: `/api/projects/${id}`,
		...requestConfig
	});
	return res.data;
}

/**
 * {@link /api/projects/:id}
 */
export function createDeleteProject(
	options: {
		mutation?: CreateMutationOptions<
			DeleteProjectMutationResponse,
			ResponseErrorConfig<Error>,
			{ id: DeleteProjectPathParams['id'] }
		>;
		client?: Partial<RequestConfig> & { client?: typeof client };
	} = {}
) {
	const { mutation: mutationOptions, client: config = {} } = options ?? {};
	const mutationKey = mutationOptions?.mutationKey ?? deleteProjectMutationKey();

	return createMutation<
		DeleteProjectMutationResponse,
		ResponseErrorConfig<Error>,
		{ id: DeleteProjectPathParams['id'] }
	>({
		mutationFn: async ({ id }) => {
			return deleteProject(id, config);
		},
		mutationKey,
		...mutationOptions
	});
}
