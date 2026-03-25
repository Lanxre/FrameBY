<script setup lang="ts">
import { SQUAD_STATUS_LABELS, SQUAD_STATUS_COLORS } from '~/types/frontend/student-squad'
import { useStudentSquads, useMySquads, useSquadStatuses } from '@/composables/api/squads/useStudentSquads'
import { useUpdateSquad } from '@/composables/api/squads/useUpdateSquad'
import type { StudentSquad } from '~/types/frontend/student-squad'
import { formatDate } from '@/utils/str'
import ModalConfirm from '@/components/common/ModalConfirm.vue'

const emit = defineEmits<{ updated: [] }>()

const { squads, total, isLoading, errorMessage, fetchSquads } = useStudentSquads()
const { squads: mySquads, fetchMySquads } = useMySquads()
const { fetchStatuses } = useSquadStatuses()
const { closeRecruitment, openRecruitment, isLoading: isUpdating } = useUpdateSquad()

const showMySquadsOnly = ref(true)
const showCloseModal = ref(false)
const squadToClose = ref<StudentSquad | null>(null)

const limit = ref(10)
const offset = ref(0)

const totalPages = computed(() => Math.ceil(total.value / limit.value))
const currentPage = computed(() => Math.floor(offset.value / limit.value) + 1)

const isMyView = computed(() => showMySquadsOnly.value)

const displayedSquads = computed(() =>
  isMyView.value ? mySquads.value : squads.value
)

watchEffect(async () => {
  if (isMyView.value) {
    await fetchMySquads()
  } else {
    await fetchSquads({ limit: limit.value, offset: offset.value })
  }
  await fetchStatuses()
})

const goToPage = (page: number) => {
  if (page < 1 || page > totalPages.value) return
  offset.value = (page - 1) * limit.value
}

const openCloseModal = (squad: StudentSquad) => {
  squadToClose.value = squad
  showCloseModal.value = true
}

const confirmClose = async () => {
  if (!squadToClose.value) return
  await closeRecruitment(squadToClose.value.id)
  showCloseModal.value = false
  squadToClose.value = null
  emit('updated')
}

const openRecruitmentFor = async (squad: StudentSquad) => {
  await openRecruitment(squad.id)
  emit('updated')
}
</script>

<template>
  <div class="space-y-5">

    <div class="flex items-center justify-between">
      <div class="flex items-center gap-2">
        <Icon name="ph:list-bullets" size="22" class="text-emerald-500" />
        <h2 class="text-lg font-semibold text-gray-800">
          Список наборов отрядов
        </h2>
      </div>

      <div class="flex items-center gap-4">
        <label class="flex items-center gap-2 cursor-pointer">
          <input type="checkbox" v-model="showMySquadsOnly" class="sr-only peer" />
          <div class="w-10 h-6 bg-gray-200 rounded-full peer-checked:bg-emerald-500 transition relative">
            <div class="absolute top-1 left-1 w-4 h-4 bg-white rounded-full transition peer-checked:translate-x-4"></div>
          </div>
          <span class="text-sm text-gray-600">Только мои</span>
        </label>

        <span class="text-sm text-gray-500">Всего: {{ total }}</span>
      </div>
    </div>

    <div v-if="errorMessage" class="text-sm text-red-600 bg-red-50 border border-red-200 px-4 py-3 rounded-xl">
      {{ errorMessage }}
    </div>

    <div v-if="isLoading" class="flex justify-center py-12">
      <Icon name="ph:circle-notch" size="32" class="animate-spin text-emerald-500" />
    </div>

    <div v-else-if="displayedSquads.length === 0" class="text-center py-12 text-gray-500">
      <Icon name="ph:users-three" size="48" class="mx-auto mb-3 text-gray-300" />
      <p>Отряды не найдены</p>
    </div>

    <div v-else class="space-y-4">
      <div
        v-for="squad in displayedSquads"
        :key="squad.id"
        class="p-5 rounded-2xl border border-emerald-100 bg-white shadow-md"
      >
        <div class="flex justify-between gap-4">

          <div class="flex-1">
            <div class="flex items-center gap-2 mb-2">
              <h3 class="font-semibold text-gray-800">
                {{ squad.title }}
              </h3>
              <span
                class="px-2 py-0.5 text-xs rounded-full"
                :class="[
                  SQUAD_STATUS_COLORS[squad.status]?.bg,
                  SQUAD_STATUS_COLORS[squad.status]?.text
                ]"
              >
                {{ SQUAD_STATUS_LABELS[squad.status] }}
              </span>
            </div>

            <div class="grid grid-cols-2 gap-3 text-sm text-gray-600 mb-3">
              <div class="flex items-center gap-1">
                <Icon name="ph:users" size="14" />
                {{ squad.current_count }} / {{ squad.max_participants }}
              </div>

              <div class="flex items-center gap-1">
                <Icon name="ph:calendar" size="14" />
                {{ formatDate(squad.created_at) }}
              </div>

              <div class="flex items-center gap-1">
                <Icon name="ph:user" size="14" />
                {{ squad.organizer_name || 'Неизвестно' }}
              </div>

              <div v-if="squad.profile" class="truncate col-span-2">
                {{ squad.profile }}
              </div>
            </div>

            <p v-if="squad.description" class="text-sm text-gray-500 line-clamp-2">
              {{ squad.description }}
            </p>

          </div>

          <div v-if="isMyView && squad.status !== 'closed'" class="flex flex-col gap-2">
            <button
              v-if="squad.status === 'recruitment_open'"
              @click="openCloseModal(squad)"
              class="btn-danger"
            >
              Закрыть
            </button>

            <button
              v-if="squad.status === 'closed'"
              @click="openRecruitmentFor(squad)"
              class="btn-success"
            >
              Открыть
            </button>
          </div>

        </div>

        <div class="mt-4">
          <div class="w-full h-2 bg-gray-200 rounded-full">
            <div
              class="h-2 rounded-full bg-emerald-500 transition-all"
              :style="{ width: `${Math.min((squad.current_count / squad.max_participants) * 100, 100)}%` }"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- Pagination -->
    <div v-if="!isMyView && totalPages > 1" class="flex justify-center gap-3 pt-4">
      <button @click="goToPage(currentPage - 1)" :disabled="currentPage === 1" class="page-btn">
        <Icon name="ph:caret-left" />
      </button>

      <span class="text-sm text-gray-600">
        {{ currentPage }} / {{ totalPages }}
      </span>

      <button @click="goToPage(currentPage + 1)" :disabled="currentPage === totalPages" class="page-btn">
        <Icon name="ph:caret-right" />
      </button>
    </div>

    <ModalConfirm
      v-model="showCloseModal"
      title="Закрыть набор"
      :description="`Закрыть набор '${squadToClose?.title}'?`"
      confirmText="Закрыть"
      :loading="isUpdating"
      @confirm="confirmClose"
    />

  </div>
</template>