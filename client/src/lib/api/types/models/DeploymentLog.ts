import type { ModelsDeployment } from './Deployment.ts';

export type ModelsDeploymentLog = {
	/**
	 * @type string
	 */
	created_at: string;
	/**
	 * @type object
	 */
	deployment: ModelsDeployment;
	/**
	 * @type string
	 */
	deployment_id: string;
	/**
	 * @type string
	 */
	id: string;
	/**
	 * @type string
	 */
	message: string;
	/**
	 * @type string
	 */
	updated_at: string;
};
