<script setup lang="ts">
import Select from "@/components/ui/Select/Select.vue";
import Pagination from "@/components/ui/Pagination/Pagination.vue";
import StudentSquadCard from "@/components/app/squads/StudentSquadCard.vue";
import { useStudentSquads } from "@/composables/api/squads/useStudentSquads";
import { SQUAD_STATUS_OPTIONS } from "@/const/squad";

const statusOptions = SQUAD_STATUS_OPTIONS;
const selectedStatus = ref<{ id: string; name: string } | null>(null);
const limit = ref(12);
const offset = ref(0);

const { squads, total, isLoading, errorMessage, fetchSquads } =
	useStudentSquads();

const totalPages = computed(() => Math.ceil(total.value / limit.value));
const currentPage = computed(() => Math.floor(offset.value / limit.value) + 1);

const statusCache = ref<Record<string, any[]>>({});

const loadSquads = async () => {
	const cacheKey = selectedStatus.value?.id || "all";

	if (statusCache.value[cacheKey] && currentPage.value === 1) {
		const cached = statusCache.value[cacheKey];
		squads.value = cached;
		return;
	}

	await fetchSquads({
		status:
			selectedStatus.value?.id === "all" ? undefined : selectedStatus.value?.id,
		limit: limit.value,
		offset: offset.value,
	});

	if (currentPage.value === 1) {
		statusCache.value[cacheKey] = [...squads.value];
	}
};

watch(selectedStatus, () => {
	offset.value = 0;
	loadSquads();
});

onMounted(() => {
	selectedStatus.value = statusOptions[0] || null;
	loadSquads();
});

const handlePageChange = (page: number) => {
	offset.value = (page - 1) * limit.value;
	loadSquads();
};
</script>

<template>
  <div class="container mx-auto px-4 py-8 max-w-6xl">
    <div class="flex items-center justify-between mb-8">
      <div class="flex items-center gap-3 w-full">
        <Icon name="ph:users-three" size="28" class="text-emerald-500" />
        <h1 class="text-2xl font-bold text-gray-800">Студенческие отряды</h1>
      </div>

      <Select
        v-model="selectedStatus"
        :options="statusOptions"
        class="max-w-32"
        placeholder="Фильтр"
        icon="ph:funnel"
      />
    </div>

    <div v-if="errorMessage" class="p-4 bg-red-50 border border-red-200 rounded-xl text-red-600 mb-6">
      {{ errorMessage }}
    </div>

    <div class="relative" :class="{ 'opacity-50 pointer-events-none': isLoading }">
      <div v-if="isLoading && squads.length === 0" class="flex justify-center py-20">
        <Icon name="ph:circle-notch" size="40" class="animate-spin text-emerald-500" />
      </div>

      <div v-else-if="squads.length === 0" class="text-center py-20 text-gray-500">
        <Icon name="ph:users-three" size="64" class="mx-auto mb-4 text-gray-300" />
        <p class="text-lg">Отряды не найдены</p>
      </div>

      <div v-else class="grid grid-cols-1 xl:grid-cols-2 gap-4 max-w-5xl mx-auto">
        <StudentSquadCard
          v-for="squad in squads"
          :key="squad.id"
          :squad="squad"
        />
      </div>
    </div>

    <div v-if="totalPages > 1" class="mt-8">
      <Pagination
        :currentPage="currentPage"
        :totalPages="totalPages"
        @pageChange="handlePageChange"
      />
    </div>
  </div>
</template>
