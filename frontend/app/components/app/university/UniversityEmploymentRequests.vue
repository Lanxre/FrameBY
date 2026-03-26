<script setup lang="ts">
import { EMPLOYMENT_STATUS_LABELS, EMPLOYMENT_STATUS_COLORS } from '@/types/frontend/employment'
import { useEmploymentRequests } from '@/composables/api/employment/useEmploymentRequests'
import { useApproveEmployment, useUpdateEmploymentStatus } from '@/composables/api/employment/useEmploymentActions'

const { requests, total, isLoading, errorMessage, fetchRequests } = useEmploymentRequests()
const { approve, reject, isLoading: isApproving, errorMessage: approveError, reset: resetApprove, isSuccess: approveSuccess } = useApproveEmployment()
const { updateStatus, isLoading: isClosing, errorMessage: closeError, reset: resetClose, isSuccess: closeSuccess } = useUpdateEmploymentStatus()

const statusFilter = ref('pending')
const statusOptions = [
  { label: 'Ожидает подтверждения', value: 'pending' },
  { label: 'Все', value: 'all' }
]

const loadRequests = () => {
  fetchRequests({
    status: statusFilter.value === 'all' ? undefined : statusFilter.value,
    limit: 50
  })
}

watch(statusFilter, loadRequests)

onMounted(loadRequests)

const handleApprove = async (requestId: string) => {
  const success = await approve(requestId)
  if (success) {
    const request = requests.value.find(r => r.id === requestId)
    if (request) request.status = 'approved'
  }
}

const handleReject = async (requestId: string) => {
  const success = await reject(requestId)
  if (success) {
    if (statusFilter.value === 'pending') {
      const index = requests.value.findIndex(r => r.id === requestId)
      if (index !== -1) requests.value.splice(index, 1)
      total.value--
    } else {
      const request = requests.value.find(r => r.id === requestId)
      if (request) request.status = 'rejected'
    }
  }
}

const handleClose = async (requestId: string) => {
  const success = await updateStatus(requestId, 'closed')
  if (success) {
    const request = requests.value.find(r => r.id === requestId)
    if (request) request.status = 'closed'
  }
}

const getStatusColor = (status: string) => {
  return EMPLOYMENT_STATUS_COLORS[status] || { bg: 'bg-gray-100', text: 'text-gray-600' }
}

const isUpdating = computed(() => isApproving.value || isClosing.value)
const updateError = computed(() => approveError.value || closeError.value)
const updateSuccess = computed(() => approveSuccess.value || closeSuccess.value)
const resetUpdate = () => {
  resetApprove()
  resetClose()
}
</script>

<template>
  <div class="space-y-4">
    <div class="flex items-center gap-4">
      <div class="flex items-center gap-2">
        <Icon name="ph:funnel" class="text-emerald-600" />
        <select
          v-model="statusFilter"
          class="px-3 py-2 rounded-lg border border-emerald-200 bg-white/80 text-sm focus:outline-none focus:ring-2 focus:ring-emerald-400"
        >
          <option v-for="opt in statusOptions" :key="opt.value" :value="opt.value">
            {{ opt.label }}
          </option>
        </select>
      </div>
      <span class="text-sm text-gray-500">Всего: {{ total }}</span>
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

    <div v-if="isLoading" class="flex justify-center py-8">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
    </div>

    <div v-else-if="requests.length === 0" class="text-center py-8 text-gray-500">
      <Icon name="ph:briefcase" size="48" class="mx-auto mb-2 text-gray-300" />
      <p>Заявки не найдены</p>
    </div>

    <div v-else class="space-y-3">
      <div
        v-for="request in requests"
        :key="request.id"
        class="bg-white/80 border border-emerald-100 rounded-xl p-4 shadow-sm hover:shadow-md transition-shadow"
      >
        <div class="flex items-start justify-between gap-4">
          <div class="flex-1 min-w-0">
            <div class="flex items-center gap-2 mb-2">
              <h3 class="font-semibold text-gray-900 truncate">{{ request.title }}</h3>
              <span
                :class="[
                  'px-2 py-0.5 rounded-full text-xs font-medium',
                  getStatusColor(request.status).bg,
                  getStatusColor(request.status).text
                ]"
              >
                {{ EMPLOYMENT_STATUS_LABELS[request.status] || request.status }}
              </span>
            </div>

            <p v-if="request.description" class="text-sm text-gray-600 mb-2">
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

            <div v-if="request.requirements" class="mt-2 text-sm text-gray-500">
              <span class="font-medium">Требования:</span> {{ request.requirements }}
            </div>
          </div>

          <div class="flex-shrink-0 flex gap-2">
            <button
              v-if="request.status === 'pending'"
              @click="handleApprove(request.id)"
              :disabled="isUpdating"
              class="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 disabled:bg-emerald-300 text-white text-sm font-medium rounded-lg transition-colors flex items-center gap-2"
            >
              <Icon v-if="!isUpdating" name="ph:check-circle" size="18" />
              <div v-else class="animate-spin rounded-full h-4 w-4 border-b-2 border-white"></div>
              Одобрить
            </button>
            <button
              v-if="request.status === 'pending'"
              @click="handleReject(request.id)"
              :disabled="isUpdating"
              class="px-4 py-2 bg-red-50 hover:bg-red-100 disabled:bg-red-50 text-red-600 text-sm font-medium rounded-lg transition-colors flex items-center gap-2"
            >
              <Icon v-if="!isUpdating" name="ph:x-circle" size="18" />
              <div v-else class="animate-spin rounded-full h-4 w-4 border-b-2 border-red-600"></div>
              Отклонить
            </button>
            <button
              v-if="request.status === 'approved'"
              @click="handleClose(request.id)"
              :disabled="isUpdating"
              class="px-4 py-2 bg-red-50 hover:bg-red-100 disabled:bg-red-50 text-red-600 text-sm font-medium rounded-lg transition-colors flex items-center gap-2"
            >
              <Icon v-if="!isUpdating" name="ph:x-circle" size="18" />
              <div v-else class="animate-spin rounded-full h-4 w-4 border-b-2 border-red-600"></div>
              Закрыть
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
