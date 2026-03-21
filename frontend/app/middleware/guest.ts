export default defineNuxtRouteMiddleware((to, from) => {
	const authStore = useAuthStore();
	
	const token = useCookie('FRAMEBY_ACCESS_TOKEN');
	
	if (token.value || authStore.isAuthenticated) {
		return navigateTo("/");
	}
});
