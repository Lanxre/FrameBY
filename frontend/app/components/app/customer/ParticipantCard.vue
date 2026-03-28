<script setup lang="ts">
import {
	PARTICIPANT_STATUS_LABELS,
	PARTICIPANT_STATUS_COLORS,
} from "@/types/frontend/employment";
import { useUpdateParticipantStatus } from "@/composables/api/employment/useEmploymentActions";
import type { EmploymentParticipant, EmploymentRequest } from "@/types/frontend/employment";

const props = withDefaults(
	defineProps<{
		participant: EmploymentParticipant;
		request: EmploymentRequest;
		onViewStudent?: (participant: EmploymentParticipant) => void;
	}>(),
	{
		isLoading: false,
	}
);

const emit = defineEmits<{
	(e: "statusUpdated", participantId: string, newStatus: string): void;
}>();

const { notify } = useNotificationStore();
const { updateStatus, isLoading } = useUpdateParticipantStatus();

const isApplied = computed(() => props.participant.status === "applied");
const isAccepted = computed(() => props.participant.status === "accepted");
const isRejected = computed(() => props.participant.status === "rejected");
const isContracted = computed(() => props.participant.status === "contracted");

const statusColorClass = computed(() => {
	const colors = PARTICIPANT_STATUS_COLORS[props.participant.status];
	return colors ? `${colors.bg} ${colors.text}` : "bg-gray-100 text-gray-600";
});

const handleCardClick = () => {
	props.onViewStudent?.(props.participant);
};

const handleApprove = async (e: Event) => {
	e.stopPropagation();
	const success = await updateStatus(props.request.id, props.participant.user_id, "accepted");
	if (success) {
		emit("statusUpdated", props.participant.id, "accepted");
		notify({
			title: "Успех",
			content: "Заявка принята",
			type: "success",
		});
	}
};

const handleReject = async (e: Event) => {
	e.stopPropagation();
	const success = await updateStatus(props.request.id, props.participant.user_id, "rejected");
	if (success) {
		emit("statusUpdated", props.participant.id, "rejected");
		notify({
			title: "Успех",
			content: "Заявка отклонена",
			type: "success",
		});
	}
};
</script>

<template>
	<div
		@click="handleCardClick"
		class="p-5 rounded-2xl border bg-white shadow-lg border-emerald-200 cursor-pointer"
	>
		<div class="flex items-start justify-between gap-2 mb-4">
			<div class="flex items-start gap-3 min-w-0 flex-1">
				<div class="w-10 h-10 rounded-full bg-gradient-to-br from-emerald-400 to-green-500 flex items-center justify-center shrink-0">
					<Icon name="ph:user" size="20" class="text-white" />
				</div>
				<div class="min-w-0 flex-1">
					<div class="flex items-center gap-2 flex-wrap">
						<span class="font-semibold text-gray-800">
							{{ participant.student_name }}
						</span>
						<span
							class="px-2 py-0.5 text-xs rounded-full shrink-0"
							:class="statusColorClass"
						>
							{{ PARTICIPANT_STATUS_LABELS[participant.status] }}
						</span>
					</div>
					<p class="text-xs text-gray-500 mt-0.5">
						Отклик от {{ new Date(participant.applied_at).toLocaleDateString("ru-RU") }}
					</p>
				</div>
			</div>
			<Icon name="ph:caret-right" size="16" class="text-gray-400 shrink-0 mt-2" />
		</div>

		<div class="mb-4 p-4 rounded-xl bg-gradient-to-r from-emerald-50/50 to-green-50/50 border border-emerald-100">
			<h4 class="font-semibold text-gray-800 mb-2 flex items-center gap-2">
				<Icon name="ph:briefcase" size="16" class="text-emerald-600" />
				{{ request.title }}
			</h4>
			<div class="space-y-2">
				<div v-if="request.enterprise_name" class="flex items-center gap-2 text-sm">
					<Icon name="ph:buildings" size="14" class="text-gray-400" />
					<span class="text-gray-600">{{ request.enterprise_name }}</span>
				</div>
				<div v-if="request.salary" class="flex items-center gap-2 text-sm">
					<Icon name="ph:currency-dollar" size="14" class="text-gray-400" />
					<span class="text-gray-600">{{ request.salary }}</span>
				</div>
				<div v-if="request.schedule" class="flex items-center gap-2 text-sm">
					<Icon name="ph:clock" size="14" class="text-gray-400" />
					<span class="text-gray-600">{{ request.schedule }}</span>
				</div>
				<div class="flex items-center gap-2 text-sm">
					<Icon name="ph:users" size="14" class="text-gray-400" />
					<span class="text-gray-600">
						{{ request.current_participants }}/{{ request.max_participants }} мест
					</span>
				</div>
			</div>
		</div>

		<div v-if="request.description || request.requirements" class="mb-4 space-y-2">
			<div v-if="request.description" class="text-sm">
				<p class="text-xs text-gray-500 mb-1">Описание:</p>
				<p class="text-gray-600 line-clamp-2">{{ request.description }}</p>
			</div>
			<div v-if="request.requirements" class="text-sm">
				<p class="text-xs text-gray-500 mb-1">Требования:</p>
				<p class="text-gray-600 line-clamp-2">{{ request.requirements }}</p>
			</div>
		</div>

		<div v-if="isAccepted" class="p-3 bg-green-50 rounded-lg border border-green-200">
			<div class="flex items-center gap-2 text-green-700 text-sm">
				<Icon name="ph:check-circle" size="16" />
				<span>Заявка принята</span>
			</div>
		</div>

		<div v-if="isRejected" class="p-3 bg-red-50 rounded-lg border border-red-200">
			<div class="flex items-center gap-2 text-red-700 text-sm">
				<Icon name="ph:x-circle" size="16" />
				<span>Заявка отклонена</span>
			</div>
		</div>

		<div v-if="isContracted" class="p-3 bg-emerald-50 rounded-lg border border-emerald-200">
			<div class="flex items-center gap-2 text-emerald-700 text-sm">
				<Icon name="ph:handshake" size="16" />
				<span>Трудоустроен</span>
			</div>
		</div>

		<div v-if="isApplied" class="flex gap-2 mt-4">
			<button
				@click="handleApprove"
				:disabled="isLoading"
				class="flex-1 px-4 py-2.5 rounded-xl text-sm font-semibold bg-gradient-to-r from-emerald-400 to-green-600 text-white flex items-center justify-center gap-2 hover:opacity-90 disabled:opacity-50 cursor-pointer transition-opacity"
			>
				<Icon v-if="!isLoading" name="ph:check" size="18" />
				<div v-else class="animate-spin rounded-full h-4 w-4 border-2 border-white border-t-transparent"></div>
				Принять
			</button>
			<button
				@click="handleReject"
				:disabled="isLoading"
				class="flex-1 px-4 py-2.5 rounded-xl text-sm font-medium bg-red-50 hover:bg-red-100 disabled:opacity-50 text-red-600 flex items-center justify-center gap-2 cursor-pointer transition-colors"
			>
				<Icon v-if="!isLoading" name="ph:x" size="18" />
				<div v-else class="animate-spin rounded-full h-4 w-4 border-2 border-red-600 border-t-transparent"></div>
				Отклонить
			</button>
		</div>
	</div>
</template>
