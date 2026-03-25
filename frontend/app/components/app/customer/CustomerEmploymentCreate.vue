<script setup lang="ts">
import type { CreateEmploymentRequestData } from '@/types/employment'
import { useEnterprises } from '@/composables/api/useEnterprises'
import { useUniversityDepartments } from '@/composables/api/useUniversityDepartments'
import { useCreateEmploymentRequest } from '@/composables/api/employment/useCreateEmploymentRequest'

const emit = defineEmits<{
  created: []
}>()

const { enterprises, fetchEnterprises } = useEnterprises()
const { departments, fetchDepartments } = useUniversityDepartments()
const { create, isLoading, errorMessage, isSuccess, reset } = useCreateEmploymentRequest()

const form = ref<CreateEmploymentRequestData>({
  enterprise_id: '',
  university_department_id: '',
  title: '',
  description: '',
  requirements: '',
  salary: '',
  schedule: '',
  location: '',
  employment_type: 'full_time'
})

const employmentTypes = [
  { value: 'full_time', label: 'Полная занятость' },
  { value: 'part_time', label: 'Частичная занятость' },
  { value: 'remote', label: 'Удалённая работа' },
  { value: 'flexible', label: 'Гибкий график' }
]

const handleSubmit = async () => {
  const success = await create(form.value)
  if (success) {
    reset()
    form.value = {
      enterprise_id: '',
      university_department_id: '',
      title: '',
      description: '',
      requirements: '',
      salary: '',
      schedule: '',
      location: '',
      employment_type: 'full_time'
    }
    emit('created')
  }
}

onMounted(() => {
  fetchEnterprises()
  fetchDepartments()
})
</script>

<template>
  <div class="space-y-4">
    <div v-if="errorMessage" class="p-4 bg-red-50 border border-red-200 rounded-lg text-red-600 text-sm">
      {{ errorMessage }}
    </div>

    <div v-if="isSuccess" class="p-4 bg-green-50 border border-green-200 rounded-lg text-green-600 text-sm">
      Заявка успешно создана!
    </div>

    <form @submit.prevent="handleSubmit" class="space-y-4">
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Предприятие *</label>
        <select
          v-model="form.enterprise_id"
          required
          class="w-full px-3 py-2 rounded-lg border border-emerald-200 bg-white/80 focus:outline-none focus:ring-2 focus:ring-emerald-400"
        >
          <option value="" disabled>Выберите предприятие</option>
          <option v-for="e in enterprises" :key="e.id" :value="e.id">{{ e.name }}</option>
        </select>
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Отдел/Кафедра Университета *</label>
        <select
          v-model="form.university_department_id"
          required
          class="w-full px-3 py-2 rounded-lg border border-emerald-200 bg-white/80 focus:outline-none focus:ring-2 focus:ring-emerald-400"
        >
          <option value="" disabled>Выберите отдел/кафедру</option>
          <option v-for="d in departments" :key="d.id" :value="d.id">{{ d.name }}</option>
        </select>
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Название вакансии *</label>
        <input
          v-model="form.title"
          type="text"
          required
          placeholder="Например: Стажер-разработчик"
          class="w-full px-3 py-2 rounded-lg border border-emerald-200 bg-white/80 focus:outline-none focus:ring-2 focus:ring-emerald-400"
        />
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Тип занятости *</label>
        <select
          v-model="form.employment_type"
          required
          class="w-full px-3 py-2 rounded-lg border border-emerald-200 bg-white/80 focus:outline-none focus:ring-2 focus:ring-emerald-400"
        >
          <option v-for="t in employmentTypes" :key="t.value" :value="t.value">{{ t.label }}</option>
        </select>
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Описание</label>
        <textarea
          v-model="form.description"
          rows="3"
          placeholder="Опишите обязанности и условия работы"
          class="w-full px-3 py-2 rounded-lg border border-emerald-200 bg-white/80 focus:outline-none focus:ring-2 focus:ring-emerald-400 resize-none"
        ></textarea>
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Требования</label>
        <textarea
          v-model="form.requirements"
          rows="2"
          placeholder="Укажите необходимые навыки и квалификацию"
          class="w-full px-3 py-2 rounded-lg border border-emerald-200 bg-white/80 focus:outline-none focus:ring-2 focus:ring-emerald-400 resize-none"
        ></textarea>
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Зарплата</label>
          <input
            v-model="form.salary"
            type="text"
            placeholder="Например: 50 000 руб."
            class="w-full px-3 py-2 rounded-lg border border-emerald-200 bg-white/80 focus:outline-none focus:ring-2 focus:ring-emerald-400"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">График</label>
          <input
            v-model="form.schedule"
            type="text"
            placeholder="Например: 5/2, 8 часов"
            class="w-full px-3 py-2 rounded-lg border border-emerald-200 bg-white/80 focus:outline-none focus:ring-2 focus:ring-emerald-400"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Место работы</label>
          <input
            v-model="form.location"
            type="text"
            placeholder="Город, адрес"
            class="w-full px-3 py-2 rounded-lg border border-emerald-200 bg-white/80 focus:outline-none focus:ring-2 focus:ring-emerald-400"
          />
        </div>
      </div>

      <div class="pt-2">
        <button
          type="submit"
          :disabled="isLoading"
          class="w-full px-4 py-3 bg-emerald-600 hover:bg-emerald-700 disabled:bg-emerald-300 text-white font-medium rounded-lg transition-colors flex items-center justify-center gap-2"
        >
          <Icon v-if="!isLoading" name="ph:plus-circle" size="20" />
          <div v-else class="animate-spin rounded-full h-5 w-5 border-b-2 border-white"></div>
          Создать заявку
        </button>
      </div>
    </form>
  </div>
</template>
