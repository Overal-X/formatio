<script lang="ts">
	import { page } from '$app/state';
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
				redirect_url: `${url}/api/github/?authId=${auth_id}&next=${page.url.href}`,
				name: `Formatio-${dayjs(new Date()).format('YYYY-MM-DD hh:mm')}`,
				url,
				hook_attributes: { url: `${url}/api/github/deploy/` },
				callback_urls: [`${url}/api/github/`],
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

<div class="w-full">
	{#if $github_apps_query.data}
		<div class="w-lg space-y-4">
			<h2>Connected Apps</h2>
			{#each $github_apps_query.data as github_app}
				<div class="group flex items-center justify-between border border-white p-4">
					<span>{github_app.app_name}</span>

					<div class="flex items-center gap-x-2">
						<a
							href={`https://github.com/apps/${github_app.app_name.toLowerCase()}/installations/new/`}
							target="_blank"
						>
							<button type="button" class="bg-black p-2 text-xs">Update</button>
						</a>
						<button type="button" class="bg-red-800 p-2 text-xs">Delete</button>
					</div>
				</div>
			{/each}
		</div>
	{:else}
		<p class="text-sm text-gray-400">
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
				class="disabled:opacity-50"
			>
				Create GitHub App
			</button>
		</form>
	{/if}
</div>
