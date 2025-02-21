import type { ModelsProject } from './Project.ts';

export type ModelsEnvironment = {
	/**
	 * @type string
	 */
	created_at: string;
	/**
	 * @type string
	 */
	id: string;
	/**
	 * @type boolean
	 */
	is_active: boolean;
	/**
	 * @type string
	 */
	name: string;
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
	updated_at: string;
	/**
	 * @type string
	 */
	variables: string;
};
