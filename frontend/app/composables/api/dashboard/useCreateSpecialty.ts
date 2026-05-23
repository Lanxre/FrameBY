import { $api } from "@/composables/api/useApi";

export interface CreateSpecialtyData {
	name: string;
	department_id: string;
}

export interface CreatedSpecialty {
	id: string;
	name: string;
	department_id: string;
}

export function useCreateSpecialty() {
	const isLoading = ref(false);
	const errorMessage = ref("");
	const isSuccess = ref(false);
	const createdSpecialty = ref<CreatedSpecialty | null>(null);

	const create = async (data: CreateSpecialtyData): Promise<boolean> => {
		isLoading.value = true;
		errorMessage.value = "";
		isSuccess.value = false;

		try {
			const response = await $api<{ specialty: CreatedSpecialty }>(
				"/specialties",
				{
					method: "POST",
					body: data,
				},
			);
			createdSpecialty.value = response.specialty;
			isSuccess.value = true;
			return true;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка создания специальности";
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	const reset = () => {
		errorMessage.value = "";
		isSuccess.value = false;
		createdSpecialty.value = null;
	};

	return {
		isLoading,
		errorMessage,
		isSuccess,
		createdSpecialty,
		create,
		reset,
	};
}
