<script lang="ts">
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';
	import { SvelteQueryDevtools } from '@tanstack/svelte-query-devtools';

	import '../app.css';

	let { children } = $props();

	const queryClient = new QueryClient({
		defaultOptions: {
			queries: {
				refetchOnWindowFocus: false
			}
		}
	});
</script>

<svelte:head>
	<title>Formatio</title>
</svelte:head>

<QueryClientProvider client={queryClient}>
	<SvelteQueryDevtools />
	{@render children()}
</QueryClientProvider>
