import { $api } from "@/composables/api/useApi";

export function useJoinSquad() {
	const isLoading = ref(false);
	const errorMessage = ref("");
	const isSuccess = ref(false);

	const join = async (squadId: string): Promise<boolean> => {
		isLoading.value = true;
		errorMessage.value = "";
		isSuccess.value = false;

		try {
			await $api(`/student-squads/join`, {
				method: "POST",
				body: { squad_id: squadId },
			});
			isSuccess.value = true;
			return true;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка вступления в отряд";
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	const reset = () => {
		errorMessage.value = "";
		isSuccess.value = false;
	};

	return {
		isLoading,
		errorMessage,
		isSuccess,
		join,
		reset,
	};
}

export function useLeaveSquad() {
	const isLoading = ref(false);
	const errorMessage = ref("");
	const isSuccess = ref(false);

	const leave = async (squadId: string): Promise<boolean> => {
		isLoading.value = true;
		errorMessage.value = "";
		isSuccess.value = false;

		try {
			await $api(`/student-squads/leave`, {
				method: "POST",
				body: { squad_id: squadId },
			});
			isSuccess.value = true;
			return true;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка выхода из отряда";
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	const reset = () => {
		errorMessage.value = "";
		isSuccess.value = false;
	};

	return {
		isLoading,
		errorMessage,
		isSuccess,
		leave,
		reset,
	};
}

export function useApproveSquad() {
	const isLoading = ref(false);
	const errorMessage = ref("");
	const isSuccess = ref(false);

	const approve = async (squadId: string): Promise<boolean> => {
		isLoading.value = true;
		errorMessage.value = "";
		isSuccess.value = false;

		try {
			await $api(`/student-squads/${squadId}/approve`, {
				method: "POST",
				body: { approved: true },
			});
			isSuccess.value = true;
			return true;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка одобрения отряда";
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	const reject = async (squadId: string): Promise<boolean> => {
		isLoading.value = true;
		errorMessage.value = "";
		isSuccess.value = false;

		try {
			await $api(`/student-squads/${squadId}/approve`, {
				method: "POST",
				body: { approved: false },
			});
			isSuccess.value = true;
			return true;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка отклонения отряда";
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	const reset = () => {
		errorMessage.value = "";
		isSuccess.value = false;
	};

	return {
		isLoading,
		errorMessage,
		isSuccess,
		approve,
		reject,
		reset,
	};
}

export function useUpdateSquad() {
	const isLoading = ref(false);
	const errorMessage = ref("");
	const isSuccess = ref(false);

	const update = async (
		squadId: string,
		data: {
			title?: string;
			description?: string;
			profile?: string;
			max_participants?: number;
			status?: string;
		},
	): Promise<boolean> => {
		isLoading.value = true;
		errorMessage.value = "";
		isSuccess.value = false;

		try {
			await $api(`/student-squads/${squadId}`, {
				method: "PATCH",
				body: data,
			});
			isSuccess.value = true;
			return true;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка обновления отряда";
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	const closeRecruitment = async (squadId: string): Promise<boolean> => {
		return update(squadId, { status: "closed" });
	};

	const openRecruitment = async (squadId: string): Promise<boolean> => {
		return update(squadId, { status: "recruitment_open" });
	};

	const reset = () => {
		errorMessage.value = "";
		isSuccess.value = false;
	};

	return {
		isLoading,
		errorMessage,
		isSuccess,
		update,
		closeRecruitment,
		openRecruitment,
		reset,
	};
}
