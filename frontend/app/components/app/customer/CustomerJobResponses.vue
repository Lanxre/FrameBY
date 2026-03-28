<script setup lang="ts">
import { useVirtualList } from "@vueuse/core";
import ParticipantCard from "./ParticipantCard.vue";
import Pagination from "@/components/ui/Pagination/Pagination.vue";
import StudentInfoModal from "@/components/common/StudentInfoModal.vue";
import { useMyOrganizationRequests } from "@/composables/api/employment/useEmploymentRequests";
import { $api } from "@/composables/api/useApi";
import type { EmploymentParticipant, EmploymentRequest } from "@/types/frontend/employment";

const {
	requests,
	isLoading: isLoadingRequests,
	fetchRequests,
} = useMyOrganizationRequests();

const page = ref(1);
const limit = ref(10);

const participantsByRequest = ref<Map<string, EmploymentParticipant[]>>(new Map());

const isLoading = computed(() => isLoadingRequests.value);

const loadParticipants = async (requestId: string) => {
	try {
		const response = await $api<{
			participants: EmploymentParticipant[];
		}>(`/employment-requests/${requestId}/participants`, {
			method: "GET",
		});
		participantsByRequest.value.set(requestId, response.participants || []);
	} catch (e) {
		participantsByRequest.value.set(requestId, []);
	}
};

const loadAllParticipants = async () => {
	for (const request of requests.value) {
		await loadParticipants(request.id);
	}
};

const flatParticipants = computed(() => {
	const result: Array<{
		participant: EmploymentParticipant;
		request: EmploymentRequest;
	}> = [];

	for (const request of requests.value) {
		const participants = participantsByRequest.value.get(request.id) || [];
		for (const participant of participants) {
			result.push({
				participant,
				request,
			});
		}
	}

	return result;
});

const total = computed(() => flatParticipants.value.length);

const paginatedParticipants = computed(() => {
	const start = (page.value - 1) * limit.value;
	const end = start + limit.value;
	return flatParticipants.value.slice(start, end);
});

const handleStatusUpdated = (participantId: string, newStatus: string) => {
	for (const [requestId, participants] of participantsByRequest.value) {
		const idx = participants.findIndex((p) => p.id === participantId);
		if (idx !== -1) {
			const updatedParticipants = [...participants];
			updatedParticipants[idx] = {
				...updatedParticipants[idx],
				status: newStatus,
				status_id: newStatus === "accepted" ? 2 : 3,
			};
			participantsByRequest.value.set(requestId, updatedParticipants);
			break;
		}
	}
};

const handlePageChange = (newPage: number) => {
	page.value = newPage;
};

const showStudentModal = ref(false);
const selectedParticipant = ref<EmploymentParticipant | null>(null);

const handleViewStudent = (participant: EmploymentParticipant) => {
	selectedParticipant.value = participant;
	showStudentModal.value = true;
};

onMounted(async () => {
	await fetchRequests();
	await loadAllParticipants();
});

const { list, containerProps, wrapperProps } = useVirtualList(paginatedParticipants, {
	itemHeight: 320,
	overscan: 5,
});
</script>

<template>
	<div class="space-y-6">
		<div class="flex flex-wrap items-center justify-between gap-4">
			<span class="text-sm text-gray-500">Всего: {{ total }}</span>
		</div>

		<div v-if="isLoading && total === 0" class="flex justify-center py-8">
			<div class="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
		</div>

		<div
			v-else-if="total === 0"
			class="flex flex-col items-center justify-center py-12 text-gray-500"
		>
			<Icon name="ph:users" size="48" class="mb-3 text-gray-300" />
			<p>Откликов пока нет</p>
		</div>

		<div
			v-else
			v-bind="containerProps"
			class="h-170 overflow-y-auto rounded-2xl border border-emerald-100 bg-white/60 backdrop-blur"
		>
			<div v-bind="wrapperProps" class="p-4 space-y-3">
				<ParticipantCard
					v-for="{ data: item } in list"
					:key="item.participant.id"
					:participant="item.participant"
					:request="item.request"
					:onViewStudent="handleViewStudent"
					@statusUpdated="handleStatusUpdated"
				/>
			</div>
		</div>

		<Pagination
			v-if="total > limit"
			:currentPage="page"
			:totalPages="Math.ceil(total / limit)"
			@pageChange="handlePageChange"
		/>

		<StudentInfoModal
			v-model="showStudentModal"
			:participant="selectedParticipant"
		/>
	</div>
</template>
