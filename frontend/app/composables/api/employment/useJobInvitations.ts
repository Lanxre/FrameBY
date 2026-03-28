import { $api } from "@/composables/api/useApi";
import type { EmploymentApplication } from "@/types/frontend/employment";

export function useJobInvitations() {
	const invitations = ref<EmploymentApplication[]>([]);
	const total = ref(0);
	const isLoading = ref(false);
	const errorMessage = ref("");

	const fetchInvitations = async (
		options: { status?: string; limit?: number; offset?: number } = {},
	) => {
		isLoading.value = true;
		errorMessage.value = "";

		try {
			const params = new URLSearchParams();
			if (options.status) params.append("status", options.status);
			if (options.limit) params.append("limit", options.limit.toString());
			if (options.offset) params.append("offset", options.offset.toString());

			const query = params.toString() ? `?${params.toString()}` : "";

			const response = await $api<{
				requests: EmploymentApplication[];
				total: number;
			}>(`/employment-requests/my-applications${query}`, {
				method: "GET",
			});

			invitations.value = response.requests;
			total.value = response.total || response.requests.length;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка загрузки заявок";
		} finally {
			isLoading.value = false;
		}
	};

	const cancelInvitation = async (requestId: string): Promise<boolean> => {
		isLoading.value = true;
		errorMessage.value = "";

		try {
			await $api("/employment-requests/cancel", {
				method: "POST",
				body: { request_id: requestId },
			});

			invitations.value = invitations.value.filter(
				(inv) => inv.id !== requestId,
			);
			total.value--;
			return true;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка отмены заявки";
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	return {
		invitations,
		total,
		isLoading,
		errorMessage,
		fetchInvitations,
		cancelInvitation,
	};
}
