<script lang="ts">
	import { PUBLIC_API_URL } from '$env/static/public';
	import { createListApps } from '$lib/api';
	import dayjs from 'dayjs';

	const url = PUBLIC_API_URL;
	const auth_id = crypto.randomUUID();

	let is_organization = $state(false);
	let organization = $state('');
	let form_action_url = $derived(
		is_organization
			? `https://github.com/organizations/${organization}/settings/apps/new?state=gh_init:${auth_id}`
			: `https://github.com/settings/apps/new?state=gh_init:${auth_id}`
	);
	const manifest = $derived(
		JSON.stringify(
			{
				redirect_url: `${url}/api/providers/github?authId=${auth_id}&next=http://localhost:5173/providers/github/`,
				name: `Formatio-${dayjs(new Date()).format('YYYY-MM-DD')}`,
				url,
				hook_attributes: {
					url: `${url}/api/deploy/github`
				},
				callback_urls: [`${url}/api/providers/github`],
				public: false,
				request_oauth_on_install: true,
				default_permissions: {
					contents: 'read',
					metadata: 'read',
					emails: 'read',
					pull_requests: 'write'
				},
				default_events: ['pull_request', 'push']
			},
			null,
			4
		)
	);

	const github_apps_query = createListApps();
</script>

<div class="grid h-full w-full place-content-center">
	{#if $github_apps_query.data}
		<div class="w-lg space-y-2 py-4">
			<h2 class="py-4">Connect Apps</h2>
			{#each $github_apps_query.data as github_app}
				<div class="group flex items-center justify-between border border-white p-4">
					<span>{github_app.appName}</span>

					<div class="flex items-center gap-x-2">
						<button type="button" class="bg-black p-2 text-xs text-white hover:cursor-pointer"
							>Update</button
						>
						<button type="button" class="bg-red-800 p-2 text-xs text-white hover:cursor-pointer"
							>Delete</button
						>
					</div>
				</div>
			{/each}
		</div>
	{:else}
		<p class="max-w-2xl text-sm text-gray-400">
			To integrate your GitHub account with our services, you'll need to create and install a GitHub
			app. This process is straightforward and only takes a few minutes. Click the "Create Github
			App" button to get started.
		</p>
		<div class="mt-4 flex w-full flex-col gap-4">
			<div class="flex flex-row gap-4">
				<span>Organization?</span>
				<input type="checkbox" bind:checked={is_organization} />
			</div>

			{#if is_organization}
				<input
					required
					placeholder="Organization name"
					bind:value={organization}
					class="border border-white p-4 text-sm"
				/>
			{/if}
		</div>
		<form action={form_action_url} method="post" class="w-lg">
			<input type="text" name="manifest" id="manifest" defaultvalue={manifest} class="invisible" />
			<br />

			<button
				disabled={is_organization && organization.length < 1}
				type="submit"
				class="bg-black p-4 text-sm transition-all duration-150 hover:cursor-pointer"
			>
				Create GitHub App
			</button>
		</form>
	{/if}
</div>
