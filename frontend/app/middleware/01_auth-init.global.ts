import { useAuthStore } from "~/stores/auth";

export default defineNuxtRouteMiddleware(async (to, from) => {
	const authStore = useAuthStore();

	const token = useCookie("FRAMEBY_ACCESS_TOKEN");

	if (token.value && !authStore.isAuthenticated) {
		try {
			await authStore.initAuth();
		} catch (e) {
			token.value = null;
		}
	}
});
