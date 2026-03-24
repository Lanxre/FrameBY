import { defu } from 'defu'
import type { ProfileRow } from '@/types/dashboard/profile'

export interface ProfilesResponse {
  profiles: ProfileRow[]
  total: number
  limit: number
  offset: number
}

export function useDashboardProfiles() {
  const profiles = ref<ProfileRow[]>([])
  const total = ref(0)
  const isLoading = ref(false)
  const errorMessage = ref('')

  const fetchProfiles = async (options: {
    type?: string
    search?: string
    limit?: number
    offset?: number
  } = {}) => {
    isLoading.value = true
    errorMessage.value = ''

    try {
      const params = new URLSearchParams()
      if (options.type) params.append('profile_type', options.type)
      if (options.search) params.append('search', options.search)
      if (options.limit) params.append('limit', options.limit.toString())
      if (options.offset) params.append('offset', options.offset.toString())

      const query = params.toString() ? `?${params.toString()}` : ''
      const defaults = {
        baseURL: '/api',
        headers: useRequestHeaders(['cookie']),
      }

      const response = await $fetch<ProfilesResponse>(`/api/admin/profiles${query}`, defaults)
      
      profiles.value = response.profiles
      total.value = response.total
    } catch (e: any) {
      errorMessage.value = e.data?.error || e.message || 'Ошибка загрузки профилей'
    } finally {
      isLoading.value = false
    }
  }

  const updateSubrole = async (userId: string, role: string, subrole: string) => {
    try {
      const defaults = {
        baseURL: '/api',
        headers: useRequestHeaders(['cookie']),
      }

      await $fetch(`/api/admin/profiles/${userId}/subrole`, {
        ...defaults,
        method: 'PATCH',
        body: { role, subrole }
      })
      return true
    } catch (e: any) {
      errorMessage.value = e.data?.error || e.message || 'Ошибка обновления суброли'
      return false
    }
  }

  const deleteProfile = async (userId: string) => {
    try {
      const defaults = {
        baseURL: '/api',
        headers: useRequestHeaders(['cookie']),
      }

      await $fetch(`/api/admin/profiles/${userId}`, {
        ...defaults,
        method: 'DELETE'
      })
      return true
    } catch (e: any) {
      errorMessage.value = e.data?.error || e.message || 'Ошибка удаления профиля'
      return false
    }
  }

  return {
    profiles,
    total,
    isLoading,
    errorMessage,
    fetchProfiles,
    updateSubrole,
    deleteProfile
  }
}
