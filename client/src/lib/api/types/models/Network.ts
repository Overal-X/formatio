import type { ModelsProject } from './Project.ts';

export type ModelsNetwork = {
	/**
	 * @type string
	 */
	created_at: string;
	/**
	 * @type string
	 */
	host_name: string;
	/**
	 * @type string
	 */
	id: string;
	/**
	 * @type integer
	 */
	port: number;
	/**
	 * @type object
	 */
	project: ModelsProject;
	/**
	 * @type string
	 */
	project_id: string;
	/**
	 * @type string
	 */
	target_id: string;
	/**
	 * @type string
	 */
	updated_at: string;
};
