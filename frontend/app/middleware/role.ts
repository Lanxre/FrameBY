import { FramebyAppRole, ROLE_WEIGHTS } from "@/types/frontend/enums/role";

export default defineNuxtRouteMiddleware((to) => {
	const authStore = useAuthStore();

	const minRole = to.meta.minRole as FramebyAppRole | undefined;

	if (!minRole) return;

	const userRole =
		(authStore.user?.role as FramebyAppRole) || FramebyAppRole.USER;
	const userWeight = ROLE_WEIGHTS[userRole] || 0;
	const requiredWeight = ROLE_WEIGHTS[minRole] || 0;

	if (userWeight < requiredWeight) {
		return navigateTo("/");
	}
});
