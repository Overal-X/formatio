<script lang="ts">
	let { children } = $props();

	const sidebar_items = [
		{
			title: 'Home',
			link: '/'
		},
		{
			title: 'Projects',
			link: '/projects'
		},
		{
			title: 'Provider',
			link: '/providers'
		}
	];
</script>

{#snippet sidebar_item(args: { title: string; link: string })}
	<a href={args.link}><button class="border-x-0 border-t-0">{args.title}</button></a>
{/snippet}

<main class="flex h-full w-full items-start justify-center gap-x-6 p-6">
	<aside class="sticky top-6 w-[20%] border border-white">
		{#each sidebar_items as item}
			{@render sidebar_item(item)}
		{/each}

		<button class="mt-auto border-x-0 border-t-0">Logout</button>
	</aside>

	<section class="h-full w-[80%]">
		{@render children()}
	</section>
</main>
