<script setup lang="ts">
import { useVirtualList } from "@vueuse/core";
import type { StudentSquad } from "@/types/frontend/student-squad";
import { useStudentSquads } from "@/composables/api/squads/useStudentSquads";
import {
	useApproveSquad,
	useUpdateSquad,
} from "@/composables/api/squads/useJoinLeaveSquad";
import Select from "@/components/ui/Select/Select.vue";
import StudentSquadCard from "@/components/app/squads/StudentSquadCard.vue";
import { formatTotalStudentSquads } from "@/utils/str";

const { notify } = useNotificationStore();

const { squads, total, isLoading, errorMessage, fetchSquads } =
	useStudentSquads();

const {
	approve,
	reject,
	isLoading: isApproving,
	errorMessage: approveError,
	reset: resetApprove,
	isSuccess: approveSuccess,
} = useApproveSquad();

const {
	closeRecruitment,
	isLoading: isClosing,
	errorMessage: closeError,
	reset: resetClose,
	isSuccess: closeSuccess,
} = useUpdateSquad();

const statusFilter = ref<{ id: string; name: string } | null>(null);
const statusOptions = [
	{ id: "pending", name: "Ожидает подтверждения" },
	{ id: "all", name: "Все" },
];

const statusCache = ref<
	Record<string, { squads: StudentSquad[]; total: number }>
>({});
const isCacheLoading = ref(false);

const loadSquads = async (status?: string) => {
	const cacheKey = status || "all";

	if (statusCache.value[cacheKey] && status === undefined) {
		const cached = statusCache.value[cacheKey];
		squads.value = cached.squads;
		total.value = cached.total;
		return;
	}

	isCacheLoading.value = true;
	try {
		await fetchSquads({
			status: status === "all" ? undefined : status,
			limit: 50,
		});

		statusCache.value[cacheKey] = {
			squads: [...squads.value],
			total: total.value,
		};
	} finally {
		isCacheLoading.value = false;
	}
};

watch(statusFilter, async (newFilter) => {
	const status = newFilter?.id;
	const cacheKey = status || "all";

	if (statusCache.value[cacheKey]) {
		const cached = statusCache.value[cacheKey];
		squads.value = cached.squads;
		total.value = cached.total;
	} else {
		await loadSquads(status);
	}
});

onMounted(() => {
	statusFilter.value = statusOptions[0];
	loadSquads("pending");
});

const handleApprove = async (squad: StudentSquad) => {
	const success = await approve(squad.id);
	if (success) {
		const item = squads.value.find((s) => s.id === squad.id);
		if (item) item.status = "recruitment_open";

		if (statusFilter.value?.id === "pending") {
			const index = squads.value.findIndex((s) => s.id === squad.id);
			if (index !== -1) squads.value.splice(index, 1);
			total.value--;
		}

		Object.keys(statusCache.value).forEach((key) => {
			const cached = statusCache.value[key];
			const cachedItem = cached.squads.find((s) => s.id === squad.id);
			if (cachedItem) cachedItem.status = "recruitment_open";
		});

		notify({
			title: "Успех",
			content: "Отряд одобрен",
			type: "success",
		});
	} else {
		notify({
			title: "Ошибка",
			content: "Не удалось одобрить отряд",
			type: "error",
		});
	}
};

const handleReject = async (squad: StudentSquad) => {
	const success = await reject(squad.id);
	if (success) {
		if (statusFilter.value?.id === "pending") {
			const index = squads.value.findIndex((s) => s.id === squad.id);
			if (index !== -1) squads.value.splice(index, 1);
			total.value--;
		} else {
			const item = squads.value.find((s) => s.id === squad.id);
			if (item) item.status = "rejected";
		}

		Object.keys(statusCache.value).forEach((key) => {
			const cached = statusCache.value[key];
			const cachedItem = cached.squads.find((s) => s.id === squad.id);
			if (cachedItem) cachedItem.status = "rejected";
		});

		notify({
			title: "Успех",
			content: "Отряд отклонён",
			type: "success",
		});
	} else {
		notify({
			title: "Ошибка",
			content: "Не удалось отклонить отряд",
			type: "error",
		});
	}
};

const handleClose = async (squadId: string) => {
	const success = await closeRecruitment(squadId);
	if (success) {
		const item = squads.value.find((s) => s.id === squadId);
		if (item) item.status = "closed";

		Object.keys(statusCache.value).forEach((key) => {
			const cached = statusCache.value[key];
			const cachedItem = cached.squads.find((s) => s.id === squadId);
			if (cachedItem) cachedItem.status = "closed";
		});

		notify({
			title: "Успех",
			content: "Набор закрыт",
			type: "success",
		});
	} else {
		notify({
			title: "Ошибка",
			content: "Не удалось закрыть набор",
			type: "error",
		});
	}
};

const isUpdating = computed(() => isApproving.value || isClosing.value);
const updateError = computed(() => approveError.value || closeError.value);
const updateSuccess = computed(
	() => approveSuccess.value || closeSuccess.value,
);

const { list, containerProps, wrapperProps } = useVirtualList(squads, {
	itemHeight: 220,
	overscan: 8,
});
</script>

<template>
  <div class="space-y-4">
    <div class="flex items-center justify-between gap-4">
      <div class="flex items-center gap-3">
        <Select
          v-model="statusFilter"
          :options="statusOptions"
          icon="ph:funnel"
          placeholder="Статус"
        />
      </div>
      <span class="text-sm text-gray-500 whitespace-nowrap">
        Всего: {{ formatTotalStudentSquads(total, { oneText: 'заявка', twoFourText: 'заявки', fivePlusText: 'заявок' }) }}
      </span>
    </div>

    <div v-if="errorMessage" class="p-4 bg-red-50 border border-red-200 rounded-lg text-red-600 text-sm">
      {{ errorMessage }}
    </div>

    <div v-if="updateError" class="p-4 bg-red-50 border border-red-200 rounded-lg text-red-600 text-sm">
      {{ updateError }}
    </div>

    <div v-if="updateSuccess" class="p-4 bg-green-50 border border-green-200 rounded-lg text-green-600 text-sm">
      Статус успешно обновлён!
    </div>

    <div v-if="isLoading || isCacheLoading" class="flex justify-center py-8">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
    </div>

    <div v-else-if="squads.length === 0" class="text-center py-8 text-gray-500">
      <Icon name="ph:users-three" size="48" class="mx-auto mb-2 text-gray-300" />
      <p>Заявки не найдены</p>
    </div>

    <div
      v-else
      v-bind="containerProps"
      class="h-150 overflow-y-auto rounded-2xl border border-emerald-100 bg-white/60 backdrop-blur"
    >
      <div v-bind="wrapperProps" class="p-4 space-y-3">
        <StudentSquadCard
          v-for="{ data: squad } in list"
          :key="squad.id"
          :squad="squad"
          :is-loading="isUpdating"
          :is-university-mode="true"
          :on-approve="handleApprove"
          :on-reject="handleReject"
          :on-close="handleClose"
        />
      </div>
    </div>
  </div>
</template>
