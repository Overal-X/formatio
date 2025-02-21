<script lang="ts">
	import { page } from '$app/state';
	import { Link } from 'lucide-svelte';
	import { SiGithub as Github } from '@icons-pack/svelte-simple-icons';

	import { createGetNetwork, createGetProject } from '$lib/api';

	const projectQuery = createGetProject(page.params.id);
	const networkQuery = createGetNetwork(page.params.id);
</script>

<div class="space-y-2">
	{#if $projectQuery.data}
		<a
			href={`https://github.com/${$projectQuery.data.repo_fullname}`}
			target="_blank"
			class="flex w-fit items-center gap-x-1"
		>
			<Github size={12} />

			<span>
				{$projectQuery.data.repo_fullname}
			</span>
		</a>
	{/if}

	{#if $networkQuery.data}
		<a
			href={`http://${$networkQuery.data.host_name}`}
			target="_blank"
			class="flex w-fit items-center gap-x-1"
		>
			<Link size={12} />

			<span>
				{`http://${$networkQuery.data.host_name}`}
			</span>
		</a>
	{/if}
</div>
