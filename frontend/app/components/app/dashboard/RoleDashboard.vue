<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import { FramebyAppRole } from '~/types/frontend/enums/role'

const props = defineProps<{
  roleLabel: string
  roleIcon: string
}>()

const authStore = useAuthStore()
const { user } = storeToRefs(authStore)
</script>

<template>
  <div class="space-y-6">
    <div class="flex items-center gap-3">
      <Icon :name="roleIcon" size="24" class="text-emerald-500" />
      <h2 class="text-xl font-semibold text-gray-800">{{ roleLabel }}</h2>
    </div>

    <div class="rounded-2xl bg-white/80 border border-emerald-100 p-6 shadow-lg">
      <p class="text-gray-600">
        Добро пожаловать, <span class="font-semibold text-emerald-600">{{ user?.login }}</span>!
      </p>
      <p class="mt-2 text-sm text-gray-500">
        Здесь будет отображаться информация для {{ roleLabel.toLowerCase() }}.
      </p>
    </div>
  </div>
</template>
