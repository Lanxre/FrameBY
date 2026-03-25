import { $api } from '@/composables/api/useApi'

export function useUpdateSquad() {
  const isLoading = ref(false)
  const errorMessage = ref('')
  const isSuccess = ref(false)

  const updateStatus = async (squadId: string, status: string): Promise<boolean> => {
    isLoading.value = true
    errorMessage.value = ''
    isSuccess.value = false

    try {
      await $api(`/student-squads/${squadId}`, {
        method: 'PUT',
        body: { status }
      })
      isSuccess.value = true
      return true
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка обновления'
      return false
    } finally {
      isLoading.value = false
    }
  }

  const closeRecruitment = async (squadId: string): Promise<boolean> => {
    return updateStatus(squadId, 'closed')
  }

  const openRecruitment = async (squadId: string): Promise<boolean> => {
    return updateStatus(squadId, 'recruitment_open')
  }

  return {
    isLoading,
    errorMessage,
    isSuccess,
    updateStatus,
    closeRecruitment,
    openRecruitment
  }
}
