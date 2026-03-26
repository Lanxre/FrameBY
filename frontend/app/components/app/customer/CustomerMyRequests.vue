<script setup lang="ts">
import { useMySquads, useUpdateSquad } from '@/composables/api/squads/useStudentSquads'
import { useMyOrganizationRequests } from '@/composables/api/employment/useEmploymentRequests'
import { useUpdateEmploymentStatus } from '@/composables/api/employment/useEmploymentActions'
import StudentSquadCard from '@/components/app/squads/StudentSquadCard.vue'
import StudentSquadEditModal from '@/components/app/squads/StudentSquadEditModal.vue'
import EmploymentRequestCard from './EmploymentRequestCard.vue'
import EmploymentRequestEditModal from './EmploymentRequestEditModal.vue'
import ModalConfirm from '@/components/common/ModalConfirm.vue'
import type { StudentSquad } from '@/types/frontend/student-squad'
import type { EmploymentRequest } from '@/types/frontend/employment'

const { squads, isLoading: isLoadingSquads, fetchMySquads, errorMessage: squadsError } = useMySquads()
const { requests, isLoading: isLoadingRequests, fetchRequests, errorMessage: requestsError } = useMyOrganizationRequests()
const { closeRecruitment, openRecruitment, isLoading: isUpdatingSquad } = useUpdateSquad()
const { updateStatus, isLoading: isUpdatingRequest } = useUpdateEmploymentStatus()

const errorMessage = computed(() => squadsError.value || requestsError.value)
const isLoading = computed(() => isLoadingSquads.value || isLoadingRequests.value)
const isUpdating = computed(() => isUpdatingSquad.value || isUpdatingRequest.value)

const activeTab = ref<'all' | 'squads' | 'employment'>('all')

const filteredSquads = computed(() => {
  if (activeTab.value === 'employment') return []
  return squads.value
})

const filteredRequests = computed(() => {
  if (activeTab.value === 'squads') return []
  return requests.value
})

const showCloseModal = ref(false)
const closeItem = ref<{ type: 'squad' | 'request'; item: StudentSquad | EmploymentRequest } | null>(null)

const showSquadEditModal = ref(false)
const editingSquad = ref<StudentSquad | null>(null)

const showRequestEditModal = ref(false)
const editingRequest = ref<EmploymentRequest | null>(null)

const openCloseModal = (type: 'squad' | 'request', item: StudentSquad | EmploymentRequest) => {
  closeItem.value = { type, item }
  showCloseModal.value = true
}

const handleOpenSquad = async (squadId: string) => {
  const success = await openRecruitment(squadId)
  if (success) {
    const squad = squads.value.find(s => s.id === squadId)
    if (squad) squad.status = 'pending'
  }
}

const handleCloseSquad = async (squadId: string) => {
  const squad = squads.value.find(s => s.id === squadId)
  if (squad) {
    openCloseModal('squad', squad)
  }
}

const handleEditSquad = (squad: StudentSquad) => {
  editingSquad.value = squad
  showSquadEditModal.value = true
}

const handleOpenRequest = async (request: EmploymentRequest) => {
  const success = await updateStatus(request.id, 'pending')
  if (success) {
    const req = requests.value.find(r => r.id === request.id)
    if (req) req.status = 'pending'
  }
}

const handleCloseRequest = async (request: EmploymentRequest) => {
  openCloseModal('request', request)
}

const handleEditRequest = (request: EmploymentRequest) => {
  editingRequest.value = request
  showRequestEditModal.value = true
}

const confirmClose = async () => {
  if (!closeItem.value) return
  if (closeItem.value.type === 'squad') {
    const success = await closeRecruitment(closeItem.value.item.id)
    if (success) {
      const squad = squads.value.find(s => s.id === closeItem.value!.item.id)
      if (squad) squad.status = 'closed'
    }
  } else {
    const success = await updateStatus(closeItem.value.item.id, 'closed')
    if (success) {
      const request = requests.value.find(r => r.id === closeItem.value!.item.id)
      if (request) request.status = 'closed'
    }
  }
  showCloseModal.value = false
  closeItem.value = null
}

