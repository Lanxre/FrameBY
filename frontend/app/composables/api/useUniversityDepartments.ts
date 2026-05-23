import { ref } from "vue";
import { $api } from "@/composables/api/useApi";
import type { UniversityDepartment } from "~/types/frontend/university";

export function useUniversityDepartments() {
  const departments = ref<UniversityDepartment[]>([]);
  const isLoading = ref(false);
  const errorMessage = ref("");

  const fetchDepartments = async () => {
    isLoading.value = true;
    errorMessage.value = "";

    try {
      const res = await $api<{ departments: UniversityDepartment[] }>(
        "/university-departments",
      );
      if (res?.departments) departments.value = res.departments;
    } catch (e: any) {
      errorMessage.value = e.message || "Ошибка загрузки данных";
    } finally {
      isLoading.value = false;
    }
  };

  return {
    departments,
    isLoading,
    errorMessage,
    fetchDepartments,
  };
}
