<script setup lang="ts">
import { useVirtualList } from "@vueuse/core";
import StudentSquadCard from "./StudentSquadCard.vue";
import Select from "@/components/ui/Select/Select.vue";
import Pagination from "@/components/ui/Pagination/Pagination.vue";
import { SQUAD_STATUS_OPTIONS } from "@/const/squad";

import {
	useStudentSquads,
	useMySquads,
} from "@/composables/api/squads/useStudentSquads";
import { useJoinSquad } from "@/composables/api/squads/useJoinLeaveSquad";
import { formatTotalStudentSquads } from "@/utils/str";
import { removeFromFirstByProp } from "@/utils/arr";
import { FramebyAppRole } from "~/types/frontend/enums/role";

const { user } = useAuthStore();

const statusOptions = computed(() => {
	if (user?.role === FramebyAppRole.STUDENT) {
		return SQUAD_STATUS_OPTIONS.filter(
			(option) => option.id !== "closed" && option.id !== "rejected",
		);
	}

	return SQUAD_STATUS_OPTIONS;
});

const showOnlyMine = ref(false);
const selectedStatus = ref<{ id: string; name: string } | null>(null);
const page = ref(1);
const limit = ref(10);

const {
	squads,
	total,
	isLoading: isLoadingSquads,
	fetchSquads,
	errorMessage,
} = useStudentSquads();
const {
	squads: mySquads,
	isLoading: isLoadingMySquads,
	fetchMySquads,
} = useMySquads();
const {
	join,
	isLoading: isJoining,
	errorMessage: joinError,
	reset: resetJoin,
} = useJoinSquad();

const { notify } = useNotificationStore();

const sourceList = ref<any[]>([]);
const isLoading = ref(false);

watch(
	[squads, mySquads, showOnlyMine],
	() => {
		sourceList.value = showOnlyMine.value ? mySquads.value : squads.value;
	},
	{ immediate: true },
);

const statusCache = ref<Record<string, { squads: any[]; total: number }>>({});

const loadSquads = async () => {
	await fetchMySquads();

	const cacheKey = selectedStatus.value?.id || "all";

	if (statusCache.value[cacheKey] && page.value === 1) {
		const cached = statusCache.value[cacheKey];
		squads.value = cached.squads;
		total.value = cached.total;
		return;
	}

	isLoading.value = true;
	try {
		await fetchSquads({
			status:
				selectedStatus.value?.id === "all"
					? undefined
					: selectedStatus.value?.id,
			limit: limit.value,
			offset: (page.value - 1) * limit.value,
		});

		if (page.value === 1) {
			let filteredSquads = squads.value;
			if (mySquads.value.length > 0) {
				filteredSquads = removeFromFirstByProp(
					[...squads.value],
					mySquads.value,
					"id",
				);
			}
			statusCache.value[cacheKey] = {
				squads: filteredSquads,
				total: total.value,
			};
			squads.value = filteredSquads;
		}
	} finally {
		isLoading.value = false;
	}

	if (mySquads.value.length > 0) {
		squads.value = removeFromFirstByProp(squads.value, mySquads.value, "id");
	}
};

const handlePageChange = (newPage: number) => {
	page.value = newPage;
	loadSquads();
};

const handleJoin = async (id: string) => {
	const success = await join(id);
	if (success) {
		resetJoin();
		notify({
			title: "Успех",
			content: "Вы успешно присоединились к отряду!",
			type: "success",
		});

		sourceList.value = sourceList.value.filter((squad) => squad.id !== id);

		if (!showOnlyMine.value) {
			total.value--;
		}

		Object.keys(statusCache.value).forEach((key) => {
			const index = statusCache.value[key].squads.findIndex((s) => s.id === id);
			if (index !== -1) {
				statusCache.value[key].squads.splice(index, 1);
				statusCache.value[key].total--;
			}
		});
	} else {
		notify({
			title: "Ошибка",
			content: "Не удалось присоединиться к отряду!",
			type: "error",
		});
	}
};

watch([selectedStatus, showOnlyMine], () => {
	page.value = 1;
	loadSquads();
});

onMounted(() => {
	selectedStatus.value = statusOptions.value[0] || null;
	loadSquads();
});

const { list, containerProps, wrapperProps } = useVirtualList(sourceList, {
	itemHeight: 140,
	overscan: 8,
});
</script>

<template>
  <div class="space-y-6">

    <div class="flex flex-wrap items-center justify-between gap-4">

      <div class="flex items-center gap-3">
        <Select
          v-model="selectedStatus"
          :options="statusOptions"
          placeholder="Фильтр"
          icon="ph:funnel"
          class="w-40"
        />
      </div>

      <span class="text-sm text-gray-500">Всего: {{ formatTotalStudentSquads(total) }} </span>
    </div>

    <div v-if="errorMessage" class="p-3 rounded-xl bg-red-50 border border-red-200 text-sm text-red-500">
      {{ errorMessage }}
    </div>

    <div v-if="joinError" class="p-3 rounded-xl bg-red-50 border border-red-200 text-sm text-red-500">
      {{ joinError }}
    </div>

    <div v-if="isLoading" class="flex justify-center py-8">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
    </div>

    <div
      v-else-if="sourceList.length === 0"
      class="flex flex-col items-center justify-center py-12 text-gray-500"
    >
      <Icon name="ph:users-three" size="48" class="mb-3 text-gray-300" />
      <p>Отряды не найдены</p>
    </div>

    <div
      v-else
      v-bind="containerProps"
      class="h-170 overflow-y-auto rounded-2xl border border-emerald-100 bg-white/60 backdrop-blur"
    >
      <div v-bind="wrapperProps" class="p-4 space-y-3">
        <StudentSquadCard
          v-for="{ data: squad } in list"
          :key="squad.id"
          :squad="squad"
          :isJoining="isJoining"
          :onJoin="handleJoin"
        />
      </div>
    </div>

    <Pagination
      v-if="!showOnlyMine && total > limit"
      :currentPage="page"
      :totalPages="Math.ceil(total / limit)"
      @pageChange="handlePageChange"
    />

  </div>
</template>
