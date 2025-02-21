<script>
	import { ArrowRight } from 'lucide-svelte';
	import { defaults, superForm } from 'sveltekit-superforms';
	import { zod } from 'sveltekit-superforms/adapters';
	import { useQueryClient } from '@tanstack/svelte-query';
	import { goto } from '$app/navigation';

	import {
		createCreateProject,
		createListApps,
		createListInstallations,
		createListRepo,
		listProjectsQueryKey
	} from '$lib/api';
	import { CreateProjectSchema } from './schema';

	const queryClient = useQueryClient();

	const createProjectMutation = createCreateProject({
		mutation: {
			onSuccess: () => {
				queryClient.invalidateQueries({ queryKey: listProjectsQueryKey() });
				goto('/projects');
			}
		}
	});
	const { form, errors, enhance } = superForm(defaults(zod(CreateProjectSchema)), {
		dataType: 'json',
		SPA: true,
		validators: zod(CreateProjectSchema),
		onUpdate: ({ form }) => {
			if (!form.valid) return;

			const repo = $repoQuery.data?.find((r) => r.id === +form.data.repo_id);

			$createProjectMutation.mutate({
				data: { ...form.data, repo_fullname: repo ? `${repo.owner_username}/${repo.name}` : '' }
			});
		}
	});

	const appsQuery = createListApps();
	const installationsQuery = $derived(
		createListInstallations($form.app_id, {
			query: { enabled: !!$form.app_id.length }
		})
	);
	const repoQuery = $derived(
		createListRepo($form.app_id, $form.installation_id, {
			query: { enabled: !!$form.installation_id }
		})
	);
</script>

<form class="grid grid-cols-2 gap-2" use:enhance>
	<div class="col-span-1 space-y-2">
		<div class="form-control">
			<label for="name">Name</label>
			<input id="name" bind:value={$form.name} />
			{#if $errors.name}<span class="form-error">{$errors.name}</span>{/if}
		</div>

		<div class="form-control">
			<label for="name">Description</label>
			<input id="name" bind:value={$form.description} />
			{#if $errors.description}<span class="form-error">{$errors.description}</span>{/if}
		</div>

		<div class="form-control">
			<label for="app">App</label>
			<select id="app" bind:value={$form.app_id}>
				{#if $appsQuery.data}
					{#each $appsQuery.data as app}
						<option value={app.id}>{app.app_name}</option>
					{/each}
				{/if}
			</select>
			{#if $errors.app_id}<span class="form-error">{$errors.app_id}</span>{/if}
		</div>

		<div class="flex w-full items-center gap-x-2">
			<div class="form-control">
				<label for="account">Account</label>
				<select id="account" bind:value={$form.installation_id}>
					{#if $installationsQuery.data}
						{#each $installationsQuery.data as installation}
							<option value={installation.id.toString()}>{installation.owner_username}</option>
						{/each}
					{/if}
				</select>
				{#if $errors.installation_id}<span class="form-error">{$errors.installation_id}</span>{/if}
			</div>

			<div class="form-control">
				<label for="repo_id">Repository</label>
				<select id="repo_id" bind:value={$form.repo_id}>
					{#if $repoQuery.data}
						{#each $repoQuery.data as repo}
							<option value={repo.id.toString()}>{repo.owner_username}/{repo.name}</option>
						{/each}
					{/if}
				</select>
				{#if $errors.repo_id}<span class="form-error">{$errors.repo_id}</span>{/if}
			</div>
		</div>
	</div>

	<div class="form-control col-span-1 row-span-4">
		<label for="env_variables">Environemnt Variables</label>
		<textarea class="h-full" id="env_variables" placeholder="PORT=8000" bind:value={$form.variables}
		></textarea>
		{#if $errors.variables}<span class="form-error">{$errors.variables}</span>{/if}
	</div>

	<button class="col-span-1 flex w-[99%] flex-row items-center justify-between" type="submit">
		<span>Deploy</span>

		<ArrowRight size={16} />
	</button>
</form>
