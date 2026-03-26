import { $api } from '@/composables/api/useApi'
import type { EmploymentRequest, CreateEmploymentRequestData } from '@/types/frontend/employment'

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

export function useApplyToRequest() {
  const isLoading = ref(false)
  const errorMessage = ref('')
  const isSuccess = ref(false)

  const apply = async (requestId: string): Promise<boolean> => {
    isLoading.value = true
    errorMessage.value = ''
    isSuccess.value = false

    try {
      await $api('/employment-requests/apply', {
        method: 'POST',
        body: { request_id: requestId }
      })
      isSuccess.value = true
      return true
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка подачи заявки'
      return false
    } finally {
      isLoading.value = false
    }
  }

  const reset = () => {
    errorMessage.value = ''
    isSuccess.value = false
  }

  return {
    isLoading,
    errorMessage,
    isSuccess,
    apply,
    reset
  }
}

export function useCancelApplication() {
  const isLoading = ref(false)
  const errorMessage = ref('')
  const isSuccess = ref(false)

  const cancel = async (requestId: string): Promise<boolean> => {
    isLoading.value = true
    errorMessage.value = ''
    isSuccess.value = false

    try {
      await $api('/employment-requests/cancel', {
        method: 'POST',
        body: { request_id: requestId }
      })
      isSuccess.value = true
      return true
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка отмены заявки'
      return false
    } finally {
      isLoading.value = false
    }
  }

  const reset = () => {
    errorMessage.value = ''
    isSuccess.value = false
  }

  return {
    isLoading,
    errorMessage,
    isSuccess,
    cancel,
    reset
  }
}
