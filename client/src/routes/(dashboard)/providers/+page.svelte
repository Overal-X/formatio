<script lang="ts">
	import { createListApps } from '$lib/api';
	import {
		SiBitbucket as Bitbucket,
		SiGithub as Github,
		SiGitlab as Gitlab
	} from '@icons-pack/svelte-simple-icons';

	const github_apps_query = createListApps();
</script>

<div class="grid grid-cols-3 gap-4">
	<a href="/providers/github/">
		<button class="flex items-center gap-2">
			<Github />
			<span>Github</span>

			<span class="ml-auto text-xs text-gray-200"
				>{$github_apps_query.data?.length || 0} App(s)</span
			>
		</button>
	</a>

	<a href="/providers/gitlab/">
		<button class="flex items-center gap-2">
			<Gitlab />
			<span>Gitlab</span>
			<span class="ml-auto text-xs text-gray-200">0 App(s)</span>
		</button>
	</a>

	<a href="/providers/bitbucket/">
		<button class="flex items-center gap-2">
			<Bitbucket />
			<span>Bitbucket</span>
			<span class="ml-auto text-xs text-gray-200">0 App(s)</span>
		</button>
	</a>
</div>
