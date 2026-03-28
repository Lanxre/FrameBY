<script setup lang="ts">
import { SQUAD_STATUS_LABELS, SQUAD_STATUS_COLORS } from "@/const/squad";

const props = defineProps<{
	squad: any;
	isJoining: boolean;
	onJoin: (id: string) => void;
}>();

const canJoin = (squad: any) =>
	squad.status === "recruitment_open" &&
	squad.current_count < squad.max_participants;
</script>

<template>
<div class="p-5 rounded-2xl
           bg-white/80 backdrop-blur
           border border-emerald-100
           shadow-md shadow-emerald-500/5
           hover:shadow-lg hover:shadow-emerald-500/10
           transition">

  <div class="flex justify-between gap-4">

    <div class="flex-1 space-y-2">

      <div class="flex items-center gap-2">
        <h3 class="font-semibold text-gray-800 truncate">
          {{ squad.title }}
        </h3>

        <span
          :class="[
            'px-2 py-0.5 rounded-full text-xs font-medium',
            SQUAD_STATUS_COLORS[squad.status]?.bg,
            SQUAD_STATUS_COLORS[squad.status]?.text
          ]"
        >
          {{ SQUAD_STATUS_LABELS[squad.status] }}
        </span>
      </div>

      <p v-if="squad.description"
         class="text-sm text-gray-500 line-clamp-2">
        {{ squad.description }}
      </p>

      <div class="flex flex-wrap gap-3 text-sm text-gray-500">
        <span class="flex items-center gap-1">
          <Icon name="ph:user" size="16" />
          {{ squad.organizer.name }}
        </span>

        <span class="flex items-center gap-1">
          <Icon name="ph:users" size="16" />
          {{ squad.current_count }}/{{ squad.max_participants }}
        </span>
      </div>
    </div>

    <div class="flex items-center">

      <button
        v-if="canJoin(squad)"
        @click="onJoin(squad.id)"
        :disabled="isJoining"
        class="px-4 py-2 rounded-xl text-sm font-semibold
               bg-linear-to-r from-emerald-400 to-green-600 text-white
               flex items-center gap-2
               hover:opacity-90
               disabled:opacity-50"
      >
        <Icon v-if="!isJoining" name="ph:user-plus" size="18" />
        <div v-else class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
        Вступить
      </button>

      <span
        v-else
        class="px-4 py-2 rounded-xl text-sm bg-gray-100 text-gray-500"
      >
        Нет мест
      </span>

    </div>
  </div>

  <div v-if="squad.status === 'recruitment_open'" class="mt-4">
    <div class="h-2 bg-gray-100 rounded-full overflow-hidden">
      <div
        class="h-full bg-gradient-to-r from-emerald-400 to-green-600 transition-all"
        :style="{ width: `${(squad.current_count / squad.max_participants) * 100}%` }"
      />
    </div>
  </div>

</div>
</template>
