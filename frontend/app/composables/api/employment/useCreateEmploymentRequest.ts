import { $api } from '@/composables/api/useApi'
import type { CreateEmploymentRequestData, EmploymentRequest } from '@/types/frontend/employment'

export function useCreateEmploymentRequest() {
  const isLoading = ref(false)
  const errorMessage = ref('')
  const isSuccess = ref(false)
  const createdRequest = ref<EmploymentRequest | null>(null)

  const create = async (data: CreateEmploymentRequestData): Promise<boolean> => {
    isLoading.value = true
    errorMessage.value = ''
    isSuccess.value = false

    try {
      const response = await $api<{ request: EmploymentRequest }>('/employment-requests', {
        method: 'POST',
        body: data
      })
      createdRequest.value = response.request
      isSuccess.value = true
      return true
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка создания заявки'
      return false
    } finally {
      isLoading.value = false
    }
  }

  const reset = () => {
    errorMessage.value = ''
    isSuccess.value = false
    createdRequest.value = null
  }

  return {
    isLoading,
    errorMessage,
    isSuccess,
    createdRequest,
    create,
    reset
  }
}
