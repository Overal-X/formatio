import { defineConfig } from '@kubb/core';
import { pluginOas } from '@kubb/plugin-oas';
import { pluginSvelteQuery } from '@kubb/plugin-svelte-query';
import { pluginTs } from '@kubb/plugin-ts';

export default defineConfig({
	input: { path: '../server/docs/swagger.yaml' },
	output: { path: './src/lib/api' },
	hooks: { done: ['npx prettier --write ./src/lib/api'] },
	plugins: [
		pluginOas(),
		pluginTs(),
		pluginSvelteQuery({
			output: { path: './hooks' },
			client: { importPath: '$lib/api/client' },
			mutation: {
				methods: ['post', 'put', 'patch', 'delete']
			},
			query: { methods: ['get'] }
		})
	]
});
