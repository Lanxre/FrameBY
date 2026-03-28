import { $api } from "@/composables/api/useApi";
import type { Chat } from "@/types/frontend/chat";

export function useChats() {
	const chats = ref<Chat[]>([]);
	const isLoading = ref(false);
	const error = ref("");

	const fetchChats = async () => {
		isLoading.value = true;
		error.value = "";
		try {
			const response = await $api<{ chats: Chat[]; total: number }>("/chats", {
				method: "GET",
			});
			chats.value = response.chats || [];
		} catch (e: any) {
			error.value = e.message || "Ошибка загрузки чатов";
		} finally {
			isLoading.value = false;
		}
	};

	return {
		chats,
		isLoading,
		error,
		fetchChats,
	};
}
