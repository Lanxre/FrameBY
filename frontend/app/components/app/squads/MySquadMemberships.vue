<script setup lang="ts">
import { SQUAD_STATUS_LABELS, SQUAD_STATUS_COLORS } from '~/types/frontend/student-squad'
import { useMySquads } from '@/composables/api/squads/useStudentSquads'
import { useLeaveSquad } from '@/composables/api/squads/useJoinLeaveSquad'

const { squads, isLoading, errorMessage, fetchMySquads } = useMySquads()
const { leave, isLoading: isLeaving, errorMessage: leaveError, reset: resetLeave, isSuccess: leaveSuccess } = useLeaveSquad()

const loadSquads = () => {
  fetchMySquads()
}

onMounted(loadSquads)

const handleLeave = async (squadId: string) => {
  const success = await leave(squadId)
  if (success) {
    loadSquads()
  }
}
</script>

<template>
  <div class="space-y-4">
    <div v-if="errorMessage" class="p-4 bg-red-50 border border-red-200 rounded-lg text-red-600 text-sm">
      {{ errorMessage }}
    </div>

    <div v-if="leaveError" class="p-4 bg-red-50 border border-red-200 rounded-lg text-red-600 text-sm">
      {{ leaveError }}
    </div>

    <div v-if="leaveSuccess" class="p-4 bg-green-50 border border-green-200 rounded-lg text-green-600 text-sm">
      Вы успешно покинули отряд!
    </div>

    <div v-if="isLoading" class="flex justify-center py-8">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
    </div>

    <div v-else-if="squads.length === 0" class="text-center py-8 text-gray-500">
      <Icon name="ph:users-three" size="48" class="mx-auto mb-2 text-gray-300" />
      <p>Вы пока не состоите в отрядах</p>
    </div>

    <div v-else class="space-y-3">
      <div
        v-for="squad in squads"
        :key="squad.id"
        class="bg-white/80 border border-emerald-100 rounded-xl p-4 shadow-sm"
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

            <p v-if="squad.description" class="text-sm text-gray-600 mb-2 line-clamp-2">
              {{ squad.description }}
            </p>

            <div class="flex flex-wrap items-center gap-3 text-sm text-gray-500">
              <span class="flex items-center gap-1">
                <Icon name="ph:user" size="16" />
                {{ squad.organizer_name }}
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

          <div class="flex-shrink-0">
            <button
              v-if="squad.status === 'recruitment_open'"
              @click="handleLeave(squad.id)"
              :disabled="isLeaving"
              class="px-4 py-2 bg-red-50 hover:bg-red-100 disabled:bg-red-50 text-red-600 text-sm font-medium rounded-lg transition-colors flex items-center gap-2"
            >
              <Icon v-if="!isLeaving" name="ph:user-minus" size="18" />
              <div v-else class="animate-spin rounded-full h-4 w-4 border-b-2 border-red-600"></div>
              Покинуть
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
