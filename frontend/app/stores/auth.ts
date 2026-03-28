import { defineStore } from "pinia";
import { computed, ref } from "vue";
import type { UserEntityResponse, UserEntity } from "@/types/backend/user";
import { useApi, $api } from "~/composables/api/useApi";
import { FramebyAppRole } from "~/types/frontend/enums/role";
import { useNotificationStore } from "@/stores/notification";

export const useAuthStore = defineStore("auth", () => {
	const user = ref<UserEntity | null>(null);
	const isAuthChecking = ref(true);

	const { notify } = useNotificationStore();

	const isAuthenticated = computed(() => !!user.value);
	const isAdmin = computed(
		() =>
			user.value?.role === FramebyAppRole.BRSM &&
			user.value.subrole === "admin",
	);
	const isBRSM = computed(() => user.value?.role === FramebyAppRole.BRSM);

	async function initAuth() {
		isAuthChecking.value = true;
		try {
			const data = await $api<UserEntityResponse>("/api/auth/me");

			if (data && data.user) {
				user.value = data.user;
			} else {
				user.value = null;
			}
		} catch (e) {
			user.value = null;
		} finally {
			isAuthChecking.value = false;
		}
	}

	async function fetchUser() {
		return initAuth();
	}

	async function logout() {
		const token = useCookie("FRAMEBY_ACCESS_TOKEN");
		try {
			await useApi("/api/auth/logout", { method: "POST" });
			notify({
				type: "success",
				content: "Вы успешно вышли из системы",
				title: "Успешный выход",
			});
		} catch (e) {
			console.error("Logout error:", e);
		} finally {
			user.value = null;
			isAuthChecking.value = false;
			token.value = undefined;
		}
	}

	return {
		user,
		isAuthChecking,
		isAuthenticated,
		isAdmin,
		isBRSM,
		initAuth,
		fetchUser,
		logout,
	};
});
