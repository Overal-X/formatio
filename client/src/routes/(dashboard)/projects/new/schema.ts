import { z } from 'zod';

export const CreateProjectSchema = z.object({
	name: z.string().nonempty().max(100),
	description: z.string().default(''),
	app_id: z.string().nonempty(),
	installation_id: z.string().nonempty(),
	repo_id: z.string().nonempty(),
	auto_deploy: z.boolean().default(true),
	require_approval: z.boolean().default(false),
	variables: z.string().default('')
});

export type CreateProjectSchemaType = z.infer<typeof CreateProjectSchema>;
