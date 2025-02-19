<script>
	import { ArrowRight } from 'lucide-svelte';
	import { writable } from 'svelte/store';

	import { createListApps, createListInstallations, createListRepo } from '$lib/api';

	let appId = writable('');
	let installationId = writable('');
	let repoId = writable('');

	const appsQuery = createListApps();
	const installationsQuery = $derived(
		createListInstallations($appId, {
			query: { enabled: !!$appId.length }
		})
	);
	const repoQuery = $derived(
		createListRepo($appId, $installationId, {
			query: { enabled: !!$installationId }
		})
	);
</script>

<form class="grid grid-cols-2 gap-2">
	<div class="col-span-1 space-y-2">
		<div class="form-control">
			<label for="name">Project Name</label>
			<input id="name" required />
		</div>

		<div class="form-control">
			<label for="app">App</label>
			<select id="app" required bind:value={$appId}>
				{#if $appsQuery.data}
					{#each $appsQuery.data as app}
						<option value={app.id}>{app.app_name}</option>
					{/each}
				{/if}
			</select>
		</div>

		<div class="flex w-full items-center gap-x-2">
			<div class="form-control">
				<label for="account">Account</label>
				<select id="account" required bind:value={$installationId}>
					{#if $installationsQuery.data}
						{#each $installationsQuery.data as installation}
							<option value={installation.id}>{installation.owner_username}</option>
						{/each}
					{/if}
				</select>
			</div>

			<div class="form-control">
				<label for="repo_id">Repository</label>
				<select id="repo_id" required bind:value={$repoId}>
					{#if $repoQuery.data}
						{#each $repoQuery.data as repo}
							<option value={repo.id}>{repo.owner_username}/{repo.name}</option>
						{/each}
					{/if}
				</select>
			</div>
		</div>
	</div>

	<div class="form-control col-span-1 row-span-4">
		<label for="env_variables">Environemnt Variables</label>
		<textarea class="h-full" id="env_variables" placeholder="PORT=8000"></textarea>
	</div>

	<button class="col-span-1 flex w-[99%] flex-row items-center justify-between" type="submit">
		<span>Deploy</span>

		<ArrowRight size={16} />
	</button>
</form>
