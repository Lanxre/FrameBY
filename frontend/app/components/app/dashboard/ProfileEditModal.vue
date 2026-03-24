<script setup lang="ts">
import { PROFILE_TYPES, type ProfileRow } from '@/types/dashboard/profile'
import Select from '@/components/ui/Select/Select.vue'
import { $api } from '@/composables/api/useApi'
import { formatAvatar } from '@/utils/str'

const props = defineProps<{
  modelValue: boolean
  profile: ProfileRow | null
}>()

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
  'saved': []
}>()

const roleOptions = PROFILE_TYPES
  .filter(t => t.value !== 'all')
  .map(t => ({
    id: t.value,
    name: t.label
  }))

const selectedRole = ref<{ id: string; name: string } | null>(null)
const fullName = ref('')
const subrole = ref('')
const faculty = ref('')
const specialty = ref('')
const grade = ref('')
const department = ref('')
const position = ref('')
const isLoading = ref(false)
const errorMessage = ref('')

watch(() => props.profile, (newProfile) => {
  if (newProfile) {
    const roleOption = roleOptions.find(r => r.id === newProfile.role)
    selectedRole.value = roleOption || null
    fullName.value = newProfile.full_name || ''
    subrole.value = newProfile.subrole || ''
    faculty.value = newProfile.faculty || ''
    specialty.value = newProfile.specialty || ''
    grade.value = newProfile.grade?.toString() || ''
    department.value = newProfile.department || ''
    position.value = newProfile.position || ''
  }
})

watch(() => props.modelValue, (isOpen) => {
  if (!isOpen) {
    errorMessage.value = ''
  }
})

const close = () => {
  emit('update:modelValue', false)
}

