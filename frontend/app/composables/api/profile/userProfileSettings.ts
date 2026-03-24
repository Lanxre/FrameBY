import { ref, reactive } from 'vue'
import { $api } from '../useApi'

export function useProfileSettings() {
  const showConfirm = ref(false)
  const isLoading = ref(false)
  
  const authStore = useAuthStore();
  const { notify } = useNotificationStore();
  
  const avatarPreview = ref<string | null>(formatAvatar(authStore.user!.login, authStore.user?.avatar) || null)
  const avatarFile = ref<File | null>(null)

  const fileInputRef = ref<HTMLInputElement | null>(null)

  const form = reactive({
    login: '',
    password: '',
    passwordConfirm: ''
  })

  const errorMessage = ref('')

  const openFileDialog = () => {
    fileInputRef.value?.click()
  }

  const handleAvatarChange = (e: Event) => {
    const file = (e.target as HTMLInputElement).files?.[0]
    if (!file) return

    avatarFile.value = file
    avatarPreview.value = URL.createObjectURL(file)
  }

  const handleSubmit = () => {
    errorMessage.value = ''

    if (form.password && form.password !== form.passwordConfirm) {
      errorMessage.value = 'Пароли не совпадают'
      return
    }

    showConfirm.value = true
  }

  const submit = async () => {
    isLoading.value = true

    try {
      const formData = new FormData()
      formData.append('login', form.login)
      formData.append('password', form.password)
      
      if (avatarFile.value) {
        formData.append('avatar', avatarFile.value)
      }
      
      await $api<any>("/user", {
        method: 'PATCH',
        body: formData
      })

      showConfirm.value = false
      await authStore.fetchUser()
      notify({
        type: "success",
        title: "Успех",
        content: "Профиль успешно обновлен",
      })
      
    } catch (e) {
      errorMessage.value = 'Ошибка при обновлении'
      notify({
        type: "error",
        title: "Ошибка",
        content: "Ошибка при обновлении профиля",
      })
    } finally {
      isLoading.value = false
    }
  }

  return {
    form,
    errorMessage,
    avatarPreview,
    avatarFile,
    fileInputRef,
    showConfirm,
    isLoading,
    openFileDialog,
    handleAvatarChange,
    handleSubmit,
    submit
  }
}