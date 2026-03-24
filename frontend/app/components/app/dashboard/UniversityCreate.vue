<script setup lang="ts">
import { useCreateUniversityDepartment } from '@/composables/api/dashboard/useCreateUniversityDepartment'

const emit = defineEmits<{
  created: []
}>()

const { isLoading, errorMessage, isSuccess, create, reset } = useCreateUniversityDepartment()

const universityName = ref('')
const departmentName = ref('')
const address = ref('')

const handleSubmit = async () => {
  if (!universityName.value.trim() || !departmentName.value.trim()) {
    return
  }

  const success = await create({
    university_name: universityName.value.trim(),
    department_name: departmentName.value.trim(),
    address: address.value.trim() || undefined
  })

  if (success) {
    universityName.value = ''
    departmentName.value = ''
    address.value = ''
    emit('created')
    setTimeout(() => {
      reset()
    }, 3000)
  }
}
</script>

<template>
  <div class="p-4 rounded-xl bg-white/60 border border-emerald-100">
    <div class="flex items-center gap-2 mb-4">
      <Icon name="ph:graduation-cap" size="20" class="text-emerald-500" />
      <h3 class="text-sm font-semibold text-gray-700">Добавить университет / кафедру</h3>
    </div>

    <div v-if="errorMessage" class="text-sm text-red-500 bg-red-50 border border-red-200 px-3 py-2 rounded-xl mb-3">
      {{ errorMessage }}
    </div>

    <div v-if="isSuccess" class="text-sm text-green-600 bg-green-50 border border-green-200 px-3 py-2 rounded-xl mb-3">
      Запись успешно создана
    </div>

    <div class="grid grid-rows-3 gap-3">
      <div class="space-y-1">
        <label class="text-xs text-gray-500 ml-2">Университет</label>
        <div class="relative">
          <Icon name="ph:buildings" size="16" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
          <input
            v-model="universityName"
            type="text"
            placeholder="БГУ"
            class="w-full pl-8 pr-3 py-2 rounded-xl text-sm
                   bg-white border border-emerald-100
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
          />
        </div>
      </div>

      <div class="space-y-1">
        <label class="text-xs text-gray-500 ml-2">Кафедра</label>
        <div class="relative">
          <Icon name="ph:book-open" size="16" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
          <input
            v-model="departmentName"
            type="text"
            placeholder="Факультет прикладной математики"
            class="w-full pl-8 pr-3 py-2 rounded-xl text-sm
                   bg-white border border-emerald-100
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
          />
        </div>
      </div>

      <div class="space-y-1">
        <label class="text-xs text-gray-500 ml-2">Адрес</label>
        <div class="relative">
          <Icon name="ph:map-pin" size="16" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
          <input
            v-model="address"
            type="text"
            placeholder="Минск, ул. Богдана Хмельницкого 5"
            class="w-full pl-8 pr-3 py-2 rounded-xl text-sm
                   bg-white border border-emerald-100
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
          />
        </div>
      </div>

      <div></div>

      <button
        class="py-2 mt-2.5 rounded-xl text-sm font-bold
               bg-linear-to-r from-emerald-400 to-green-600 text-white
               hover:opacity-90 transition
               disabled:opacity-50 disabled:cursor-not-allowed
               shadow-sm shadow-emerald-500/20"
        :disabled="isLoading || !universityName.trim() || !departmentName.trim()"
        @click="handleSubmit"
      >
        <span v-if="isLoading">Создание...</span>
        <span v-else>Создать</span>
      </button>
    </div>
  </div>
</template>
