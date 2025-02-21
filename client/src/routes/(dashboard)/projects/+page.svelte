<script lang="ts">
	import { createListProjects } from '$lib/api';

	const projectsQuery = createListProjects();
</script>

<div class="flex items-center justify-between">
	<h2>Projects</h2>

	<a href="/projects/new/">
		<button>New Project</button>
	</a>
</div>

<div class="grid grid-cols-4 gap-4">
	{#if $projectsQuery.data}
		{#each $projectsQuery.data as project}
			<a href={`/projects/${project.id}`}>
				<button>
					<h2 class="card-title">{project.name}</h2>
					<p class="text-sm">{project.description}</p>
				</button>
			</a>
		{/each}
	{/if}
</div>
