<script setup lang="ts">
import { useVirtualList } from "@vueuse/core";
import JobInvitationCard from "./JobInvitationCard.vue";
import Pagination from "@/components/ui/Pagination/Pagination.vue";
import Select from "@/components/ui/Select/Select.vue";
import { useJobInvitations } from "@/composables/api/employment/useJobInvitations";

const { notify } = useNotificationStore();

const statusOptions = [
    { id: 'all', name: 'Все' },
    { id: 'applied', name: 'Подана' },
    { id: 'accepted', name: 'Принята' },
    { id: 'rejected', name: 'Отклонена' },
    { id: 'contracted', name: 'Трудоустроен' },
];

const selectedStatus = ref<{ id: string; name: string } | null>(null);
const page = ref(1);
const limit = ref(10);

const {
    invitations,
    total,
    isLoading,
    errorMessage,
    fetchInvitations,
    cancelInvitation,
} = useJobInvitations();

const statusCache = ref<Record<string, any[]>>({});

const loadInvitations = async () => {
    const cacheKey = selectedStatus.value?.id || "all";

    if (statusCache.value[cacheKey] && page.value === 1) {
        const cached = statusCache.value[cacheKey];
        invitations.value = cached;
        return;
    }

    try {
        await fetchInvitations({
            status: selectedStatus.value?.id === "all" ? undefined : selectedStatus.value?.id,
            limit: limit.value,
            offset: (page.value - 1) * limit.value,
        });

        if (page.value === 1) {
            statusCache.value[cacheKey] = [...invitations.value];
        }
    } catch (e) {
        console.error("Error loading invitations:", e);
    }
};

const handlePageChange = (newPage: number) => {
    page.value = newPage;
    loadInvitations();
};

const handleCancel = async (id: string) => {
    const success = await cancelInvitation(id);
    if (success) {
        notify({
            title: "Успех",
            content: "Заявка отменена",
            type: "success",
        });
    } else {
        notify({
            title: "Ошибка",
            content: "Не удалось отменить заявку",
            type: "error",
        });
    }
};

watch(selectedStatus, () => {
    page.value = 1;
    loadInvitations();
});

onMounted(() => {
    if (statusOptions[0]) {
        selectedStatus.value = statusOptions[0];
    }
    loadInvitations();
});

const { list, containerProps, wrapperProps } = useVirtualList(invitations, {
    itemHeight: 280,
    overscan: 5,
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

            <span class="text-sm text-gray-500">Всего: {{ total }}</span>
        </div>

        <div v-if="errorMessage" class="p-3 rounded-xl bg-red-50 border border-red-200 text-sm text-red-500">
            {{ errorMessage }}
        </div>

        <div v-if="isLoading && invitations.length === 0" class="flex justify-center py-8">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
        </div>

        <div
            v-else-if="invitations.length === 0"
            class="flex flex-col items-center justify-center py-12 text-gray-500"
        >
            <Icon name="ph:clipboard-text" size="48" class="mb-3 text-gray-300" />
            <p>Мои заявки не найдены</p>
        </div>

        <div
            v-else
            v-bind="containerProps"
            class="h-170 overflow-y-auto rounded-2xl border border-emerald-100 bg-white/60 backdrop-blur"
        >
            <div v-bind="wrapperProps" class="p-4 space-y-3">
                <JobInvitationCard
                    v-for="{ data: invitation } in list"
                    :key="invitation.id"
                    :invitation="invitation"
                    :isLoading="isLoading"
                    :onCancel="handleCancel"
                />
            </div>
        </div>

        <Pagination
            v-if="total > limit"
            :currentPage="page"
            :totalPages="Math.ceil(total / limit)"
            @pageChange="handlePageChange"
        />
    </div>
</template>
