<script setup lang="ts">
import { EMPLOYMENT_STATUS_LABELS, EMPLOYMENT_STATUS_COLORS } from '@/types/frontend/employment'
import { useMyEmploymentRequests } from '@/composables/api/employment/useEmploymentRequests'

const emit = defineEmits<{
  refresh: []
}>()

const { requests, isLoading, errorMessage, fetchMyRequests } = useMyEmploymentRequests()

const loadRequests = () => {
  fetchMyRequests()
}

onMounted(loadRequests)

watch(() => emit('refresh'), loadRequests)
</script>

<template>
  <div class="space-y-4">
    <div v-if="errorMessage" class="p-4 bg-red-50 border border-red-200 rounded-lg text-red-600 text-sm">
      {{ errorMessage }}
    </div>

    <div v-if="isLoading" class="flex justify-center py-8">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
    </div>

    <div v-else-if="requests.length === 0" class="text-center py-8 text-gray-500">
      <Icon name="ph:briefcase" size="48" class="mx-auto mb-2 text-gray-300" />
      <p>Вы пока не создали ни одной заявки</p>
    </div>

    <div v-else class="space-y-3">
      <div
        v-for="request in requests"
        :key="request.id"
        class="bg-white/80 border border-emerald-100 rounded-xl p-4 shadow-sm"
      >
        <div class="flex items-start justify-between gap-4">
          <div class="flex-1 min-w-0">
            <div class="flex items-center gap-2 mb-2">
              <h3 class="font-semibold text-gray-900 truncate">{{ request.title }}</h3>
              <span
                :class="[
                  'px-2 py-0.5 rounded-full text-xs font-medium',
                  EMPLOYMENT_STATUS_COLORS[request.status]?.bg || 'bg-gray-100',
                  EMPLOYMENT_STATUS_COLORS[request.status]?.text || 'text-gray-600'
                ]"
              >
                {{ EMPLOYMENT_STATUS_LABELS[request.status] || request.status }}
              </span>
            </div>

            <p v-if="request.description" class="text-sm text-gray-600 mb-2 line-clamp-2">
              {{ request.description }}
            </p>

            <div class="flex flex-wrap items-center gap-3 text-sm text-gray-500">
              <span class="flex items-center gap-1">
                <Icon name="ph:buildings" size="16" />
                {{ request.enterprise_name }}
              </span>
              <span class="flex items-center gap-1">
                <Icon name="ph:graduation-cap" size="16" />
                {{ request.university_department_name }}
              </span>
              <span v-if="request.salary" class="flex items-center gap-1">
                <Icon name="ph:currency-rub" size="16" />
                {{ request.salary }}
              </span>
              <span v-if="request.employment_type" class="flex items-center gap-1">
                <Icon name="ph:clock" size="16" />
                {{ request.employment_type }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
