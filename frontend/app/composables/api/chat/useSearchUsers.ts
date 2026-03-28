import { $api } from "@/composables/api/useApi";
import type { SearchUser } from "@/types/frontend/chat";

export function useSearchUsers() {
	const results = ref<SearchUser[]>([]);
	const isLoading = ref(false);
	const error = ref("");

	const searchUsers = async (query: string) => {
		if (!query.trim()) {
			results.value = [];
			return;
		}

		isLoading.value = true;
		error.value = "";
		try {
			const response = await $api<{ profiles: SearchUser[]; total: number }>(
				`/users/search?search=${encodeURIComponent(query)}&limit=10`,
				{
					method: "GET",
				}
			);
			results.value = response.profiles || [];
		} catch (e: any) {
			error.value = e.message || "Ошибка поиска";
			results.value = [];
		} finally {
			isLoading.value = false;
		}
	};

	const clearResults = () => {
		results.value = [];
	};

	return {
		results,
		isLoading,
		error,
		searchUsers,
		clearResults,
	};
}
