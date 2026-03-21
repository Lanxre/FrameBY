import { reactive, ref } from "vue";
import { useApi } from "~/composables/api/useApi";
import { useAuthStore } from "@/stores/auth";

export function useLogin() {
	const authStore = useAuthStore();

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
			const { data, error } = await useApi<any>("/api/auth/login", {
				method: "POST",
				body: form,
			});

			if (error.value) {
				throw new Error(error.value.data?.error || "Неверный email или пароль");
			}

			if (data.value) {
				authStore.user = data.value;

				await navigateTo("/profile");
			}
		} catch (err: any) {
			errorMessage.value = err.message || "Ошибка входа";
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