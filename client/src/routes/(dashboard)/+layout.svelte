<script lang="ts">
	import { page } from '$app/state';

	let { children } = $props();

	const base_pathname = $derived.by(() => {
		const p = page.url.pathname
			.split('/')
			.filter((path) => path !== '')
			.at(0);
		if (!p) return '/';

		return '/' + p;
	});
</script>

{#snippet sidebar_item(args: { text: string; href: string })}
	<a href={args.href}>
		<button
			class="border-x-0 border-t-0 aria-selected:bg-[rgba(0,0,0,0.25)]"
			role="tab"
			aria-selected={base_pathname === args.href}
		>
			{args.text}
		</button>
	</a>
{/snippet}

<main class="flex h-full w-full items-start justify-center gap-x-6 p-6">
	<aside class="sticky top-6 w-[20%] border border-white">
		{@render sidebar_item({ text: 'Home', href: '/' })}
		{@render sidebar_item({ text: 'Projects', href: '/projects' })}
		{@render sidebar_item({ text: 'Providers', href: '/providers' })}

		<button class="mt-auto border-x-0 border-t-0">Logout</button>
	</aside>

	<section class="h-full w-[80%]">
		{@render children()}
	</section>
</main>
