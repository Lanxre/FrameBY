<script setup lang="ts">
import Select from '@/components/ui/Select/Select.vue'
import ModalConfirm from '@/components/common/ModalConfirm.vue'
import type { CreateEmploymentRequestData } from '@/types/frontend/employment'
import { useUniversityDepartments } from '@/composables/api/useUniversityDepartments'
import { useCreateEmploymentRequest } from '@/composables/api/employment/useCreateEmploymentRequest'
import { useNotificationStore } from '@/stores/notification'

const emit = defineEmits<{
  created: []
}>()

const { departments, fetchDepartments } = useUniversityDepartments()
const { create, isLoading, errorMessage, isSuccess, reset } = useCreateEmploymentRequest()
const notification = useNotificationStore()

const selectedDepartment = ref<{ id: string; name: string } | null>(null)
const showConfirmModal = ref(false)

const form = ref<CreateEmploymentRequestData>({
  university_department_id: '',
  title: '',
  description: '',
  requirements: '',
  salary: '',
  schedule: '',
  max_participants: 10
})

const errors = ref({
  department: '',
  title: '',
  maxParticipants: ''
})

watch(selectedDepartment, (val) => {
  form.value.university_department_id = val?.id || ''
  if (val) errors.value.department = ''
})

watch(() => form.value.title, (val) => {
  if (val?.trim()) errors.value.title = ''
})

watch(() => form.value.max_participants, (val) => {
  if (val > 0) errors.value.maxParticipants = ''
})

const validate = (): boolean => {
  let isValid = true

  if (!selectedDepartment.value) {
    errors.value.department = 'Выберите отдел/кафедру'
    isValid = false
  }

  if (!form.value.title?.trim()) {
    errors.value.title = 'Введите название вакансии'
    isValid = false
  }

  if (form.value.max_participants <= 0) {
    errors.value.maxParticipants = 'Укажите количество участников'
    isValid = false
  }

  return isValid
}

const handleSubmit = async () => {
  if (!validate()) return

  showConfirmModal.value = true
}

const confirmSubmit = async () => {
  showConfirmModal.value = false

  const success = await create(form.value)
  if (success) {
    notification.notify({
      type: 'success',
      title: 'Успешно',
      content: 'Заявка успешно создана!'
    })
    reset()
    selectedDepartment.value = null
    form.value = {
      university_department_id: '',
      title: '',
      description: '',
      requirements: '',
      salary: '',
      schedule: '',
      max_participants: 10
    }
    emit('created')
  } else if (errorMessage.value) {
    notification.notify({
      type: 'error',
      title: 'Ошибка',
      content: errorMessage.value
    })
  }
}

onMounted(() => {
  fetchDepartments()
})
</script>

<template>
  <div class="space-y-4">
    <div class="flex items-center gap-2 mb-4">
      <Icon name="ph:briefcase" size="22" class="text-emerald-500" />
      <h2 class="text-lg font-semibold text-gray-800">Создание заявки на работу</h2>
    </div>

    <form @submit.prevent="handleSubmit" class="space-y-4">
      <div class="space-y-1">
        <label class="text-sm font-medium text-gray-700 ml-1">Отдел/Кафедра Университета <span class="text-red-500">*</span></label>
        <Select
          v-model="selectedDepartment"
          :options="departments"
          placeholder="Выберите отдел/кафедру"
          icon="ph:graduation-cap"
        />
        <p v-if="errors.department" class="text-xs text-red-500 ml-1 mt-1">{{ errors.department }}</p>
      </div>

      <div class="space-y-1">
        <label class="text-sm font-medium text-gray-700 ml-1">Название вакансии <span class="text-red-500">*</span></label>
        <div class="relative">
          <Icon name="ph:text-t" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
          <input
            v-model="form.title"
            type="text"
            placeholder="Например: Стажер-разработчик"
            class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm
                   bg-white border border-emerald-300
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
            :class="{ 'border-red-400': errors.title }"
          />
        </div>
        <p v-if="errors.title" class="text-xs text-red-500 ml-1 mt-1">{{ errors.title }}</p>
      </div>

      <div class="space-y-1">
        <label class="text-sm font-medium text-gray-700 ml-1">Количество участников <span class="text-red-500">*</span></label>
        <div class="relative">
          <Icon name="ph:users" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
          <input
            v-model.number="form.max_participants"
            type="number"
            min="1"
            class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm
                   bg-white border border-emerald-300
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
            :class="{ 'border-red-400': errors.maxParticipants }"
          />
        </div>
        <p v-if="errors.maxParticipants" class="text-xs text-red-500 ml-1 mt-1">{{ errors.maxParticipants }}</p>
      </div>

      <div class="space-y-1">
        <label class="text-sm font-medium text-gray-700 ml-1">Описание</label>
        <div class="relative">
          <Icon name="ph:text-align-left" size="18" class="absolute left-3 top-3 text-gray-400" />
          <textarea
            v-model="form.description"
            rows="3"
            placeholder="Опишите обязанности и условия работы"
            class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm resize-none
                   bg-white border border-emerald-300
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
          ></textarea>
        </div>
      </div>

      <div class="space-y-1">
        <label class="text-sm font-medium text-gray-700 ml-1">Требования</label>
        <div class="relative">
          <Icon name="ph:check-circle" size="18" class="absolute left-3 top-3 text-gray-400" />
          <textarea
            v-model="form.requirements"
            rows="2"
            placeholder="Укажите необходимые навыки и квалификацию"
            class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm resize-none
                   bg-white border border-emerald-300
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
          ></textarea>
        </div>
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div class="space-y-1">
          <label class="text-sm font-medium text-gray-700 ml-1">Зарплата</label>
          <div class="relative">
            <input
              v-model="form.salary"
              type="text"
              class="w-full pl-8 pr-4 py-2.5 rounded-xl text-sm
                     bg-white border border-emerald-300
                     focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
            />
          </div>
        </div>

        <div class="space-y-1">
          <label class="text-sm font-medium text-gray-700 ml-1">График</label>
          <div class="relative">
            <Icon name="ph:calendar-blank" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
            <input
              v-model="form.schedule"
              type="text"
              placeholder="Например: 5/2, 8 часов"
              class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm
                     bg-white border border-emerald-300
                     focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
            />
          </div>
        </div>
      </div>

      <button
        type="submit"
        :disabled="isLoading"
        class="w-full py-2.5 rounded-xl text-sm font-semibold
               bg-linear-to-r from-emerald-400 to-green-600 text-white
               hover:opacity-90 transition
               disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer
               shadow-lg shadow-emerald-500/20 flex items-center justify-center gap-2"
      >
        <Icon v-if="!isLoading" name="ph:plus-circle" size="18" />
        <div v-else class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
        <span>{{ isLoading ? 'Создание...' : 'Создать заявку' }}</span>
      </button>
    </form>

    <ModalConfirm
      v-model="showConfirmModal"
      title="Создание заявки"
      description="Вы уверены, что хотите создать заявку на работу?"
      confirm-text="Создать"
      :loading="isLoading"
      @confirm="confirmSubmit"
    />
  </div>
</template>
