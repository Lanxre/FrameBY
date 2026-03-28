import { $api } from "@/composables/api/useApi";
import type { StudentEmploymentListResponse } from "@/types/frontend/university";

export function useUniversityStudents() {
	const students = ref<StudentEmploymentListResponse["students"]>([]);
	const isLoading = ref(false);
	const error = ref("");
	const total = ref(0);
	const limit = ref(20);
	const offset = ref(0);

	const fetchStudents = async (newLimit?: number, newOffset?: number) => {
		isLoading.value = true;
		error.value = "";

		const queryLimit = newLimit ?? limit.value;
		const queryOffset = newOffset ?? offset.value;

		try {
			const response = await $api<StudentEmploymentListResponse>(
				`/profile/university/students?limit=${queryLimit}&offset=${queryOffset}`,
				{ method: "GET" }
			);
			students.value = response.students || [];
			total.value = response.total;
			limit.value = response.limit;
			offset.value = response.offset;
		} catch (e: any) {
			error.value = e.message || "Ошибка загрузки студентов";
		} finally {
			isLoading.value = false;
		}
	};

	const loadMore = async () => {
		if (students.value.length >= total.value || isLoading.value) return;
		await fetchStudents(limit.value, offset.value + limit.value);
	};

	const hasMore = computed(() => students.value.length < total.value);

	return {
		students,
		isLoading,
		error,
		total,
		limit,
		offset,
		fetchStudents,
		loadMore,
		hasMore,
	};
}
