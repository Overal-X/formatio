import { z } from 'zod';

export const DeployProjectSchema = z.object({
	commit_sha: z.string().nonempty()
});
