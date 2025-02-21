<script lang="ts">
	import { page } from '$app/state';

	import { createListDeploymentLogs, type ModelsDeploymentLog } from '$lib/api';

	const deployment_logs_query = createListDeploymentLogs(page.params.deployment_id);
</script>

{#snippet deployment_log_item(args: ModelsDeploymentLog)}
	<button
		class="flex items-center justify-start gap-4 border-0 py-2 text-sm hover:bg-[rgba(255,255,255,0.25)]"
	>
		{args.message}
	</button>
{/snippet}

<pre class="flex flex-col border border-white bg-black">
	{#if $deployment_logs_query.data}
		{#each $deployment_logs_query.data as deployment}
			{@render deployment_log_item(deployment)}
		{/each}
	{/if}
</pre>
