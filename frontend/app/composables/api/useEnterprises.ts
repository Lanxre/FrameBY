import { ref } from "vue";
import { $api } from "@/composables/api/useApi";

export function useEnterprises() {
	const enterprises = ref<{ id: string; name: string }[]>([]);
	const isLoading = ref(false);
	const errorMessage = ref("");

	const fetchEnterprises = async () => {
		isLoading.value = true;
		errorMessage.value = "";

		try {
			const res: { enterprises: any[] } = await $api("/enterprises");
			enterprises.value = res.enterprises.map((e: any) => ({
				id: e.id,
				name: e.name,
			}));
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка загрузки данных";
		} finally {
			isLoading.value = false;
		}
	};

	return {
		enterprises,
		isLoading,
		errorMessage,
		fetchEnterprises,
	};
}
