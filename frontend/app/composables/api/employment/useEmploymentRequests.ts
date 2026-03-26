import { $api } from '@/composables/api/useApi'
import type { EmploymentRequest, EmploymentStatus, ParticipantStatus } from '@/types/frontend/employment'

export interface EmploymentResponse {
  requests: EmploymentRequest[]
  total: number
  limit: number
  offset: number
}

export function useEmploymentRequests() {
  const requests = ref<EmploymentRequest[]>([])
  const total = ref(0)
  const isLoading = ref(false)
  const errorMessage = ref('')

  const fetchRequests = async (options: {
    status?: string
    limit?: number
    offset?: number
  } = {}) => {
    isLoading.value = true
    errorMessage.value = ''

    try {
      const params = new URLSearchParams()
      if (options.status) params.append('status', options.status)
      if (options.limit) params.append('limit', options.limit.toString())
      if (options.offset) params.append('offset', options.offset.toString())

      const query = params.toString() ? `?${params.toString()}` : ''

      const response = await $api<EmploymentResponse>(`/employment-requests${query}`, {
        method: 'GET'
      })

      requests.value = response.requests
      total.value = response.total
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка загрузки заявок'
    } finally {
      isLoading.value = false
    }
  }

  return {
    requests,
    total,
    isLoading,
    errorMessage,
    fetchRequests
  }
}

export function useMyOrganizationRequests() {
  const requests = ref<EmploymentRequest[]>([])
  const isLoading = ref(false)
  const errorMessage = ref('')

  const fetchRequests = async () => {
    isLoading.value = true
    errorMessage.value = ''

    try {
      const response = await $api<{ requests: EmploymentRequest[] }>('/employment-requests/my-organization', {
        method: 'GET'
      })
      requests.value = response.requests || []
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка загрузки'
    } finally {
      isLoading.value = false
    }
  }

  return {
    requests,
    isLoading,
    errorMessage,
    fetchRequests
  }
}

export function useMyApplications() {
  const requests = ref<EmploymentRequest[]>([])
  const isLoading = ref(false)
  const errorMessage = ref('')

  const fetchRequests = async () => {
    isLoading.value = true
    errorMessage.value = ''

    try {
      const response = await $api<{ requests: EmploymentRequest[] }>('/employment-requests/my-applications', {
        method: 'GET'
      })
      requests.value = response.requests
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка загрузки'
    } finally {
      isLoading.value = false
    }
  }

  return {
    requests,
    isLoading,
    errorMessage,
    fetchRequests
  }
}

export function useEmploymentStatuses() {
  const statuses = ref<EmploymentStatus[]>([])
  const isLoading = ref(false)
  const errorMessage = ref('')

  const fetchStatuses = async () => {
    isLoading.value = true
    errorMessage.value = ''

    try {
      const response = await $api<{ statuses: EmploymentStatus[] }>('/employment-requests/statuses', {
        method: 'GET'
      })
      statuses.value = response.statuses
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка загрузки статусов'
    } finally {
      isLoading.value = false
    }
  }

  return {
    statuses,
    isLoading,
    errorMessage,
    fetchStatuses
  }
}

export function useParticipantStatuses() {
  const statuses = ref<ParticipantStatus[]>([])
  const isLoading = ref(false)
  const errorMessage = ref('')

  const fetchStatuses = async () => {
    isLoading.value = true
    errorMessage.value = ''

    try {
      const response = await $api<{ statuses: ParticipantStatus[] }>('/employment-requests/participant-statuses', {
        method: 'GET'
      })
      statuses.value = response.statuses
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка загрузки статусов'
    } finally {
      isLoading.value = false
    }
  }

  return {
    statuses,
    isLoading,
    errorMessage,
    fetchStatuses
  }
}

export function useUpdateEmploymentRequest() {
  const isLoading = ref(false)
  const errorMessage = ref('')
  const isSuccess = ref(false)

  const update = async (
    requestId: string,
    data: {
      title?: string
      description?: string
      requirements?: string
      salary?: string
      schedule?: string
      max_participants?: number
    }
  ): Promise<boolean> => {
    isLoading.value = true
    errorMessage.value = ''
    isSuccess.value = false

    try {
      await $api(`/employment-requests/${requestId}`, {
        method: 'PUT',
        body: data
      })
      isSuccess.value = true
      return true
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка обновления заявки'
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
    update,
    reset
  }
}
