<script setup lang="ts">
import EmploymentRequestCard from "./EmploymentRequestCard.vue";
import { useMyOrganizationRequests } from "@/composables/api/employment/useEmploymentRequests";
import { useUpdateEmploymentStatus } from "@/composables/api/employment/useEmploymentActions";
import ModalConfirm from "@/components/common/ModalConfirm.vue";
import type { EmploymentRequest } from "@/types/frontend/employment";

const { requests, isLoading, errorMessage, fetchRequests } =
	useMyOrganizationRequests();
const { updateStatus, isLoading: isUpdating } = useUpdateEmploymentStatus();

const showCloseModal = ref(false);
const requestToClose = ref<EmploymentRequest | null>(null);

const openCloseModal = (request: EmploymentRequest) => {
	requestToClose.value = request;
	showCloseModal.value = true;
};

const openRecruitment = async (request: EmploymentRequest) => {
	await updateStatus(request.id, "approved");
	await fetchRequests();
};

const confirmClose = async () => {
	if (!requestToClose.value) return;
	await updateStatus(requestToClose.value.id, "closed");
	showCloseModal.value = false;
	requestToClose.value = null;
	await fetchRequests();
};

onMounted(fetchRequests);
</script>

<template>
  <div class="space-y-4">
    <div class="flex items-center justify-between">
      <div class="flex items-center gap-2">
        <Icon name="ph:list-bullets" size="22" class="text-emerald-500" />
        <h2 class="text-lg font-semibold text-gray-800">Мои заявки на работу</h2>
      </div>
      <span class="text-sm text-gray-500">Всего: {{ requests.length }}</span>
    </div>

    <div v-if="errorMessage" class="p-4 bg-red-50 border border-red-200 rounded-lg text-red-600 text-sm">
      {{ errorMessage }}
    </div>

    <div v-if="isLoading && requests.length === 0" class="flex justify-center py-8">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
    </div>

    <div v-else-if="requests.length === 0" class="text-center py-8 text-gray-500">
      <Icon name="ph:briefcase" size="48" class="mx-auto mb-2 text-gray-300" />
      <p>Вы пока не создали ни одной заявки</p>
    </div>

    <div v-else class="space-y-3">
      <EmploymentRequestCard
        v-for="request in requests"
        :key="request.id"
        :request="request"
        show-actions
        @close="openCloseModal"
        @open="openRecruitment"
      />
    </div>

    <ModalConfirm
      v-model="showCloseModal"
      title="Закрыть заявку"
      :description="`Закрыть заявку '${requestToClose?.title}'?`"
      confirm-text="Закрыть"
      :loading="isUpdating"
      @confirm="confirmClose"
    />
  </div>
</template>