const handleSave = async () => {
  if (!props.profile || !selectedRole.value) return
  
  isLoading.value = true
  errorMessage.value = ''
  
  try {
    const body: any = {
      role: selectedRole.value.id
    }

    if (selectedRole.value.id !== 'user') {
      if (fullName.value) body.full_name = fullName.value
    }

    if (selectedRole.value.id === 'brsm' || selectedRole.value.id === 'university' || selectedRole.value.id === 'customer') {
      if (subrole.value) body.subrole = subrole.value
    }

    if (selectedRole.value.id === 'student') {
      if (faculty.value) body.faculty = faculty.value
      if (specialty.value) body.specialty = specialty.value
      if (grade.value) body.grade = parseFloat(grade.value)
    }

    if (selectedRole.value.id === 'university') {
      if (faculty.value) body.faculty = faculty.value
      if (department.value) body.department = department.value
    }

    if (selectedRole.value.id !== 'user' && position.value) {
      body.position = position.value
    }

    await $api(`/admin/profiles/${props.profile.user_id}/subrole`, {
      method: 'PATCH',
      body
    })
    emit('saved')
    close()
  } catch (e: any) {
    errorMessage.value = e.data?.error || e.message || 'Ошибка сохранения'
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <Teleport to="body">
    <Transition
      enter-active-class="transition duration-200 ease-out"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition duration-150 ease-in"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div 
        v-if="modelValue"
        class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/30 backdrop-blur-sm"
        @click.self="close"
      >
        <Transition
          enter-active-class="transition duration-200 ease-out"
          enter-from-class="transform scale-95 opacity-0"
          enter-to-class="transform scale-100 opacity-100"
          leave-active-class="transition duration-150 ease-in"
          leave-from-class="transform scale-100 opacity-100"
          leave-to-class="transform scale-95 opacity-0"
        >
          <div 
            v-if="modelValue"
            class="w-full max-w-lg rounded-2xl bg-white/95 backdrop-blur-xl 
                   border border-emerald-100 shadow-xl shadow-emerald-500/10"
          >
            <div class="flex items-center justify-between px-6 py-4 border-b border-emerald-100">
              <h3 class="text-lg font-semibold text-gray-800">
                Редактирование профиля
              </h3>
              <button 
                @click="close"
                class="w-9 h-9 flex items-center justify-center rounded-xl hover:bg-emerald-50 transition cursor-pointer"
              >
                <Icon name="ph:x" size="20" class="text-gray-500" />
              </button>
            </div>

            <div class="p-6 space-y-5 max-h-[70vh] overflow-y-auto" v-if="profile">
              <div class="flex items-center gap-4 p-4 rounded-xl bg-emerald-50/50 border border-emerald-100">
                <div class="w-12 h-12 rounded-full bg-linear-to-r from-emerald-400 to-green-600 flex items-center justify-center text-white font-semibold text-lg shrink-0">
                    <img v-if="profile.avatar" :src="formatAvatar(profile.login, profile.avatar)" alt="Avatar" class="w-full h-full object-cover rounded-full" />
                    <span v-else>{{ profile.login.charAt(0).toUpperCase() }}</span>
                </div>
                <div class="min-w-0">
                  <p class="font-medium text-gray-800 truncate">{{ profile.login }}</p>
                  <p class="text-sm text-gray-500 truncate">{{ profile.email }}</p>
                </div>
              </div>

              <div v-if="errorMessage" class="text-sm text-red-500 bg-red-50 border border-red-200 px-4 py-3 rounded-xl">
                {{ errorMessage }}
              </div>

              <div class="space-y-2">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Роль</label>
                <Select
                  v-model="selectedRole"
                  :options="roleOptions"
                  placeholder="Выберите роль"
                  icon="ph:identification-card"
                />
              </div>

              <div class="space-y-2" v-if="selectedRole?.id !== 'user'">
                <label class="text-sm font-medium text-gray-700 ml-1.5">ФИО</label>
                <div class="relative">
                  <Icon name="ph:user" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
                  <input
                    v-model="fullName"
                    type="text"
                    placeholder="Иванов Иван Иванович"
                    class="w-full pl-9 pr-4 py-2.5 rounded-xl
                           bg-white/60 border border-emerald-100
                           focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                           text-sm cursor-text"
                  />
                </div>
              </div>

              <div class="space-y-2" v-if="selectedRole?.id === 'brsm' || selectedRole?.id === 'university' || selectedRole?.id === 'customer'">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Суброль</label>
                <div class="relative">
                  <Icon name="ph:address-book" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
                  <input
                    v-model="subrole"
                    type="text"
                    placeholder="Например: координатор, декан..."
                    class="w-full pl-9 pr-4 py-2.5 rounded-xl
                           bg-white/60 border border-emerald-100
                           focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                           text-sm cursor-text"
                  />
                </div>
              </div>

              <div v-if="selectedRole?.id === 'student' || selectedRole?.id === 'university'" class="space-y-2">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Факультет</label>
                <div class="relative">
                  <Icon name="ph:graduation-cap" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
                  <input
                    v-model="faculty"
                    type="text"
                    placeholder="Факультет"
                    class="w-full pl-9 pr-4 py-2.5 rounded-xl
                           bg-white/60 border border-emerald-100
                           focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                           text-sm cursor-text"
                  />
                </div>
              </div>

              <div v-if="selectedRole?.id === 'student'" class="space-y-2">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Специальность</label>
                <div class="relative">
                  <Icon name="ph:book-open" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
                  <input
                    v-model="specialty"
                    type="text"
                    placeholder="Специальность"
                    class="w-full pl-9 pr-4 py-2.5 rounded-xl
                           bg-white/60 border border-emerald-100
                           focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                           text-sm cursor-text"
                  />
                </div>
              </div>

              <div v-if="selectedRole?.id === 'student'" class="space-y-2">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Оценка (1-10)</label>
                <div class="relative">
                  <Icon name="ph:star" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
                  <input
                    v-model="grade"
                    type="number"
                    min="1"
                    max="10"
                    step="0.1"
                    placeholder="9.5"
                    class="w-full pl-9 pr-4 py-2.5 rounded-xl
                           bg-white/60 border border-emerald-100
                           focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                           text-sm cursor-text"
                  />
                </div>
              </div>

              <div v-if="selectedRole?.id === 'university'" class="space-y-2">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Кафедра</label>
                <div class="relative">
                  <Icon name="ph:buildings" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
                  <input
                    v-model="department"
                    type="text"
                    placeholder="Кафедра"
                    class="w-full pl-9 pr-4 py-2.5 rounded-xl
                           bg-white/60 border border-emerald-100
                           focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                           text-sm cursor-text"
                  />
                </div>
              </div>

              <div class="space-y-2" v-if="selectedRole?.id !== 'user'">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Должность</label>
                <div class="relative">
                  <Icon name="ph:briefcase" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
                  <input
                    v-model="position"
                    type="text"
                    placeholder="Должность"
                    class="w-full pl-9 pr-4 py-2.5 rounded-xl
                           bg-white/60 border border-emerald-100
                           focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                           text-sm cursor-text"
                  />
                </div>
              </div>
            </div>

            <div class="flex items-center justify-end gap-3 px-6 py-4 border-t border-emerald-100">
              <button
                @click="close"
                class="px-5 py-2.5 rounded-xl text-sm font-medium
                       bg-white border border-emerald-200
                       text-gray-700 hover:bg-emerald-50 transition cursor-pointer"
              >
                Отмена
              </button>
              <button
                @click="handleSave"
                :disabled="isLoading || !selectedRole"
                class="px-5 py-2.5 rounded-xl text-sm font-semibold
                       bg-linear-to-r from-emerald-400 to-green-600 text-white
                       hover:opacity-90 transition cursor-pointer
                       disabled:opacity-50 disabled:cursor-not-allowed
                       shadow-lg shadow-emerald-500/20"
              >
                <span v-if="isLoading">Сохранение...</span>
                <span v-else>Сохранить</span>
              </button>
            </div>
          </div>
        </Transition>
      </div>
    </Transition>
  </Teleport>
</template>
