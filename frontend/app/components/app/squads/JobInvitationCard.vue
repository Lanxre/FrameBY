<script setup lang="ts">
import { PARTICIPANT_STATUS_LABELS, PARTICIPANT_STATUS_COLORS } from "@/types/frontend/employment";
import type { EmploymentApplication } from "@/types/frontend/employment";
import ToolTip from "@/components/ui/ToolTip.vue";

const props = withDefaults(
    defineProps<{
        invitation: EmploymentApplication;
        isLoading?: boolean;
        onCancel?: (id: string) => void;
    }>(),
    {
        isLoading: false,
    },
);

const canCancel = computed(
    () => props.invitation.participant_status === "applied" && props.onCancel,
);

const isApplied = computed(() => props.invitation.participant_status === "applied");
const isAccepted = computed(() => props.invitation.participant_status === "accepted");
const isRejected = computed(() => props.invitation.participant_status === "rejected");
const isContracted = computed(() => props.invitation.participant_status === "contracted");
</script>

<template>
    <div class="p-5 rounded-2xl border border-emerald-100 bg-white shadow-lg">
        <div class="flex items-center justify-between gap-2 mb-3">
            <div class="flex items-center gap-2 min-w-0">
                <h3 class="font-semibold text-gray-800 truncate">
                    {{ invitation.title }}
                </h3>
                <span
                    class="px-2 py-0.5 text-xs rounded-full shrink-0"
                    :class="[
                        PARTICIPANT_STATUS_COLORS[invitation.participant_status]?.bg,
                        PARTICIPANT_STATUS_COLORS[invitation.participant_status]?.text,
                    ]"
                >
                    {{ PARTICIPANT_STATUS_LABELS[invitation.participant_status] }}
                </span>
            </div>
            <div class="flex items-center gap-1 shrink-0">
                <ToolTip text="Отменить заявку">
                    <button
                        v-if="canCancel"
                        @click="onCancel?.(invitation.id)"
                        :disabled="isLoading"
                        class="p-2 rounded-xl text-gray-500 flex items-center justify-center disabled:opacity-50 cursor-pointer hover:text-red-600"
                    >
                        <Icon
                            name="ph:x"
                            size="18"
                        />
                    </button>
                </ToolTip>
            </div>
        </div>

        <div class="space-y-2 text-sm text-gray-600 mb-3">
            <div class="flex items-start gap-2">
                <Icon name="ph:building-office" size="14" class="text-gray-400 mt-0.5 shrink-0" />
                <div class="min-w-0">
                    <span class="text-emerald-600 font-medium">{{ invitation.enterprise_name }}</span>
                    <span v-if="invitation.enterprise_address" class="text-gray-500 text-xs block truncate">
                        {{ invitation.enterprise_address }}
                    </span>
                </div>
            </div>

            <div class="flex items-start gap-2">
                <Icon name="ph:graduation-cap" size="14" class="text-gray-400 mt-0.5 shrink-0" />
                <div class="min-w-0">
                    <span class="text-gray-700">{{ invitation.university_name }}</span>
                    <span class="text-gray-500 text-xs block truncate">
                        {{ invitation.department_name }}
                    </span>
                </div>
            </div>

            <div v-if="invitation.salary" class="flex items-center gap-2">
                <Icon name="ph:currency-rub" size="14" class="text-gray-400" />
                <span class="text-emerald-600 font-medium">{{ invitation.salary }}</span>
            </div>

            <div v-if="invitation.schedule" class="flex items-center gap-2">
                <Icon name="ph:calendar" size="14" class="text-gray-400" />
                <span>{{ invitation.schedule }}</span>
            </div>
        </div>

        <p
            v-if="invitation.description"
            class="text-sm text-gray-500 line-clamp-2 mb-3"
        >
            {{ invitation.description }}
        </p>

        <div v-if="invitation.requirements" class="text-sm text-gray-600 mb-3">
            <div class="flex items-start gap-2">
                <Icon name="ph:list-checks" size="14" class="text-gray-400 mt-0.5 shrink-0" />
                <span class="line-clamp-2">{{ invitation.requirements }}</span>
            </div>
        </div>

        <div class="flex items-center gap-2 text-sm text-gray-600 mb-3">
            <Icon name="ph:users" size="14" class="text-gray-400" />
            <span>
                Мест: {{ invitation.current_participants }} / {{ invitation.max_participants }}
            </span>
        </div>

        <div class="text-xs text-gray-400 mb-3">
            Подана: {{ new Date(invitation.participant_applied_at).toLocaleDateString('ru-RU') }}
        </div>

        <div v-if="isAccepted" class="mt-3 p-3 bg-green-50 rounded-lg border border-green-200">
            <div class="flex items-center gap-2 text-green-700 text-sm">
                <Icon name="ph:check-circle" size="16" />
                <span>Заявка принята организацией</span>
            </div>
        </div>

        <div v-if="isRejected" class="mt-3 p-3 bg-red-50 rounded-lg border border-red-200">
            <div class="flex items-center gap-2 text-red-700 text-sm">
                <Icon name="ph:x-circle" size="16" />
                <span>Заявка отклонена организацией</span>
            </div>
        </div>

        <div v-if="isContracted" class="mt-3 p-3 bg-emerald-50 rounded-lg border border-emerald-200">
            <div class="flex items-center gap-2 text-emerald-700 text-sm">
                <Icon name="ph:handshake" size="16" />
                <span>Трудоустроен</span>
            </div>
        </div>

        <div v-if="isApplied" class="mt-4 p-3 cursor-default bg-blue-50 rounded-lg border border-emerald-200">
            <div class="flex items-center gap-2 text-emerald-700 text-sm">
                <Icon name="ph:hourglass-medium" size="16" />
                <span>Ожидает решения организации</span>
            </div>
        </div>
    </div>
</template>
