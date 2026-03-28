import { $api } from "@/composables/api/useApi";
import type { ChatMessage } from "@/types/frontend/chat";

export function useMessages() {
	const messages = ref<ChatMessage[]>([]);
	const isLoading = ref(false);
	const error = ref("");

	const fetchMessages = async (chatId: string) => {
		isLoading.value = true;
		error.value = "";
		try {
			const response = await $api<{ messages: ChatMessage[]; total: number }>(
				`/chats/${chatId}/messages`,
				{
					method: "GET",
				}
			);
			messages.value = response.messages.sort((a, b) => a.created_at.localeCompare(b.created_at)) || [];
		} catch (e: any) {
			error.value = e.message || "Ошибка загрузки сообщений";
		} finally {
			isLoading.value = false;
		}
	};

	const sendMessage = async (chatId: string, type: string, content: string) => {
		try {
			const response = await $api<ChatMessage>(`/chats/${chatId}/messages`, {
				method: "POST",
				body: { type, content },
			});
			return response;
		} catch (e: any) {
			error.value = e.message || "Ошибка отправки сообщения";
			return null;
		}
	};

	return {
		messages,
		isLoading,
		error,
		fetchMessages,
		sendMessage,
	};
}
