import { ref, reactive } from 'vue'

export function useProfileSettings() {
  const showConfirm = ref(false)
  const isLoading = ref(false)

  const avatarPreview = ref<string | null>(null)
  const avatarFile = ref<File | null>(null)

  const fileInputRef = ref<HTMLInputElement | null>(null)

  const form = reactive({
    username: '',
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
      console.log({
        username: form.username,
        password: form.password,
        avatar: avatarFile.value
      })

      showConfirm.value = false
    } catch (e) {
      errorMessage.value = 'Ошибка при обновлении'
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