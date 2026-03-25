import { $api } from '@/composables/api/useApi'
import type { StudentSquad, SquadStatus } from '@/types/frontend/student-squad'

export interface SquadsResponse {
  squads: StudentSquad[]
  total: number
  limit: number
  offset: number
}

export function useStudentSquads() {
  const squads = ref<StudentSquad[]>([])
  const total = ref(0)
  const isLoading = ref(false)
  const errorMessage = ref('')

  const fetchSquads = async (options: {
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

      const response = await $api<SquadsResponse>(`/student-squads${query}`, {
        method: 'GET'
      })

      squads.value = response.squads
      total.value = response.total
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка загрузки отрядов'
    } finally {
      isLoading.value = false
    }
  }

  return {
    squads,
    total,
    isLoading,
    errorMessage,
    fetchSquads
  }
}

export function useMySquads() {
  const squads = ref<StudentSquad[]>([])
  const isLoading = ref(false)
  const errorMessage = ref('')

  const fetchMySquads = async () => {
    isLoading.value = true
    errorMessage.value = ''

    try {
      const response = await $api<{ squads: StudentSquad[] }>('/student-squads/my', {
        method: 'GET'
      })
      squads.value = response.squads
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка загрузки'
    } finally {
      isLoading.value = false
    }
  }

  return {
    squads,
    isLoading,
    errorMessage,
    fetchMySquads
  }
}

export function useSquadStatuses() {
  const statuses = ref<SquadStatus[]>([])
  const isLoading = ref(false)
  const errorMessage = ref('')

  const fetchStatuses = async () => {
    isLoading.value = true
    errorMessage.value = ''

    try {
      const response = await $api<{ statuses: SquadStatus[] }>('/student-squads/statuses', {
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
