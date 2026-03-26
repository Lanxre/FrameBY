<script setup lang="ts">
import { SQUAD_STATUS_LABELS, SQUAD_STATUS_COLORS } from '@/const/squad'
import type { StudentSquad } from '@/types/frontend/student-squad'
import type { FramebyAppRole } from '~/types/frontend/enums/role'
import ToolTip from '@/components/ui/ToolTip.vue';

const props = defineProps<{
  squad: StudentSquad
  isJoining?: boolean
  isLeaving?: boolean
  onJoin?: (id: string) => void
  onLeave?: (id: string) => void
  onEdit?: (squad: StudentSquad) => void
  onClose?: (squadId: string) => void
  onOpen?: (squadId: string) => void
}>()

const canJoin = computed(() =>
  props.squad.status === 'recruitment_open' &&
  props.squad.current_count < props.squad.max_participants &&
  props.onJoin
)

const canLeave = computed(() =>
  props.squad.status === 'recruitment_open' &&
  props.onLeave
)

const canEdit = computed(() => props.onEdit !== undefined)
const canClose = computed(() => props.squad.status !== 'closed' && props.onClose)
const canOpen = computed(() => props.squad.status === 'closed' && props.onOpen)

const isPending = computed(() => props.squad.status === 'pending')
const isRejected = computed(() => props.squad.status === 'rejected')
</script>

<template>
  <div class="p-5 rounded-2xl border border-emerald-100 bg-white shadow-lg">
    <div class="flex items-center justify-between gap-2 mb-3">
     <div class="flex items-center gap-2">
         <h3 class="font-semibold text-gray-800 truncate">{{ squad.title }}</h3>
         <span
           class="px-2 py-0.5 text-xs rounded-full shrink-0"
           :class="[SQUAD_STATUS_COLORS[squad.status]?.bg, SQUAD_STATUS_COLORS[squad.status]?.text]"
         >
           {{ SQUAD_STATUS_LABELS[squad.status] }}
         </span>
     </div>
      <div class="flex items-center justify-center gap-2">
          <ToolTip text="Редактировать">
              <button
                v-if="canEdit"
                @click="onEdit?.(squad)"
                class="px-2 py-2 rounded-xl text-sm font-medium
                       text-gray-600 flex items-center justify-center gap-2 cursor-pointer"
              >
                <Icon name="ph:pencil" size="18" class="hover:text-emerald-600" />
              </button>
          </ToolTip>
    
          <ToolTip text="Закрыть">
              <button
                v-if="canClose"
                @click="onClose?.(squad.id)"
                class="px-2 py-2 rounded-xl text-sm font-medium
                       hover:text-red-600 text-gray-500 flex items-center justify-center gap-2 cursor-pointer"
              >
                <Icon name="ph:x" size="18" />
              </button>
          </ToolTip>
    
          <ToolTip text="Открыть">
              <button
                v-if="canOpen"
                @click="onOpen?.(squad.id)"
                class="px-2 py-2 rounded-xl text-sm font-medium
                       hover:text-green-600 text-gray-600 flex items-center justify-center gap-2 cursor-pointer"
              >
                <Icon name="ph:plus" size="18" />
              </button>
          </ToolTip>
      </div>
    </div>

    <div class="space-y-2 text-sm text-gray-600 mb-3">
      <div class="grid-rows-subgrid col-span-2 mt-4">
        <div class="flex flex-col w-full gap-4">
          <div class="flex gap-1 items-center justify-start text-nowrap">
            <Icon name="ph:user" size="14" />
            {{ formatRole(squad.organizer.role as FramebyAppRole) }}: {{ squad.organizer.name || 'Неизвестно' }}
          </div>
          <div v-if="squad.organizer.position" class="flex gap-1 items-center justify-start text-nowrap">
            <Icon name="ph:briefcase" size="14" />
            Должность: {{ squad.organizer.position }}
          </div>
          <div v-if="squad.organizer.phone" class="flex gap-1 items-center justify-start text-nowrap">
            <Icon name="ph:phone" size="14" />
            Тел: {{ squad.organizer.phone }}
          </div>
        </div>
      </div>
      
      <div v-if="squad.profile" class="flex items-center gap-2">
        <Icon name="ph:briefcase" size="14" class="text-gray-400" />
        <span class="truncate text-emerald-600">{{ squad.profile }}</span>
      </div>
    </div>

    <p v-if="squad.description" class="text-sm text-gray-500 line-clamp-2 mb-3">
      Описание работ: {{ squad.description }}
    </p>
    
    <div class="flex items-center gap-2 text-sm text-gray-600">
      <Icon name="ph:users" size="14" class="text-gray-400" />
      <span> Необходимо: {{ squad.current_count }} / {{ squad.max_participants }}</span>
    </div>
    
    <div v-if="squad.status === 'recruitment_open'" class="mt-3 mb-6">
      <div class="w-full h-2 bg-gray-100 rounded-full overflow-hidden">
        <div
          class="h-full bg-emerald-500 transition-all duration-500"
          :style="{ width: `${Math.min((squad.current_count / squad.max_participants) * 100, 100)}%` }"
        />
      </div>
    </div>

    <div v-if="isPending" class="mt-4 p-3 bg-amber-50 rounded-lg border border-amber-200">
      <div class="flex items-center gap-2 text-amber-700 text-sm">
        <Icon name="ph:hourglass-medium" size="16" />
        <span>Ожидает подтверждения от университета</span>
      </div>
    </div>

    <div v-if="isRejected" class="mt-4 p-3 bg-red-50 rounded-lg border border-red-200">
      <div class="flex items-center gap-2 text-red-700 text-sm">
        <Icon name="ph:x-circle" size="16" />
        <span>Заявка отклонена</span>
      </div>
    </div>
    
    <div class="flex gap-2 mt-4">
      <button
        v-if="canJoin"
        @click="onJoin?.(squad.id)"
        :disabled="isJoining"
        class="flex-1 px-4 py-2 rounded-xl text-sm font-semibold
               bg-linear-to-r from-emerald-400 to-green-600 text-white
               flex items-center justify-center gap-2
               hover:opacity-90 disabled:opacity-50 cursor-pointer"
      >
        <Icon v-if="!isJoining" name="ph:user-plus" size="18" />
        <div v-else class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
        Вступить
      </button>

      <button
        v-if="canLeave"
        @click="onLeave?.(squad.id)"
        :disabled="isLeaving"
        class="flex-1 px-4 py-2 rounded-xl text-sm font-medium
               bg-red-50 hover:bg-red-100 disabled:opacity-50
               text-red-600 flex items-center justify-center gap-2 cursor-pointer"
      >
        <Icon v-if="!isLeaving" name="ph:user-minus" size="18" />
        <div v-else class="w-4 h-4 border-2 border-red-600 border-t-transparent rounded-full animate-spin"></div>
        Покинуть
      </button>

      
    </div>
  </div>
</template>
