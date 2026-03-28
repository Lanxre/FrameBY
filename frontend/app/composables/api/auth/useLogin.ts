import { reactive, ref } from "vue";
import { $api } from "@/composables/api/useApi";
import { useAuthStore } from "@/stores/auth";

export function useLogin() {
	const authStore = useAuthStore();
	const { notify } = useNotificationStore();
	const isLoading = ref(false);
	const errorMessage = ref("");

	const form = reactive({
		email: "",
		password: "",
	});

	const handleManualLogin = async () => {
		errorMessage.value = "";
		isLoading.value = true;

		try {
			await $api<any>("/api/auth/login", {
				method: "POST",
				body: form,
			});

			await authStore.fetchUser();
			await navigateTo("/profile");

			notify({
				title: "Успешный вход",
				content: "Вы успешно вошли в систему",
				type: "success",
			});
		} catch (err: any) {
			errorMessage.value = err.message || "Ошибка входа";
			notify({
				title: "Ошибка входа",
				content: errorMessage.value,
				type: "error",
			});
		} finally {
			isLoading.value = false;
		}
	};

	return {
		form,
		isLoading,
		errorMessage,
		handleManualLogin,
	};
}
