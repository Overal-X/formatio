<script lang="ts">
	import { defaults, superForm } from 'sveltekit-superforms';
	import { zod } from 'sveltekit-superforms/adapters';
	import { page } from '$app/state';
	import { useQueryClient } from '@tanstack/svelte-query';

	import { createDeployProject, listDeploymentsQueryKey } from '$lib/api';
	import { DeployProjectSchema } from './schema';

	let { on_success }: { on_success?: () => void } = $props();

	const query_client = useQueryClient();

	const createProjectMutation = createDeployProject({
		mutation: {
			onSuccess: () => {
				on_success?.();
				query_client.invalidateQueries({ queryKey: [listDeploymentsQueryKey(page.params.id)] });
			}
		}
	});
	const { form, errors, enhance } = superForm(defaults(zod(DeployProjectSchema)), {
		dataType: 'json',
		SPA: true,
		validators: zod(DeployProjectSchema),
		onUpdate: ({ form }) => {
			if (!form.valid) return;
			$createProjectMutation.mutate({ id: page.params.id, data: form.data });
		}
	});
</script>

<form use:enhance>
	<div class="form-control">
		<label for="commit_sha">Commit SHA / Branch</label>
		<input id="commit_sha" bind:value={$form.commit_sha} />
		{#if $errors.commit_sha}<span class="form-error">{$errors.commit_sha}</span>{/if}
	</div>

	<button>Deploy</button>
</form>
