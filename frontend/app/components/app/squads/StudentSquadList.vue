<script setup lang="ts">
import { useVirtualList } from '@vueuse/core'
import SquadCard from './SquadCard.vue'
import Select from '@/components/ui/Select/Select.vue'
import Pagination from '@/components/ui/Pagination/Pagination.vue'

import { useStudentSquads, useMySquads } from '@/composables/api/squads/useStudentSquads'
import { useJoinSquad } from '@/composables/api/squads/useJoinLeaveSquad'

const statusOptions = [
  { id: 'all', name: 'Все' },
  { id: 'pending', name: 'Ожидает' },
  { id: 'recruitment_open', name: 'Набор' },
  { id: 'closed', name: 'Закрыта' }
]

const showOnlyMine = ref(false)
const selectedStatus = ref(statusOptions[0])
const page = ref(1)
const limit = ref(10)

const { squads, total, isLoading, errorMessage, fetchSquads } = useStudentSquads()
const { squads: mySquads, fetchMySquads } = useMySquads()
const { join, isLoading: isJoining, errorMessage: joinError, reset: resetJoin, isSuccess: joinSuccess } = useJoinSquad()

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
  const ok = await join(id)
  if (ok) {
    resetJoin()
    await loadSquads()
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

    <div v-if="joinSuccess" class="p-3 rounded-xl bg-emerald-50 border border-emerald-200 text-sm text-emerald-600">
      Вы успешно вступили в отряд
    </div>

    <div
      v-else
      v-bind="containerProps"
      class="h-[600px] overflow-y-auto rounded-2xl border border-emerald-100 bg-white/60 backdrop-blur"
    >
      <div v-bind="wrapperProps" class="p-4 space-y-3">
        <SquadCard
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