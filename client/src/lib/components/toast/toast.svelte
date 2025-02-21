<script lang="ts">
	import { store } from '$lib/state.svelte';
</script>

{#if $store.modal.visible}
	<div class="fixed right-6 bottom-6 w-[400px] border border-white bg-[#333] p-4">
		<p>{$store.modal.message}</p>
	</div>
{/if}
