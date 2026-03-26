import { $api } from '@/composables/api/useApi'

export function useApplyToEmployment() {
  const isLoading = ref(false)
  const errorMessage = ref('')
  const isSuccess = ref(false)

  const apply = async (requestId: string): Promise<boolean> => {
    isLoading.value = true
    errorMessage.value = ''
    isSuccess.value = false

    try {
      await $api(`/employment-requests/${requestId}/apply`, {
        method: 'POST'
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

export function useApproveEmployment() {
  const isLoading = ref(false)
  const errorMessage = ref('')
  const isSuccess = ref(false)

  const approve = async (requestId: string): Promise<boolean> => {
    isLoading.value = true
    errorMessage.value = ''
    isSuccess.value = false

    try {
      await $api(`/employment-requests/${requestId}/approve`, {
        method: 'POST',
        body: { approved: true }
      })
      isSuccess.value = true
      return true
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка одобрения заявки'
      return false
    } finally {
      isLoading.value = false
    }
  }

  const reject = async (requestId: string): Promise<boolean> => {
    isLoading.value = true
    errorMessage.value = ''
    isSuccess.value = false

    try {
      await $api(`/employment-requests/${requestId}/approve`, {
        method: 'POST',
        body: { approved: false }
      })
      isSuccess.value = true
      return true
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка отклонения заявки'
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
    approve,
    reject,
    reset
  }
}

export function useUpdateEmploymentStatus() {
  const isLoading = ref(false)
  const errorMessage = ref('')
  const isSuccess = ref(false)

  const updateStatus = async (
    requestId: string,
    status: string
  ): Promise<boolean> => {
    isLoading.value = true
    errorMessage.value = ''
    isSuccess.value = false

    try {
      await $api(`/employment-requests/${requestId}`, {
        method: 'PUT',
        body: { status }
      })
      isSuccess.value = true
      return true
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка обновления статуса'
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
    updateStatus,
    reset
  }
}

export function useUpdateParticipantStatus() {
  const isLoading = ref(false)
  const errorMessage = ref('')
  const isSuccess = ref(false)

  const updateStatus = async (
    requestId: string,
    userId: string,
    status: string
  ): Promise<boolean> => {
    isLoading.value = true
    errorMessage.value = ''
    isSuccess.value = false

    try {
      await $api(`/employment-requests/${requestId}/participants/${userId}/status`, {
        method: 'PATCH',
        body: { status }
      })
      isSuccess.value = true
      return true
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка обновления статуса'
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
    updateStatus,
    reset
  }
}
