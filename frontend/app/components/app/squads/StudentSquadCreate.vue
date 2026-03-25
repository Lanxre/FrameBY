<script setup lang="ts">
import { useCreateSquad } from '@/composables/api/squads/useCreateSquad'

const emit = defineEmits<{
  created: []
}>()

const { isLoading, errorMessage, isSuccess, create, reset } = useCreateSquad()

const title = ref('')
const description = ref('')
const profile = ref('')
const maxParticipants = ref<number>(10)

const handleSubmit = async () => {
  if (!title.value.trim() || maxParticipants.value <= 0) {
    return
  }

  const success = await create({
    title: title.value.trim(),
    description: description.value.trim() || undefined,
    profile: profile.value.trim() || undefined,
    max_participants: maxParticipants.value
  })

  if (success) {
    title.value = ''
    description.value = ''
    profile.value = ''
    maxParticipants.value = 10
    emit('created')
    setTimeout(() => {
      reset()
    }, 3000)
  }
}
</script>

<template>
  <div class="space-y-4">
    <div class="flex items-center gap-2 mb-4">
      <Icon name="ph:users-three" size="22" class="text-emerald-500" />
      <h2 class="text-lg font-semibold text-gray-800">Создание заявки на студенческий отряд</h2>
    </div>

    <div v-if="errorMessage" class="text-sm text-red-500 bg-red-50 border border-red-200 px-4 py-3 rounded-xl">
      {{ errorMessage }}
    </div>

    <div v-if="isSuccess" class="text-sm text-green-600 bg-green-50 border border-green-200 px-4 py-3 rounded-xl">
      Заявка успешно создана! Ожидайте подтверждения от представителя университета.
    </div>

    <form @submit.prevent="handleSubmit" class="space-y-4">
      <div class="space-y-1">
        <label class="text-sm font-medium text-gray-700 ml-1">Название отряда</label>
        <div class="relative">
          <Icon name="ph:flag" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
          <input
            v-model="title"
            type="text"
            placeholder="Например: Отряд по уборке территории"
            class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm
                   bg-white border border-emerald-100
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
          />
        </div>
      </div>

      <div class="space-y-1">
        <label class="text-sm font-medium text-gray-700 ml-1">Описание</label>
        <div class="relative">
          <Icon name="ph:text-align-left" size="18" class="absolute left-3 top-3 text-gray-400" />
          <textarea
            v-model="description"
            rows="3"
            placeholder="Опишите основные задачи и обязанности отряда"
            class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm resize-none
                   bg-white border border-emerald-100
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
          />
        </div>
      </div>

      <div class="space-y-1">
        <label class="text-sm font-medium text-gray-700 ml-1">Профиль работы</label>
        <div class="relative">
          <Icon name="ph:briefcase" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
          <input
            v-model="profile"
            type="text"
            placeholder="Например: Строительные работы, уборка, благоустройство"
            class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm
                   bg-white border border-emerald-100
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
          />
        </div>
      </div>

      <div class="space-y-1">
        <label class="text-sm font-medium text-gray-700 ml-1">Необходимое количество человек</label>
        <div class="relative">
          <Icon name="ph:users" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
          <input
            v-model.number="maxParticipants"
            type="number"
            min="1"
            max="100"
            placeholder="10"
            class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm
                   bg-white border border-emerald-100
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
          />
        </div>
      </div>

      <button
        type="submit"
        :disabled="isLoading || !title.trim() || maxParticipants <= 0"
        class="w-full py-2.5 rounded-xl text-sm font-semibold
               bg-linear-to-r from-emerald-400 to-green-600 text-white
               hover:opacity-90 transition
               disabled:opacity-50 disabled:cursor-not-allowed
               shadow-lg shadow-emerald-500/20"
      >
        <span v-if="isLoading">Создание...</span>
        <span v-else>Создать заявку</span>
      </button>
    </form>
  </div>
</template>
