<script setup lang="ts">
import { useVirtualList } from "@vueuse/core";
import AvailableJobCard from "./AvailableJobCard.vue";
import Pagination from "@/components/ui/Pagination/Pagination.vue";
import { useEmploymentRequests } from "@/composables/api/employment/useEmploymentRequests";
import { useApplyToEmployment } from "@/composables/api/employment/useApplyToEmployment";
import { $api } from "@/composables/api/useApi";

const { notify } = useNotificationStore();

const page = ref(1);
const limit = ref(10);

const {
    requests,
    total,
    isLoading,
    errorMessage,
    fetchRequests,
} = useEmploymentRequests();

const { apply, isLoading: isApplying, errorMessage: applyError, reset: resetApply } = useApplyToEmployment();

const appliedRequests = ref<Set<string>>(new Set());
const filteredRequests = ref<any[]>([]);

const loadAppliedRequests = async () => {
    try {
        const response = await $api<{ requests: any[] }>("/employment-requests/my-applications", {
            method: "GET",
        });
        response.requests.forEach((req: any) => {
            appliedRequests.value.add(req.id);
        });
    } catch (e) {
        console.error("Error loading applied requests:", e);
    }
};

const filterAppliedRequests = () => {
    filteredRequests.value = requests.value.filter(
        (req) => !appliedRequests.value.has(req.id) && req.max_participants !== req.current_participants
    );
};

const loadJobs = async () => {
    try {
        await fetchRequests({
            limit: limit.value,
            offset: (page.value - 1) * limit.value,
        });

        filterAppliedRequests();
    } catch (e) {
        console.error("Error loading jobs:", e);
    }
};

const handlePageChange = (newPage: number) => {
    page.value = newPage;
    loadJobs();
};

const handleApply = async (id: string) => {
    const success = await apply(id);
    if (success) {
        appliedRequests.value.add(id);
        filteredRequests.value = filteredRequests.value.filter((req) => req.id !== id);
        resetApply();
        notify({
            title: "Успех",
            content: "Заявка подана!",
            type: "success",
        });
    } else {
        notify({
            title: "Ошибка",
            content: applyError.value || "Не удалось подать заявку",
            type: "error",
        });
    }
};

onMounted(async () => {
    await loadAppliedRequests();
    loadJobs();
});

const { list, containerProps, wrapperProps } = useVirtualList(filteredRequests, {
    itemHeight: 280,
    overscan: 5,
});
</script>

<template>
    <div class="space-y-6">
        <div class="flex flex-wrap items-center justify-between gap-4">
            <span class="text-sm text-gray-500">Всего: {{ filteredRequests.length }}</span>
        </div>

        <div v-if="errorMessage" class="p-3 rounded-xl bg-red-50 border border-red-200 text-sm text-red-500">
            {{ errorMessage }}
        </div>

        <div v-if="isLoading && filteredRequests.length === 0" class="flex justify-center py-8">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
        </div>

        <div
            v-else-if="filteredRequests.length === 0"
            class="flex flex-col items-center justify-center py-12 text-gray-500"
        >
            <Icon name="ph:briefcase" size="48" class="mb-3 text-gray-300" />
            <p>Нет доступных вакансий</p>
        </div>

        <div
            v-else
            v-bind="containerProps"
            class="h-170 overflow-y-auto rounded-2xl border border-emerald-100 bg-white/60 backdrop-blur"
        >
            <div v-bind="wrapperProps" class="p-4 space-y-3">
                <AvailableJobCard
                    v-for="{ data: job } in list"
                    :key="job.id"
                    :job="job"
                    :isLoading="isApplying"
                    :isApplied="appliedRequests.has(job.id)"
                    :onApply="handleApply"
                />
            </div>
        </div>

        <Pagination
            v-if="total > limit"
            :currentPage="page"
            :totalPages="Math.ceil(filteredRequests.length / limit)"
            @pageChange="handlePageChange"
        />
    </div>
</template>
