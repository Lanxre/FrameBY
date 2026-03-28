import { $api } from "@/composables/api/useApi";
import type { EmploymentParticipant } from "@/types/frontend/employment";

export function useEmploymentParticipants() {
	const participants = ref<EmploymentParticipant[]>([]);
	const isLoading = ref(false);
	const errorMessage = ref("");

	const fetchParticipants = async (requestId: string): Promise<void> => {
		isLoading.value = true;
		errorMessage.value = "";

		try {
			const response = await $api<{
				participants: EmploymentParticipant[];
			}>(`/employment-requests/${requestId}/participants`, {
				method: "GET",
			});
			participants.value = response.participants || [];
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка загрузки участников";
		} finally {
			isLoading.value = false;
		}
	};

	const fetchAllParticipants = async (requestIds: string[]): Promise<Map<string, EmploymentParticipant[]>> => {
		const result = new Map<string, EmploymentParticipant[]>();

		await Promise.all(
			requestIds.map(async (requestId) => {
				try {
					const response = await $api<{
						participants: EmploymentParticipant[];
					}>(`/employment-requests/${requestId}/participants`, {
						method: "GET",
					});
					result.set(requestId, response.participants || []);
				} catch (e) {
					result.set(requestId, []);
				}
			})
		);

		return result;
	};

	const reset = () => {
		participants.value = [];
		errorMessage.value = "";
	};

	return {
		participants,
		isLoading,
		errorMessage,
		fetchParticipants,
		fetchAllParticipants,
		reset,
	};
}
