<script setup lang="ts">
import { useVirtualList } from "@vueuse/core";
import type { EmploymentRequest } from "@/types/frontend/employment";
import { useEmploymentRequests } from "@/composables/api/employment/useEmploymentRequests";
import { useUpdateEmploymentStatus } from "@/composables/api/employment/useEmploymentActions";
import { formatTotalStudentSquads } from "@/utils/str";
import Select from "@/components/ui/Select/Select.vue";
import EmploymentRequestCard from "@/components/app/customer/EmploymentRequestCard.vue";

const { requests, total, isLoading, errorMessage, fetchRequests } =
	useEmploymentRequests();

const {
	updateStatus,
	isLoading: isClosing,
	errorMessage: closeError,
	reset: resetClose,
	isSuccess: closeSuccess,
} = useUpdateEmploymentStatus();

const { notify } = useNotificationStore();

const statusFilter = ref<{ id: string; name: string } | null>(null);
const statusOptions = [
	{ id: "all", name: "Все" },
];

const statusCache = ref<
	Record<string, { requests: EmploymentRequest[]; total: number }>
>({});
const isCacheLoading = ref(false);

const loadRequests = async (status?: string) => {
	const cacheKey = status || "all";

	if (statusCache.value[cacheKey] && status === undefined) {
		const cached = statusCache.value[cacheKey];
		requests.value = cached.requests;
		total.value = cached.total;
		return;
	}

	isCacheLoading.value = true;
	try {
		await fetchRequests({
			status: status === "all" ? undefined : status,
			limit: 50,
		});

		statusCache.value[cacheKey] = {
			requests: [...requests.value],
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
		requests.value = cached.requests;
		total.value = cached.total;
	} else {
		await loadRequests(status);
	}
});

onMounted(() => {
	if (statusOptions[0]) statusFilter.value = statusOptions[0];

	loadRequests();
});

const handleClose = async (request: EmploymentRequest) => {
	const success = await updateStatus(request.id, "closed");
	if (success) {
		const item = requests.value.find((r) => r.id === request.id);
		if (item) item.status = "closed";

		Object.keys(statusCache.value).forEach((key) => {
			const cached = statusCache.value[key];
			const cachedItem = cached.requests.find((r) => r.id === request.id);
			if (cachedItem) cachedItem.status = "closed";
		});

		notify({
			title: "Успех",
			content: "Заявка закрыта",
			type: "success",
		});
	} else {
		notify({
			title: "Ошибка",
			content: "Не удалось закрыть заявку",
			type: "error",
		});
	}
};

const isUpdating = computed(() => isClosing.value);
const updateError = computed(() => closeError.value);
const updateSuccess = computed(() => closeSuccess.value);

const { list, containerProps, wrapperProps } = useVirtualList(requests, {
	itemHeight: 200,
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

    <div v-else-if="requests.length === 0" class="text-center py-8 text-gray-500">
      <Icon name="ph:briefcase" size="48" class="mx-auto mb-2 text-gray-300" />
      <p>Заявки не найдены</p>
    </div>

    <div
      v-else
      v-bind="containerProps"
      class="h-150 overflow-y-auto rounded-2xl border border-emerald-100 bg-white/60 backdrop-blur"
    >
      <div v-bind="wrapperProps" class="p-4 space-y-3">
        <EmploymentRequestCard
          v-for="{ data: request } in list"
          :key="request.id"
          :request="request"
          :show-actions="true"
          :is-university-mode="true"
          :is-loading="isUpdating"
          @close="handleClose"
        />
      </div>
    </div>
  </div>
</template>
