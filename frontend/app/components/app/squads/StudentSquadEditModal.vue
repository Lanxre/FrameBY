<script setup lang="ts">
import type { StudentSquad } from '@/types/frontend/student-squad'
import { useUpdateSquad } from '@/composables/api/squads/useUpdateSquad'

const props = defineProps<{
  modelValue: boolean
  squad: StudentSquad | null
}>()

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
  saved: [data: { id: string; title: string; description: string | null; profile: string | null; max_participants: number }]
}>()

const { update, isLoading, errorMessage, reset } = useUpdateSquad()

const form = ref({
  title: '',
  description: '',
  profile: '',
  max_participants: 1
})

watch(() => props.squad, (squad) => {
  if (squad) {
    form.value = {
      title: squad.title || '',
      description: squad.description || '',
      profile: squad.profile || '',
      max_participants: squad.max_participants || 1
    }
  }
}, { immediate: true })

watch(() => props.modelValue, (open) => {
  if (!open) {
    reset()
    errorMessage.value = ''
  }
})

const close = () => {
  emit('update:modelValue', false)
}

const handleSubmit = async () => {
  if (!props.squad) return

  const success = await update(props.squad.id, {
    title: form.value.title || undefined,
    description: form.value.description || undefined,
    profile: form.value.profile || undefined,
    max_participants: form.value.max_participants || undefined
  })

  if (success) {
    emit('saved', {
      id: props.squad!.id,
      title: form.value.title,
      description: form.value.description || null,
      profile: form.value.profile || null,
      max_participants: form.value.max_participants
    })
    close()
  }
}
</script>

<template>
  <Transition name="fade">
    <div v-if="modelValue" class="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div class="absolute inset-0 bg-black/40 rounded-2xl backdrop-blur-sm" @click="close"></div>

      <div class="relative w-full max-w-lg bg-white rounded-2xl shadow-xl p-6">
        <div class="flex items-center justify-between mb-6">
          <div class="flex items-center gap-2">
            <Icon name="ph:pencil-simple" size="22" class="text-emerald-500" />
            <h2 class="text-xl font-semibold text-gray-800">Редактирование отряда</h2>
          </div>
          <button @click="close" class="p-1 cursor-pointer text-gray-400 hover:text-gray-600 transition">
            <Icon name="ph:x" size="20" />
          </button>
        </div>

        <div v-if="errorMessage" class="mb-4 p-3 rounded-xl bg-red-50 border border-red-200 text-sm text-red-500">
          {{ errorMessage }}
        </div>

        <form @submit.prevent="handleSubmit" class="space-y-4">
          <div class="space-y-1">
            <label class="text-sm font-medium text-gray-700 ml-1">Название <span class="text-red-500">*</span></label>
            <div class="relative">
              <Icon name="ph:flag" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
              <input
                v-model="form.title"
                type="text"
                required
                class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm
                       bg-white border border-emerald-100
                       focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
                placeholder="Название отряда"
              />
            </div>
          </div>

          <div class="space-y-1">
            <label class="text-sm font-medium text-gray-700 ml-1">Количество участников <span class="text-red-500">*</span></label> 
            <div class="relative">
              <Icon name="ph:users" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
              <input
                v-model.number="form.max_participants"
                type="number"
                required
                min="1"
                class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm
                       bg-white border border-emerald-100
                       focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
              />
            </div>
          </div>

          <div class="space-y-1">
            <label class="text-sm font-medium text-gray-700 ml-1">Профиль работ</label>
            <div class="relative">
              <Icon name="ph:briefcase" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
              <input
                v-model="form.profile"
                type="text"
                class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm
                       bg-white border border-emerald-100
                       focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
                placeholder="Например: Информационные работы, Благоустройство"
              />
            </div>
          </div>

          <div class="space-y-1">
            <label class="text-sm font-medium text-gray-700 ml-1">Описание</label>
            <div class="relative">
              <Icon name="ph:text-align-left" size="18" class="absolute left-3 top-3 text-gray-400" />
              <textarea
                v-model="form.description"
                rows="4"
                class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm resize-none
                       bg-white border border-emerald-100
                       focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
                placeholder="Опишите задачи и условия"
              ></textarea>
            </div>
          </div>

          <div class="flex gap-3 pt-2">
            <button
              type="submit"
              :disabled="isLoading"
              class="flex-1 px-4 py-2 rounded-xl bg-emerald-500
              text-white font-medium
              hover:bg-emerald-600 
              disabled:opacity-50
              transition flex
              items-center 
              justify-center gap-2
              cursor-pointer"
            >
              <Icon v-if="isLoading" name="ph:circle-notch" size="18" class="animate-spin" />
              <span>{{ isLoading ? 'Сохранение...' : 'Сохранить' }}</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </Transition>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
