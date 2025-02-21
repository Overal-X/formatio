import { store } from '$lib/state.svelte';
import Toast from './toast.svelte';

export { Toast };

export const toast = (message: string, duration = 3000) => {
	store.update((state) => ({
		...state,
		modal: { visible: true, message }
	}));

	setTimeout(() => {
		store.update((state) => ({
			...state,
			modal: { visible: false, message: '' }
		}));
	}, duration);
};
