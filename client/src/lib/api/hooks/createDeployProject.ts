import client from '$lib/api/client';
import type { RequestConfig, ResponseErrorConfig } from '$lib/api/client';
import type {
	DeployProjectMutationRequest,
	DeployProjectMutationResponse,
	DeployProjectPathParams
} from '../types/DeployProject.ts';
import type { CreateMutationOptions } from '@tanstack/svelte-query';
import { createMutation } from '@tanstack/svelte-query';

export const deployProjectMutationKey = () => [{ url: '/api/projects/{id}/deploy' }] as const;

export type DeployProjectMutationKey = ReturnType<typeof deployProjectMutationKey>;

/**
 * {@link /api/projects/:id/deploy}
 */
export async function deployProject(
	id: DeployProjectPathParams['id'],
	data: DeployProjectMutationRequest,
	config: Partial<RequestConfig<DeployProjectMutationRequest>> & { client?: typeof client } = {}
) {
	const { client: request = client, ...requestConfig } = config;

	const res = await request<
		DeployProjectMutationResponse,
		ResponseErrorConfig<Error>,
		DeployProjectMutationRequest
	>({
		method: 'POST',
		url: `/api/projects/${id}/deploy`,
		data,
		...requestConfig
	});
	return res.data;
}

/**
 * {@link /api/projects/:id/deploy}
 */
export function createDeployProject(
	options: {
		mutation?: CreateMutationOptions<
			DeployProjectMutationResponse,
			ResponseErrorConfig<Error>,
			{ id: DeployProjectPathParams['id']; data: DeployProjectMutationRequest }
		>;
		client?: Partial<RequestConfig<DeployProjectMutationRequest>> & { client?: typeof client };
	} = {}
) {
	const { mutation: mutationOptions, client: config = {} } = options ?? {};
	const mutationKey = mutationOptions?.mutationKey ?? deployProjectMutationKey();

	return createMutation<
		DeployProjectMutationResponse,
		ResponseErrorConfig<Error>,
		{ id: DeployProjectPathParams['id']; data: DeployProjectMutationRequest }
	>({
		mutationFn: async ({ id, data }) => {
			return deployProject(id, data, config);
		},
		mutationKey,
		...mutationOptions
	});
}
