<script setup lang="ts">
import {
	EMPLOYMENT_STATUS_LABELS,
	EMPLOYMENT_STATUS_COLORS,
} from "@/types/frontend/employment";
import type { EmploymentRequest } from "@/types/frontend/employment";
import { formatDate } from "@/utils/str";
import ToolTip from "@/components/ui/ToolTip.vue";

const props = defineProps<{
	request: EmploymentRequest;
	showActions?: boolean;
	isUniversityMode?: boolean;
	isLoading?: boolean;
}>();

const emit = defineEmits<{
	edit: [request: EmploymentRequest];
	close: [request: EmploymentRequest];
	open: [request: EmploymentRequest];
	approve: [request: EmploymentRequest];
	reject: [request: EmploymentRequest];
}>();

const isPending = computed(() => props.request.status === "pending");
const isApproved = computed(() => props.request.status === "approved");
const isClosed = computed(() => props.request.status === "closed");
const isRejected = computed(() => props.request.status === "rejected");
</script>

<template>
  <div class="p-5 rounded-2xl border border-emerald-100 bg-white shadow-lg">
    <div class="flex justify-between gap-4">
      <div class="flex-1">
        <div class="flex items-center gap-2 mb-2">
          <h3 class="font-semibold text-gray-800 truncate">{{ request.title }}</h3>
          <span
            class="px-2 py-0.5 text-xs rounded-full"
            :class="[EMPLOYMENT_STATUS_COLORS[request.status]?.bg, EMPLOYMENT_STATUS_COLORS[request.status]?.text]"
          >
            {{ EMPLOYMENT_STATUS_LABELS[request.status] }}
          </span>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-3 text-sm text-gray-600 mb-3">
          <div class="flex items-center gap-1">
            <Icon name="ph:buildings" size="14" />
            {{ request.enterprise_name }}
          </div>

          <div class="flex items-center gap-1">
            <Icon name="ph:users" size="14" />
            {{ request.current_participants }} / {{ request.max_participants }}
          </div>

          <div class="flex items-center gap-1">
            <Icon name="ph:graduation-cap" size="14" />
            {{ request.university_name }} / {{ request.department_name }}
          </div>

          <div class="flex items-center gap-1">
            <Icon name="ph:calendar" size="14" />
            {{ formatDate(request.created_at) }}
          </div>

          <div v-if="request.salary" class="flex items-center gap-1">
            <Icon name="ph:currency-rub" size="14" />
            {{ request.salary }}
          </div>

          <div v-if="request.schedule" class="flex items-center gap-1">
            <Icon name="ph:clock" size="14" />
            {{ request.schedule }}
          </div>
        </div>

        <div v-if="request.requirements" class="text-xs text-gray-500 mb-2">
          <span class="font-medium">Требования:</span> {{ request.requirements }}
        </div>

        <p v-if="request.description" class="text-sm text-gray-500 line-clamp-2">
          {{ request.description }}
        </p>
      </div>

      <div v-if="showActions" class="flex flex-col gap-2">
        <template v-if="isUniversityMode">
          <ToolTip text="Закрыть">
            <button
              v-if="isApproved"
              @click="emit('close', request)"
              :disabled="isLoading"
              class="p-1.5 text-gray-400 hover:text-red-500 transition rounded-lg cursor-pointer disabled:opacity-50"
            >
              <Icon name="ph:x" size="18" />
            </button>
          </ToolTip>
        </template>

        <template v-else>
          <ToolTip text="Редактировать">
            <button
              @click="emit('edit', request)"
              class="p-1.5 text-gray-400 hover:text-emerald-500 transition rounded-lg cursor-pointer"
            >
              <Icon name="ph:pencil" size="18" />
            </button>
          </ToolTip>

          <ToolTip v-if="isApproved || isPending" text="Закрыть">
            <button
              @click="emit('close', request)"
              class="p-1.5 text-gray-400 hover:text-red-500 transition rounded-lg cursor-pointer"
            >
              <Icon name="ph:x" size="18" />
            </button>
          </ToolTip>

          <ToolTip v-if="isClosed" text="Открыть">
            <button
              @click="emit('open', request)"
              class="p-1.5 text-gray-400 hover:text-green-500 transition rounded-lg cursor-pointer"
            >
              <Icon name="ph:plus" size="18" />
            </button>
          </ToolTip>

          
        </template>
      </div>
    </div>

    <div v-if="isApproved" class="mt-4">
      <div class="w-full h-2 bg-gray-100 rounded-full overflow-hidden">
        <div
          class="h-full bg-emerald-500 transition-all"
          :style="{ width: `${Math.min((request.current_participants / request.max_participants) * 100, 100)}%` }"
        />
      </div>
    </div>

    <div v-if="isPending" class="mt-4 p-3 bg-amber-50 rounded-lg border border-amber-200">
      <div class="flex items-center gap-2 text-amber-700 text-sm">
        <Icon name="ph:hourglass-medium" size="16" />
        <span>Заявка ожидает подтверждения от университета</span>
      </div>
    </div>

    <div v-if="isRejected" class="mt-4 p-3 bg-red-50 rounded-lg border border-red-200">
      <div class="flex items-center gap-2 text-red-700 text-sm">
        <Icon name="ph:x-circle" size="16" />
        <span>Заявка отклонена</span>
      </div>
    </div>

    <div v-if="isClosed" class="mt-4 p-3 bg-gray-50 rounded-lg border border-gray-200">
      <div class="flex items-center gap-2 text-gray-600 text-sm">
        <Icon name="ph:lock" size="16" />
        <span>Набор закрыт</span>
      </div>
    </div>
  </div>
</template>
