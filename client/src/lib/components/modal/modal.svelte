<script lang="ts">
	import { X } from 'lucide-svelte';

	let { visible = $bindable<boolean>(), children } = $props();
</script>

{#if visible}
	<div
		class="fixed top-0 left-0 z-20 grid h-full w-full place-content-center bg-[rgba(0,0,0,0.25)] backdrop-blur-sm"
	>
		<div class="relative min-w-[500px] border border-white bg-[#333] p-4">
			<button
				class="absolute -top-2 -right-2 w-fit bg-[#333] p-1"
				onclick={() => (visible = false)}
			>
				<X size={12} />
			</button>
			{@render children()}
		</div>
	</div>
{/if}
