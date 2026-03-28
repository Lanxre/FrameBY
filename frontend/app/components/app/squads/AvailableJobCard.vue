<script setup lang="ts">
import type { EmploymentRequest } from "@/types/frontend/employment";

const props = withDefaults(
    defineProps<{
        job: EmploymentRequest;
        isLoading?: boolean;
        isApplied?: boolean;
        onApply?: (id: string) => void;
    }>(),
    {
        isLoading: false,
        isApplied: false,
    },
);

const canApply = computed(
    () => props.job.status === "approved" && 
           props.job.current_participants < props.job.max_participants && 
           props.onApply && 
           !props.isApplied,
);
</script>

<template>
    <div class="p-5 rounded-2xl border border-emerald-100 bg-white shadow-lg">
        <div class="flex items-center justify-between gap-2 mb-3">
            <div class="flex items-center gap-2 min-w-0">
                <h3 class="font-semibold text-gray-800 truncate">
                    {{ job.title }}
                </h3>
            </div>
            <div v-if="isApplied" class="px-2 py-0.5 text-xs rounded-full bg-blue-100 text-blue-600 shrink-0">
                Подана
            </div>
        </div>

        <div class="space-y-2 text-sm text-gray-600 mb-3">
            <div class="flex items-start gap-2">
                <Icon name="ph:building-office" size="14" class="text-gray-400 mt-0.5 shrink-0" />
                <div class="min-w-0">
                    <span class="text-emerald-600 font-medium">{{ job.enterprise_name }}</span>
                    <span v-if="job.enterprise_address" class="text-gray-500 text-xs block truncate">
                        {{ job.enterprise_address }}
                    </span>
                </div>
            </div>

            <div class="flex items-start gap-2">
                <Icon name="ph:graduation-cap" size="14" class="text-gray-400 mt-0.5 shrink-0" />
                <div class="min-w-0">
                    <span class="text-gray-700">{{ job.university_name }}</span>
                    <span class="text-gray-500 text-xs block truncate">
                        {{ job.department_name }}
                    </span>
                </div>
            </div>

            <div v-if="job.salary" class="flex items-center gap-2">
                <Icon name="ph:currency-rub" size="14" class="text-gray-400" />
                <span class="text-emerald-600 font-medium">{{ job.salary }}</span>
            </div>

            <div v-if="job.schedule" class="flex items-center gap-2">
                <Icon name="ph:calendar" size="14" class="text-gray-400" />
                <span>{{ job.schedule }}</span>
            </div>
        </div>

        <p
            v-if="job.description"
            class="text-sm text-gray-500 line-clamp-2 mb-3"
        >
            {{ job.description }}
        </p>

        <div v-if="job.requirements" class="text-sm text-gray-600 mb-3">
            <div class="flex items-start gap-2">
                <Icon name="ph:list-checks" size="14" class="text-gray-400 mt-0.5 shrink-0" />
                <span class="line-clamp-2">{{ job.requirements }}</span>
            </div>
        </div>

        <div class="flex items-center gap-2 text-sm text-gray-600 mb-3">
            <Icon name="ph:users" size="14" class="text-gray-400" />
            <span>
                Мест: {{ job.current_participants }} / {{ job.max_participants }}
            </span>
        </div>

        <div v-if="canApply" class="mt-4 p-3 bg-green-50 rounded-lg border border-green-200">
            <div class="flex items-center gap-2 text-green-700 text-sm">
                <Icon name="ph:check-circle" size="16" />
                <span>Идёт набор</span>
            </div>
        </div>

        <div class="flex gap-2 mt-4">
            <button
                v-if="canApply"
                @click="onApply?.(job.id)"
                :disabled="isLoading"
                class="flex-1 px-4 py-2 rounded-xl text-sm font-semibold bg-linear-to-r from-emerald-400 to-green-600 text-white flex items-center justify-center gap-2 hover:opacity-90 disabled:opacity-50 cursor-pointer"
            >
                <Icon v-if="!isLoading" name="ph:paper-plane-tilt" size="18" />
                <div
                    v-else
                    class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"
                ></div>
                Откликнуться
            </button>

            <button
                v-if="isApplied"
                :disabled="true"
                class="flex-1 px-4 py-2 rounded-xl text-sm font-medium bg-blue-50 text-blue-600 flex items-center justify-center gap-2 cursor-not-allowed"
            >
                <Icon name="ph:check" size="18" />
                Заявка подана
            </button>
        </div>
    </div>
</template>
