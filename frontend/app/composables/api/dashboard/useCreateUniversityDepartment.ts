import { $api } from "@/composables/api/useApi";

export interface CreateUniversityDepartmentData {
	university_name: string;
	department_name: string;
	address?: string;
}

export interface CreatedUniversityDepartment {
	id: string;
	university_name: string;
	department_name: string;
	address?: string;
}

export function useCreateUniversityDepartment() {
	const isLoading = ref(false);
	const errorMessage = ref("");
	const isSuccess = ref(false);
	const createdDepartment = ref<CreatedUniversityDepartment | null>(null);

	const create = async (
		data: CreateUniversityDepartmentData,
	): Promise<boolean> => {
		isLoading.value = true;
		errorMessage.value = "";
		isSuccess.value = false;

		try {
			const response = await $api<{ department: CreatedUniversityDepartment }>(
				"/university-departments",
				{
					method: "POST",
					body: data,
				},
			);
			createdDepartment.value = response.department;
			isSuccess.value = true;
			return true;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка создания записи";
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	const reset = () => {
		errorMessage.value = "";
		isSuccess.value = false;
		createdDepartment.value = null;
	};

	return {
		isLoading,
		errorMessage,
		isSuccess,
		createdDepartment,
		create,
		reset,
	};
}
