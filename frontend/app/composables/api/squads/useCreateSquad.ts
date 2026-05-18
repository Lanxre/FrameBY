import { $api } from "@/composables/api/useApi";
import type {
	CreateStudentSquadData,
	StudentSquad,
} from "@/types/frontend/student-squad";

export function useCreateSquad() {
	const isLoading = ref(false);
	const errorMessage = ref("");
	const isSuccess = ref(false);
	const { notify } = useNotificationStore();
	const createdSquad = ref<StudentSquad | null>(null);

	const create = async (data: CreateStudentSquadData): Promise<boolean> => {
		isLoading.value = true;
		errorMessage.value = "";
		isSuccess.value = false;

		try {
			const response = await $api<{ squad: StudentSquad }>("/student-squads", {
				method: "POST",
				body: data,
			});
			createdSquad.value = response.squad;
			isSuccess.value = true;
			notify({
				title: "Успех",
				type: "success",
			});
			return true;
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка создания отряда";
			notify({
				title: "Ошибка",
				content: errorMessage.value,
				type: "error",
			});
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	const reset = () => {
		errorMessage.value = "";
		isSuccess.value = false;
		createdSquad.value = null;
	};

	return {
		isLoading,
		errorMessage,
		isSuccess,
		createdSquad,
		create,
		reset,
	};
}
