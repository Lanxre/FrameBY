<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const { user } = storeToRefs(authStore)

const stats = [
  { label: 'Всего профилей', value: 0, icon: 'ph:users', color: 'from-emerald-400 to-green-600' },
  { label: 'Студентов', value: 0, icon: 'ph:student', color: 'from-blue-400 to-blue-600' },
  { label: 'БРСМ', value: 0, icon: 'ph:users-three', color: 'from-purple-400 to-purple-600' },
  { label: 'Университетов', value: 0, icon: 'ph:graduation-cap', color: 'from-amber-400 to-amber-600' }
]
</script>

<template>
  <div class="space-y-6">
    <div class="flex items-center gap-3">
      <Icon name="ph:student" size="24" class="text-emerald-500" />
      <h2 class="text-xl font-semibold text-gray-800">Панель студента</h2>
    </div>

    <div class="rounded-2xl bg-white/80 border border-emerald-100 p-6 shadow-lg">
      <p class="text-gray-600">
        Добро пожаловать, <span class="font-semibold text-emerald-600">{{ user?.login }}</span>!
      </p>
      <p class="mt-2 text-sm text-gray-500">
        Здесь будет отображаться информация для студентов.
      </p>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
      <div v-for="(stat, index) in stats" :key="index"
           class="rounded-2xl bg-white/80 border border-emerald-100 p-4 shadow-lg">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-xs text-gray-500 uppercase tracking-wider">{{ stat.label }}</p>
            <p class="text-2xl font-bold text-gray-800 mt-1">{{ stat.value }}</p>
          </div>
          <div :class="`w-10 h-10 rounded-xl bg-linear-to-r ${stat.color} flex items-center justify-center shadow-lg`">
            <Icon :name="stat.icon" size="20" class="text-white" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
