import { $api } from "@/composables/api/useApi";

export interface CreateEnterpriseData {
	name: string;
	address?: string;
}

export interface CreatedEnterprise {
	id: string;
	name: string;
	address?: string;
}

export function useCreateEnterprise() {
	const isLoading = ref(false);
	const errorMessage = ref("");
	const isSuccess = ref(false);
	const createdEnterprise = ref<CreatedEnterprise | null>(null);

	const create = async (data: CreateEnterpriseData): Promise<boolean> => {
		isLoading.value = true;
		errorMessage.value = "";
		isSuccess.value = false;

		try {
			const response = await $api<{ enterprise: CreatedEnterprise }>(
				"/enterprises",
				{
					method: "POST",
					body: data,
				},
			);
			createdEnterprise.value = response.enterprise;
			isSuccess.value = true;
			return true;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка создания организации";
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	const reset = () => {
		errorMessage.value = "";
		isSuccess.value = false;
		createdEnterprise.value = null;
	};

	return {
		isLoading,
		errorMessage,
		isSuccess,
		createdEnterprise,
		create,
		reset,
	};
}
