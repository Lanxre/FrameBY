import { $api } from "@/composables/api/useApi";
import type { Chat } from "@/types/frontend/chat";

export function useCreateChat() {
	const isLoading = ref(false);
	const error = ref("");

	const createChat = async (userId: string): Promise<Chat | null> => {
		isLoading.value = true;
		error.value = "";
		try {
			const response = await $api<{ id: string; type: string; name?: string }>("/chats", {
				method: "POST",
				body: {
					type: "direct",
					users: [userId],
				},
			});
			return {
				id: response.id,
				type: response.type,
				name: response.name,
				created_at: new Date().toISOString(),
			};
		} catch (e: any) {
			error.value = e.message || "Ошибка создания чата";
			return null;
		} finally {
			isLoading.value = false;
		}
	};

	return {
		isLoading,
		error,
		createChat,
	};
}
