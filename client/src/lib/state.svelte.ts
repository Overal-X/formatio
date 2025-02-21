import { writable } from 'svelte/store';

interface Store {
	modal: { visible: boolean; message?: string };
}

export const store = writable<Store>({
	modal: { visible: false }
});
