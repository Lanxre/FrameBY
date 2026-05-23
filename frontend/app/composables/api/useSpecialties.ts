import { ref } from "vue";
import { $api } from "@/composables/api/useApi";

export interface Specialty {
	id: string;
	name: string;
	department_id: string;
}

export function useSpecialties() {
	const specialties = ref<Specialty[]>([]);
	const isLoading = ref(false);
	const errorMessage = ref("");

	const fetchByUniversityDepartment = async (universityDeptId: string) => {
		isLoading.value = true;
		errorMessage.value = "";

		try {
			const res = await $api<Specialty[]>(
				`/specialties?university_department_id=${universityDeptId}`,
			);
      if (res) specialties.value = res;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка загрузки специальностей";
		} finally {
			isLoading.value = false;
		}
	};

	return { specialties, isLoading, errorMessage, fetchByUniversityDepartment };
}
