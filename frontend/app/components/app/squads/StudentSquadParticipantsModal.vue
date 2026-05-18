<script setup lang="ts">
import { computed } from "vue";
import type { SquadParticipant } from "@/types/frontend/student-squad";
import ModalWindow from "@/components/common/ModalWindow.vue";
import ToolTip from "@/components/ui/ToolTip.vue";

const props = defineProps<{
  modelValue: boolean;
  participants: SquadParticipant[];
}>();

const emit = defineEmits<{
  (e: "update:modelValue", value: boolean): void;
}>();

const isOpen = computed({
  get: () => props.modelValue,
  set: (val) => emit("update:modelValue", val),
});
</script>

<template>
  <ModalWindow
    v-model="isOpen"
    title="Участники отряда"
    customWidth="800px"
    @close="isOpen = false"
  >
    <div class="space-y-4 cursor-default">
      <div v-if="!participants?.length" class="flex flex-col items-center justify-center text-center text-gray-500 py-12 gap-3">
        <Icon name="ph:users-three-light" size="48" class="text-gray-300" />
        <span class="text-sm">В этом отряде пока нет участников.</span>
      </div>
      
      <div 
            v-else 
            class="grid gap-4 max-h-[60vh] overflow-y-auto pr-2 justify-items-center"
            :class="participants.length === 1 ? 'grid-cols-1' : 'grid-cols-1 sm:grid-cols-2'"
        >
            <div 
            v-for="participant in participants" 
            :key="participant.id"
            class="w-full max-w-sm p-5 rounded-xl border border-emerald-100 bg-white/70 hover:bg-emerald-50/50 transition-colors shadow-sm flex flex-col items-center text-center justify-between gap-4"
            >
            <div class="flex flex-col items-center w-full gap-2">
                <div class="w-12 h-12 rounded-full bg-emerald-50 text-emerald-600 flex items-center justify-center border border-emerald-100/60 shadow-inner">
                    <div class="w-9 h-9 rounded-full
                                bg-linear-to-r from-emerald-400 to-green-600
                                flex items-center justify-center
                                text-white text-sm font-semibold">
                        <img v-if="participant.avatar" :src="formatAvatar(participant.full_name, participant.avatar)" alt="Avatar" class="w-full h-full object-cover rounded-full" />
                        <span v-else>{{ participant.full_name.split(' ').map((word) => word.charAt(0).toUpperCase()).join(' ').slice(0, -1) }}</span>
                    </div>
                </div>
                
                <div class="w-full">
                <h4 class="font-semibold text-gray-800 line-clamp-2 px-2" :title="participant.full_name">
                    {{ participant.full_name }}
                </h4>
                </div>
    
                <ToolTip v-if="participant.grade" text="Средний балл / Оценка диплома" class="mt-1">
                <span class="inline-flex items-center gap-1 px-2.5 py-0.5 rounded-full bg-amber-100 text-amber-700 text-xs font-semibold whitespace-nowrap">
                    <Icon name="ph:star-fill" size="12" />
                    {{ participant.grade.toFixed(1) }}
                </span>
                </ToolTip>
            </div>
            
            <div class="space-y-3 w-full pt-3 border-t border-gray-100 mt-auto flex flex-col items-center">
                <div v-if="participant.specialty" class="flex flex-col items-center gap-1 text-sm text-gray-600 w-full px-2">
                <div class="flex items-center gap-1 text-gray-400 font-medium text-xs uppercase tracking-wider">
                    <Icon name="ph:book-open" size="14" />
                    <span>Специальность</span>
                </div>
                <span class="line-clamp-2 text-gray-700 text-xs leading-relaxed" :title="participant.specialty">
                    {{ participant.specialty }}
                </span>
                </div>
                
                <div v-if="participant.phone" class="flex flex-col items-center gap-1 text-sm text-gray-600 w-full">
                <div class="flex items-center gap-1 text-gray-400 font-medium text-xs uppercase tracking-wider">
                    <Icon name="ph:phone" size="14" />
                    <span>Контакты</span>
                </div>
                <a :href="`tel:${participant.phone}`" class="text-emerald-600 hover:text-emerald-700 hover:underline font-medium transition-colors">
                    {{ participant.phone }}
                </a>
                </div>
            </div>
            </div>
        </div>
    </div>
  </ModalWindow>
</template>