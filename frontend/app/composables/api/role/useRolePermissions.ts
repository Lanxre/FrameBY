import { computed } from "vue";
import { useAuthStore } from "@/stores/auth";
import { FramebyAppRole, ROLE_WEIGHTS } from "~/types/frontend/enums/role";

export function useRolePermissions() {
	const authStore = useAuthStore();

	const currentRole = computed((): FramebyAppRole => {
		return (authStore.user?.role as FramebyAppRole) || FramebyAppRole.USER;
	});

	const currentWeight = computed(() => ROLE_WEIGHTS[currentRole.value] || 0);

	const getWeight = (role: string): number => {
		return ROLE_WEIGHTS[role as FramebyAppRole] || 0;
	};

	const hasPermission = (requiredRole: FramebyAppRole): boolean => {
		const requiredWeight = ROLE_WEIGHTS[requiredRole];
		return currentWeight.value >= requiredWeight;
	};

	const isExactRole = (role: FramebyAppRole): boolean => {
		return currentRole.value === role;
	};

	const canManageUser = (targetUserRole: string): boolean => {
		const targetWeight = getWeight(targetUserRole);
		return currentWeight.value > targetWeight;
	};

	return {
		currentRole,
		currentWeight,
		hasPermission,
		isExactRole,
		canManageUser,
	};
}
