import { ref } from 'vue'
import { $api } from '@/composables/api/useApi'

export function useUniversityDepartments() {
  const departments = ref<{ id: string; name: string }[]>([])
  const isLoading = ref(false)
  const errorMessage = ref('')

  const fetchDepartments = async () => {
    isLoading.value = true
    errorMessage.value = ''

    try {
      const res: { departments: any[] } = await $api('/university-departments')
      departments.value = res.departments.map((d: any) => ({
        id: d.id,
        name: `${d.university_name} / ${d.department_name}`
      }))
    } catch (e: any) {
      errorMessage.value = e.message || 'Ошибка загрузки данных'
    } finally {
      isLoading.value = false
    }
  }

  return {
    departments,
    isLoading,
    errorMessage,
    fetchDepartments
  }
}
