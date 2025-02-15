import axios, { AxiosError, type AxiosRequestConfig, type AxiosResponse } from 'axios';

import { PUBLIC_API_URL } from '$env/static/public';

type RefreshAccessTokenMutationResponse = unknown;
type RefreshAccessTokenMutationRequest = unknown;

const ACCESS_TOKEN_KEY = 'ACCESS_TOKEN';
const REFRESH_TOKEN_KEY = 'REFRESH_TOKEN';

export const axiosInstance = axios.create({
	baseURL: PUBLIC_API_URL,
	headers: { 'Content-Type': 'application/json' }
});

// if request status is 401, add token to header
axiosInstance.interceptors.request.use((config) => {
	config.headers.Authorization = `Bearer ${localStorage.getItem(ACCESS_TOKEN_KEY)}`;

	return config;
});

// if reponse status is 401, call referesh token api and retry
axiosInstance.interceptors.response.use(
	(response) => response,
	async (error: AxiosError<Error>) => {
		if (error.config && error?.response?.status === 401) {
			return await refreshAuthToken(error.config);
		}

		return Promise.reject(error);
	}
);

function redirectToLogin() {
	window.location.href = `/login/?next=${window.location.pathname}`;

	localStorage.deleteItem(ACCESS_TOKEN_KEY);
	localStorage.deleteItem(REFRESH_TOKEN_KEY);
}

async function refreshAuthToken(errorConfig: AxiosRequestConfig) {
	const refreshToken = localStorage.getItem(REFRESH_TOKEN_KEY);
	if (!refreshToken) {
		redirectToLogin();

		return Promise.reject();
	}

	try {
		// Perform the refresh token mutation
		const { data } = await axiosClient<
			RefreshAccessTokenMutationResponse,
			unknown,
			RefreshAccessTokenMutationRequest
		>({
			method: 'POST',
			url: `/auth/refresh-access-token/`,
			data: { refreshToken }
		});

		// Update the access token in your storage
		localStorage.setItem(ACCESS_TOKEN_KEY, 'data');
		localStorage.setItem(REFRESH_TOKEN_KEY, 'data');

		// Update the Authorization header in axios
		axiosInstance.defaults.headers.common['Authorization'] = `Bearer ${data}`;

		// Retry the original request with the new token
		return axiosInstance(errorConfig);
	} catch (error) {
		// If refreshing token fails, redirect to login or handle appropriately
		redirectToLogin();
		return Promise.reject(error);
	}
}

/**
 * Subset of AxiosRequestConfig
 */
type RequestConfig<TData = unknown> = {
	baseURL?: string;
	url?: string;
	method: 'GET' | 'PUT' | 'PATCH' | 'POST' | 'DELETE';
	params?: unknown;
	data?: TData;
	responseType?: 'arraybuffer' | 'blob' | 'document' | 'json' | 'text' | 'stream';
	signal?: AbortSignal;
	headers?: AxiosRequestConfig['headers'];
};

type ResponseConfig<TData = unknown> = {
	data: TData;
	status: number;
	statusText: string;
	headers?: AxiosResponse['headers'];
};

type ResponseErrorConfig<TError = unknown> = AxiosError<TError>;

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const axiosClient = async <TData, _TError = unknown, TVariables = unknown>(
	config: RequestConfig<TVariables>
): Promise<ResponseConfig<TData>> => {
	const promise = axiosInstance.request(config).catch((e) => {
		throw e;
	});

	return promise;
};

export { axiosClient, type RequestConfig, type ResponseConfig, type ResponseErrorConfig };
export default axiosClient;
