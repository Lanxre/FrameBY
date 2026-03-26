<script setup lang="ts">
import { SQUAD_STATUS_LABELS, SQUAD_STATUS_COLORS } from '@/const/squad'
import { useStudentSquads } from '@/composables/api/squads/useStudentSquads'
import { useApproveSquad, useUpdateSquad } from '@/composables/api/squads/useJoinLeaveSquad'

const { squads, total, isLoading, errorMessage, fetchSquads } = useStudentSquads()
const { approve, reject, isLoading: isApproving, errorMessage: approveError, reset: resetApprove, isSuccess: approveSuccess } = useApproveSquad()
const { closeRecruitment, isLoading: isClosing, errorMessage: closeError, reset: resetClose, isSuccess: closeSuccess } = useUpdateSquad()

const statusFilter = ref('pending')
const statusOptions = [
  { label: 'Ожидает подтверждения', value: 'pending' },
  { label: 'Все', value: 'all' }
]

const loadSquads = () => {
  fetchSquads({
    status: statusFilter.value === 'all' ? undefined : statusFilter.value,
    limit: 50
  })
}

watch(statusFilter, loadSquads)

onMounted(loadSquads)

const handleApprove = async (squadId: string) => {
  const success = await approve(squadId)
  if (success) {
    const squad = squads.value.find(s => s.id === squadId)
    if (squad) {
      squad.status = 'recruitment_open'
    }
  }
}

const handleReject = async (squadId: string) => {
  const success = await reject(squadId)
  if (success) {
    if (statusFilter.value === 'pending') {
      const index = squads.value.findIndex(s => s.id === squadId)
      if (index !== -1) squads.value.splice(index, 1)
      total.value--
    } else {
      const squad = squads.value.find(s => s.id === squadId)
      if (squad) squad.status = 'rejected'
    }
  }
}

const handleClose = async (squadId: string) => {
  const success = await closeRecruitment(squadId)
  if (success) {
    const squad = squads.value.find(s => s.id === squadId)
    if (squad) squad.status = 'closed'
  }
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

    <div v-else-if="squads.length === 0" class="text-center py-8 text-gray-500">
      <Icon name="ph:users-three" size="48" class="mx-auto mb-2 text-gray-300" />
      <p>Заявки не найдены</p>
    </div>

    <div v-else class="space-y-3">
      <div
        v-for="squad in squads"
        :key="squad.id"
        class="bg-white/80 border border-emerald-100 rounded-xl p-4 shadow-sm hover:shadow-md transition-shadow"
      >
        <div class="flex items-start justify-between gap-4">
          <div class="flex-1 min-w-0">
            <div class="flex items-center gap-2 mb-2">
              <h3 class="font-semibold text-gray-900 truncate">{{ squad.title }}</h3>
              <span
                :class="[
                  'px-2 py-0.5 rounded-full text-xs font-medium',
                  SQUAD_STATUS_COLORS[squad.status]?.bg || 'bg-gray-100',
                  SQUAD_STATUS_COLORS[squad.status]?.text || 'text-gray-600'
                ]"
              >
                {{ SQUAD_STATUS_LABELS[squad.status] || squad.status }}
              </span>
            </div>

            <p v-if="squad.description" class="text-sm text-gray-600 mb-2">
              {{ squad.description }}
            </p>

            <div class="flex flex-wrap items-center gap-3 text-sm text-gray-500">
              <span class="flex items-center gap-1">
                <Icon name="ph:user" size="16" />
                {{ squad.organizer.name }} ({{ squad.organizer.role }})
              </span>
              <span v-if="squad.profile" class="flex items-center gap-1">
                <Icon name="ph:briefcase" size="16" />
                {{ squad.profile }}
              </span>
              <span class="flex items-center gap-1">
                <Icon name="ph:users" size="16" />
                {{ squad.current_count }}/{{ squad.max_participants }}
              </span>
            </div>
          </div>

          <div class="flex-shrink-0 flex gap-2">
            <button
              v-if="squad.status === 'pending'"
              @click="handleApprove(squad.id)"
              :disabled="isUpdating"
              class="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 disabled:bg-emerald-300 text-white text-sm font-medium rounded-lg transition-colors flex items-center gap-2"
            >
              <Icon v-if="!isUpdating" name="ph:check-circle" size="18" />
              <div v-else class="animate-spin rounded-full h-4 w-4 border-b-2 border-white"></div>
              Одобрить
            </button>
            <button
              v-if="squad.status === 'pending'"
              @click="handleReject(squad.id)"
              :disabled="isUpdating"
              class="px-4 py-2 bg-red-50 hover:bg-red-100 disabled:bg-red-50 text-red-600 text-sm font-medium rounded-lg transition-colors flex items-center gap-2"
            >
              <Icon v-if="!isUpdating" name="ph:x-circle" size="18" />
              <div v-else class="animate-spin rounded-full h-4 w-4 border-b-2 border-red-600"></div>
              Отклонить
            </button>
            <button
              v-if="squad.status === 'recruitment_open'"
              @click="handleClose(squad.id)"
              :disabled="isUpdating"
              class="px-4 py-2 bg-red-50 hover:bg-red-100 disabled:bg-red-50 text-red-600 text-sm font-medium rounded-lg transition-colors flex items-center gap-2"
            >
              <Icon v-if="!isUpdating" name="ph:x-circle" size="18" />
              <div v-else class="animate-spin rounded-full h-4 w-4 border-b-2 border-red-600"></div>
              Закрыть набор
            </button>
          </div>
        </div>

        <div v-if="squad.status === 'recruitment_open'" class="mt-3">
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div
              class="bg-emerald-500 h-2 rounded-full transition-all"
              :style="{ width: `${Math.min((squad.current_count / squad.max_participants) * 100, 100)}%` }"
            ></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
