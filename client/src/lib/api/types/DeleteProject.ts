export type DeleteProjectPathParams = {
	/**
	 * @description Project ID
	 * @type string
	 */
	id: string;
};

/**
 * @description No Content
 */
export type DeleteProject204 = any;

export type DeleteProjectMutationResponse = DeleteProject204;

export type DeleteProjectMutation = {
	Response: DeleteProject204;
	PathParams: DeleteProjectPathParams;
	Errors: any;
};
