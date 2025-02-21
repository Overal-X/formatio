<script lang="ts">
	import { Check, CircleDashed, RefreshCw, X } from 'lucide-svelte';
	import { page } from '$app/state';

	import { Modal } from '$lib/components/modal';
	import DeployForm from './deploy-form.svelte';
	import { toast } from '$lib/components/toast';
	import { createListDeployments, type ModelsDeploymentStatusEnum } from '$lib/api';

	let deploy_modal_visibility = $state(false);
	const deployments_query = createListDeployments(page.params.id);
</script>

<Modal bind:visible={deploy_modal_visibility}>
	<DeployForm
		on_success={() => {
			deploy_modal_visibility = false;
			toast('Deployment in progress ...');
		}}
	/>
</Modal>

<div class="mb-6 flex items-center justify-between">
	<span>{$deployments_query.data?.length || 0} Deployment(s)</span>

	<button class="w-fit px-6 py-2" onclick={() => (deploy_modal_visibility = true)}>Deploy</button>
</div>

{#snippet deployment_item(args: { id: string; title: string; status: ModelsDeploymentStatusEnum })}
	<a
		href="/projects/{page.params.id}/deployments/{args.id}"
		class="block border border-t-0 border-white first:border-t"
	>
		<button class="flex items-center justify-start gap-4 border-0">
			<span class="w-fit">
				{#if args.status === 'SUCCESS'}
					<Check />
				{:else if args.status === 'PENDING'}
					<CircleDashed />
				{:else if args.status === 'RUNNING'}
					<RefreshCw class="animate-spin" />
				{:else}
					<X />
				{/if}
			</span>
			<p class="w-full">{args.title}</p>
		</button>
	</a>
{/snippet}

<div>
	{#if $deployments_query.data}
		{#each $deployments_query.data as deployment}
			{@render deployment_item({
				id: deployment.id,
				title: deployment.message,
				status: deployment.status
			})}
		{/each}
	{/if}

	<!-- {@render deployment_item({ id: '2', title: 'Deployment 2', status: 'pending' })}
	{@render deployment_item({ id: '3', title: 'Deployment 3', status: 'failed' })}
	{@render deployment_item({ id: '4', title: 'Deployment 4', status: 'running' })} -->
</div>
