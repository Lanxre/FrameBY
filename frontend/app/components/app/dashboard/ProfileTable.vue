<script setup lang="ts">
import { formatDate } from '@/utils/str'
import { getProfileTypeInfo } from '@/types/dashboard/profile'
import type { ProfileRow } from '@/types/dashboard/profile'

const props = defineProps<{
  profiles: ProfileRow[]
  isLoading: boolean
  total: number
  currentPage: number
  totalPages: number
}>()

const emit = defineEmits<{
  'edit': [profile: ProfileRow]
  'page-change': [page: number]
}>()

const columns = [
  { key: 'user', label: 'Пользователь', width: 'w-48' },
  { key: 'type', label: 'Роль', width: 'w-36' },
  { key: 'fullName', label: 'ФИО', width: 'w-44' },
  { key: 'phone', label: 'Телефон', width: 'w-36' },
  { key: 'profile', label: 'Профиль', width: 'w-52' },
  { key: 'subrole', label: 'Суброль', width: 'w-36' },
  { key: 'updated', label: 'Обновлено', width: 'w-36' },
  { key: 'actions', label: '', width: 'w-16' }
]

</script>

<template>
  <div class="rounded-2xl bg-white/80 border border-emerald-100 shadow-lg overflow-hidden flex flex-col max-h-[calc(100vh-220px)]">
    <div class="overflow-x-auto overflow-y-auto flex-1 scrollbar-thin scrollbar-thumb-emerald-200 scrollbar-track-transparent">
      <table class="w-full">
        <thead class="sticky top-0 z-10">
          <tr class="bg-emerald-50/80">
            <th 
              v-for="col in columns" 
              :key="col.key"
              class="px-4 py-3 text-center text-xs font-semibold text-emerald-700 uppercase tracking-wider whitespace-nowrap"
              :class="col.width"
            >
              {{ col.label }}
            </th>
          </tr>
        </thead>
        <tbody class="divide-y divide-emerald-100">
          <tr v-if="isLoading">
            <td :colspan="columns.length" class="px-4 py-12 text-center">
              <div class="flex items-center justify-center gap-2">
                <Icon name="ph:circle-notch" size="24" class="animate-spin text-emerald-500" />
                <span class="text-gray-500">Загрузка...</span>
              </div>
            </td>
          </tr>
          <tr v-else-if="profiles.length === 0">
            <td :colspan="columns.length" class="px-4 py-12 text-center text-gray-500">
              Профили не найдены
            </td>
          </tr>
          <template v-else>
            <tr 
              v-for="row in profiles" 
              :key="row.user_id"
              class="hover:bg-emerald-50/30 transition"
            >
              <td class="px-4 py-3">
                <div class="flex items-center justify-center gap-3">
                  <div class="w-9 h-9 rounded-full bg-linear-to-r from-emerald-400 to-green-600 flex items-center justify-center text-white text-sm font-semibold shrink-0">
                      <img v-if="row.avatar" :src="formatAvatar(row.login, row.avatar)" alt="Avatar" class="w-full h-full object-cover rounded-full" />
                      <span v-else>{{ row.login.charAt(0).toUpperCase() }}</span>
                  </div>
                  <div class="text-left min-w-0">
                    <p class="text-sm font-medium text-gray-800 truncate max-w-32">{{ row.login }}</p>
                    <p class="text-xs text-gray-500 truncate max-w-32">{{ row.email }}</p>
                  </div>
                </div>
              </td>
              
              <td class="px-4 py-3">
                <div class="flex items-center flex-col justify-center gap-2">
                  <div 
                    class="w-8 h-8 rounded-lg flex items-center justify-center shrink-0" 
                    :class="[getProfileTypeInfo(row.role).bgColor, getProfileTypeInfo(row.role).color]"
                  >
                    <Icon :name="getProfileTypeInfo(row.role).icon" size="16" />
                  </div>
                  <span class="text-sm text-gray-700 whitespace-nowrap">{{ getProfileTypeInfo(row.role).label }}</span>
                </div>
              </td>
              
              <td class="px-4 py-3 text-center">
                <span class="text-sm text-gray-800">{{ row.full_name || '-' }}</span>
              </td>
              
              <td class="px-4 py-3 text-center">
                <span class="text-sm text-gray-700">{{ row.phone || '-' }}</span>
              </td>
              
              <td class="px-4 py-3">
                <div class="flex flex-col gap-2 text-xs text-gray-600 space-y-0.5 text-center">
                  <p v-if="row.faculty">Факультет: {{ row.faculty }}</p>
                  <p v-if="row.specialty">Специальность: {{ row.specialty }}</p>
                  <p v-if="row.grade">Оценка: {{ row.grade }}</p>
                  <p v-if="row.department">Кафедра: {{ row.department }}</p>
                  <p v-if="row.position">Должность: {{ row.position }}</p>
                  <p v-if="!row.faculty && !row.specialty && !row.grade && !row.department && !row.position" class="text-gray-400">-</p>
                </div>
              </td>
              
              <td class="px-4 py-3 text-center">
                <span class="text-sm" :class="row.subrole ? 'text-gray-700' : 'text-gray-400'">
                  {{ row.subrole || 'Не назначена' }}
                </span>
              </td>
              
              <td class="px-4 py-3 text-center">
                <span class="text-sm text-gray-500 whitespace-nowrap">{{ formatDate(row.updated_at) }}</span>
              </td>
              
              <td class="px-4 py-3">
                <div class="flex items-center justify-center">
                  <button
                    @click="emit('edit', row)"
                    class="w-9 h-9 flex items-center justify-center rounded-xl hover:bg-emerald-100 transition cursor-pointer"
                    title="Редактировать"
                  >
                    <Icon name="ph:pencil-simple" size="18" class="text-emerald-500" />
                  </button>
                </div>
              </td>
            </tr>
          </template>
        </tbody>
      </table>
    </div>

    <div class="px-4 py-3 bg-emerald-50/30 border-t border-emerald-100 flex items-center justify-between shrink-0">
      <p class="text-sm text-gray-600">
        Показано {{ profiles.length }} из {{ total }}
      </p>
      <div v-if="totalPages > 1" class="flex items-center gap-2">
        <button
          @click="emit('page-change', Math.max(1, currentPage - 1))"
          :disabled="currentPage === 1"
          class="w-9 h-9 flex items-center justify-center rounded-lg bg-white border border-emerald-200
                 hover:bg-emerald-50 disabled:opacity-50 disabled:cursor-not-allowed transition cursor-pointer"
        >
          <Icon name="ph:caret-left" size="18" />
        </button>
        <span class="text-sm text-gray-600 min-w-16 text-center">
          {{ currentPage }} / {{ totalPages }}
        </span>
        <button
          @click="emit('page-change', Math.min(totalPages, currentPage + 1))"
          :disabled="currentPage === totalPages"
          class="w-9 h-9 flex items-center justify-center rounded-lg bg-white border border-emerald-200
                 hover:bg-emerald-50 disabled:opacity-50 disabled:cursor-not-allowed transition cursor-pointer"
        >
          <Icon name="ph:caret-right" size="18" />
        </button>
      </div>
    </div>
  </div>
</template>