const handleSaved = (data: EmploymentRequest) => {
  const job = requests.value.find((job) => job.id === data.id)
  if (job) {
    job.title = data.title
    job.description = data.description
    job.requirements = data.requirements
    job.salary = data.salary
    job.schedule = data.schedule
    job.max_participants = data.max_participants
  }
  showRequestEditModal.value = false
  editingRequest.value = null
}

const handleSquadSaved = (data: StudentSquad) => {
  const squad = squads.value.find(s => s.id === data.id)
  if (squad) {
    squad.title = data.title
    squad.description = data.description
    squad.profile = data.profile
    squad.max_participants = data.max_participants
  }
  showSquadEditModal.value = false
  editingSquad.value = null
}

onMounted(async () => {
  await Promise.all([fetchMySquads(), fetchRequests()])
})
</script>

<template>
  <div class="space-y-4">
    <div class="flex items-center justify-between">
      <div class="flex items-center gap-2">
        <Icon name="ph:list-bullets" size="22" class="text-emerald-500" />
        <h2 class="text-lg font-semibold text-gray-800">Мои заявки</h2>
      </div>
      <span class="text-sm text-gray-500">
        {{ filteredSquads.length + filteredRequests.length }} заявок
      </span>
    </div>

    <div v-if="errorMessage" class="p-4 bg-red-50 border border-red-200 rounded-lg text-red-600 text-sm">
      {{ errorMessage }}
    </div>

    <div v-if="requestsError" class="p-4 bg-red-50 border border-red-200 rounded-lg text-red-600 text-sm">
      Ошибка загрузки заявок на работу: {{ requestsError }}
    </div>

    <div class="flex gap-2 border-b border-gray-200">
      <button
        v-for="tab in [
          { value: 'all', label: 'Все' },
          { value: 'squads', label: 'Отряды' },
          { value: 'employment', label: 'Работа' }
        ]"
        :key="tab.value"
        @click="activeTab = tab.value as any"
        :class="[
          'px-4 py-2 text-sm cursor-pointer font-medium border-b-2 -mb-px transition-colors',
          activeTab === tab.value
            ? 'border-emerald-500 text-emerald-600'
            : 'border-transparent text-gray-500 hover:text-gray-700'
        ]"
      >
        {{ tab.label }}
      </button>
    </div>

    <div v-if="isLoading" class="flex justify-center py-8">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
    </div>

    <div v-else-if="filteredSquads.length === 0 && filteredRequests.length === 0" class="text-center py-8 text-gray-500">
      <Icon name="ph:folder-open" size="48" class="mx-auto mb-2 text-gray-300" />
      <p>Заявки не найдены</p>
    </div>

    <div v-else class="space-y-4">
      <template v-for="squad in filteredSquads" :key="squad.id">
        <StudentSquadCard
          :squad="squad"
          :on-edit="handleEditSquad"
          :on-close="handleCloseSquad"
          :on-open="handleOpenSquad"
        />
      </template>

      <template v-for="request in filteredRequests" :key="request.id">
        <EmploymentRequestCard
          :request="request"
          show-actions
          @edit="handleEditRequest"
          @close="handleCloseRequest"
          @open="handleOpenRequest"
        />
      </template>
    </div>

    <StudentSquadEditModal
      v-model="showSquadEditModal"
      :squad="editingSquad"
      @saved="(data) => handleSquadSaved(data)"
    />

    <EmploymentRequestEditModal
      v-model="showRequestEditModal"
      :request="editingRequest"
      @saved="(data) => handleSaved(data)"
    />

    <ModalConfirm
      v-model="showCloseModal"
      title="Закрыть заявку"
      :description="`Закрыть заявку '${('title' in (closeItem?.item ?? {})) ? (closeItem?.item as any).title : ''}'?`"
      confirm-text="Закрыть"
      :loading="isUpdating"
      @confirm="confirmClose"
    />
  </div>
</template>
