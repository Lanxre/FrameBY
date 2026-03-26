<script setup lang="ts">
import { useVirtualList } from '@vueuse/core'
import StudentSquadCard from './StudentSquadCard.vue'
import Select from '@/components/ui/Select/Select.vue'
import Pagination from '@/components/ui/Pagination/Pagination.vue'
import { SQUAD_STATUS_OPTIONS } from '@/const/squad'

import { useStudentSquads, useMySquads } from '@/composables/api/squads/useStudentSquads'
import { useJoinSquad } from '@/composables/api/squads/useJoinLeaveSquad'
import { FramebyAppRole } from '~/types/frontend/enums/role'

const { user } = useAuthStore()


const statusOptions = computed(() => {
  if (user?.role === FramebyAppRole.STUDENT) {
    return SQUAD_STATUS_OPTIONS.filter(option => option.id !== 'closed' && option.id !== 'rejected')
  }
  
  return SQUAD_STATUS_OPTIONS
})

const showOnlyMine = ref(false)
const selectedStatus = ref(statusOptions.value[0]!)
const page = ref(1)
const limit = ref(10)

const { squads, total, errorMessage, fetchSquads } = useStudentSquads()
const { squads: mySquads, fetchMySquads } = useMySquads()
const { join, isLoading: isJoining, errorMessage: joinError, reset: resetJoin } = useJoinSquad()

const { notify } = useNotificationStore()

const sourceList = ref<any[]>([])

watch([squads, mySquads, showOnlyMine], () => {
  sourceList.value = showOnlyMine.value ? mySquads.value : squads.value
}, { immediate: true })

const loadSquads = async () => {
  if (showOnlyMine.value) {
    await fetchMySquads()
  } else {
    await fetchSquads({
      status: selectedStatus.value.id === 'all' ? undefined : selectedStatus.value.id,
      limit: limit.value,
      offset: (page.value - 1) * limit.value
    })
  }
}

const handlePageChange = (newPage: number) => {
  page.value = newPage
  loadSquads()
}

const handleJoin = async (id: string) => {
  const success = await join(id)
  if (success) {
    resetJoin()
    notify({
      title: 'Успех',
      content: 'Вы успешно присоединились к отряду!',
      type: 'success'
    })
    squads.value = squads.value.filter(squad => squad.id !== id)
  } else {
    notify({
      title: 'Ошибка',
      content: 'Не удалось присоединиться к отряду!',
      type: 'error'
    })
  }
}

watch([selectedStatus, showOnlyMine], () => {
  page.value = 1
  loadSquads()
})

onMounted(loadSquads)

const { list, containerProps, wrapperProps } = useVirtualList(sourceList, {
  itemHeight: 140,
  overscan: 8
})
</script>

<template>
  <div class="space-y-6">

    <div class="flex flex-wrap items-center justify-between gap-4">

      <div class="flex items-center gap-3">
        <Select
          v-model="selectedStatus"
          :options="statusOptions"
          placeholder="Фильтр"
          icon="ph:funnel"
          class="w-40"
        />
      </div>

      <span class="text-sm text-gray-500">{{ total }} отрядов</span>
    </div>

    <div v-if="errorMessage" class="p-3 rounded-xl bg-red-50 border border-red-200 text-sm text-red-500">
      {{ errorMessage }}
    </div>

    <div v-if="joinError" class="p-3 rounded-xl bg-red-50 border border-red-200 text-sm text-red-500">
      {{ joinError }}
    </div>

    <div
      v-bind="containerProps"
      class="h-150 overflow-y-auto rounded-2xl border border-emerald-100 bg-white/60 backdrop-blur"
    >
      <div v-bind="wrapperProps" class="p-4 space-y-3">
        <StudentSquadCard
          v-for="{ data: squad } in list"
          :key="squad.id"
          :squad="squad"
          :isJoining="isJoining"
          :onJoin="handleJoin"
        />
      </div>
    </div>

    <Pagination
      v-if="!showOnlyMine && total > limit"
      :currentPage="page"
      :totalPages="Math.ceil(total / limit)"
      @pageChange="handlePageChange"
    />

  </div>
</template>