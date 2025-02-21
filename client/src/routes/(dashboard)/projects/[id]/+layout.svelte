<script lang="ts">
	import type { LayoutProps } from './$types';
	import { page } from '$app/state';

	let { children }: LayoutProps = $props();
</script>

{#snippet link(args: { text: string; href: string })}
	<a
		class="border border-r-0 border-white last:border-r aria-selected:bg-[rgba(0,0,0,0.25)]"
		href={args.href}
		role="tab"
		aria-selected={page.url.pathname === args.href}
	>
		<button class="border-0">{args.text}</button>
	</a>
{/snippet}

<div class="sticky top-6 grid grid-flow-col bg-[#333]">
	{@render link({ text: 'Overview', href: `/projects/${page.params.id}/` })}
	{@render link({ text: 'Deployments', href: `/projects/${page.params.id}/deployments/` })}
	{@render link({ text: 'Logs', href: `/projects/${page.params.id}/logs/` })}
	{@render link({ text: 'Settings', href: `/projects/${page.params.id}/settings/` })}
</div>

<section class="py-6">
	{@render children()}
</section>
