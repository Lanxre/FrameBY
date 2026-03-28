import { $api } from "@/composables/api/useApi";
import type { StudentSquad, SquadStatus } from "@/types/frontend/student-squad";
import { FramebyAppRole } from "~/types/frontend/enums/role";

export interface SquadsResponse {
	squads: StudentSquad[];
	total: number;
	limit: number;
	offset: number;
}

export function useStudentSquads() {
	const squads = ref<StudentSquad[]>([]);
	const total = ref(0);
	const isLoading = ref(false);
	const errorMessage = ref("");

	const fetchSquads = async (
		options: { status?: string; limit?: number; offset?: number } = {},
	) => {
		isLoading.value = true;
		try {
			const params = new URLSearchParams();
			if (options.status) params.append("status", options.status);
			if (options.limit) params.append("limit", options.limit.toString());
			if (options.offset) params.append("offset", options.offset.toString());

			const response = await $api<SquadsResponse>(
				`/student-squads?${params.toString()}`,
				{ method: "GET" },
			);

			squads.value = response.squads;
			total.value = response.total;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка загрузки отрядов";
		} finally {
			isLoading.value = false;
		}
	};

	return { squads, total, isLoading, errorMessage, fetchSquads };
}

export function useMySquads() {
	const squads = ref<StudentSquad[]>([]);
	const isLoading = ref(false);
	const errorMessage = ref("");

	const fetchMySquads = async () => {
		isLoading.value = true;
		try {
			const response = await $api<{ squads: StudentSquad[] }>(
				"/student-squads/my",
				{ method: "GET" },
			);
			squads.value = response.squads;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка загрузки";
		} finally {
			isLoading.value = false;
		}
	};

	return { squads, isLoading, errorMessage, fetchMySquads };
}

export function useSquadStatuses() {
	const statuses = ref<SquadStatus[]>([]);
	const fetchStatuses = async () => {
		try {
			const response = await $api<{ statuses: SquadStatus[] }>(
				"/student-squads/statuses",
				{ method: "GET" },
			);
			statuses.value = response.statuses;
		} catch (e: any) {
			console.error(e);
		}
	};
	return { statuses, fetchStatuses };
}

export function useUpdateSquad() {
	const isLoading = ref(false);
	const errorMessage = ref("");

	const updateStatus = async (
		squadId: string,
		status: string,
	): Promise<boolean> => {
		isLoading.value = true;
		try {
			await $api(`/student-squads/${squadId}`, {
				method: "PATCH",
				body: { status },
			});
			return true;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка обновления";
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	const closeRecruitment = (squadId: string) => updateStatus(squadId, "closed");
	const openRecruitment = (squadId: string) => updateStatus(squadId, "pending");

	return { isLoading, errorMessage, closeRecruitment, openRecruitment };
}
