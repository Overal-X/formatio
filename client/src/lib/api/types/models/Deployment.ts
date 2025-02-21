import type { ModelsDeploymentStatus } from './DeploymentStatus.ts';
import type { ModelsEnvironment } from './Environment.ts';
import type { ModelsProject } from './Project.ts';

export type ModelsDeployment = {
	/**
	 * @type string
	 */
	created_at: string;
	/**
	 * @type object
	 */
	environment: ModelsEnvironment;
	/**
	 * @type string
	 */
	environment_id: string;
	/**
	 * @type string
	 */
	id: string;
	/**
	 * @type string
	 */
	message: string;
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
	status: ModelsDeploymentStatus;
	/**
	 * @type string
	 */
	updated_at: string;
};
