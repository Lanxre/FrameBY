import { $api } from '@/composables/api/useApi'

export function useUpdateSquad() {
  const isLoading = ref(false)
  const errorMessage = ref('')
  const isSuccess = ref(false)

  const update = async (
    squadId: string,
    data: {
      title?: string
      description?: string
      profile?: string
      max_participants?: number
      status?: string
    }
  ): Promise<boolean> => {
    isLoading.value = true
    errorMessage.value = ''
    isSuccess.value = false

    try {
      await $api(`/student-squads/${squadId}`, {
        method: 'PATCH',
        body: data
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
    return update(squadId, { status: 'closed' })
  }

  const openRecruitment = async (squadId: string): Promise<boolean> => {
    return update(squadId, { status: 'recruitment_open' })
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
    closeRecruitment,
    openRecruitment,
    reset
  }
}
