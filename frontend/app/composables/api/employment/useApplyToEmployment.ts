import { $api } from "@/composables/api/useApi";
import type { EmploymentRequest } from "@/types/frontend/employment";

export function useApplyToEmployment() {
	const isLoading = ref(false);
	const errorMessage = ref("");

	const apply = async (requestId: string): Promise<boolean> => {
		isLoading.value = true;
		errorMessage.value = "";

		try {
			await $api("/employment-requests/apply", {
				method: "POST",
				body: { request_id: requestId },
			});
			return true;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка подачи заявки";
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	const reset = () => {
		errorMessage.value = "";
	};

	return {
		isLoading,
		errorMessage,
		apply,
		reset,
	};
}
