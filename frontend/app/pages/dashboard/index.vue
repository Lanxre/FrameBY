<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import { FramebyAppRole } from '~/types/frontend/enums/role'
import { useRolePermissions } from '~/composables/api/role/useRolePermissions'

import AdminDashboard from '~/components/app/dashboard/AdminDashboard.vue'
import BrsmDashboard from '~/components/app/squads/BrsmDashboard.vue'
import StudentDashboard from '~/components/app/squads/StudentDashboard.vue'
import UniversityDashboard from '~/components/app/university/UniversityDashboard.vue'
import CustomerDashboard from '~/components/app/customer/CustomerDashboard.vue'


definePageMeta({
  middleware: ['auth']
})

const { isExactRole } = useRolePermissions()

const currentDashboard = computed(() => {
  if (isExactRole(FramebyAppRole.ADMIN)) return FramebyAppRole.ADMIN
  if (isExactRole(FramebyAppRole.STUDENT)) return FramebyAppRole.STUDENT
  if (isExactRole(FramebyAppRole.BRSM)) return FramebyAppRole.BRSM
  if (isExactRole(FramebyAppRole.UNIVERSITY)) return FramebyAppRole.UNIVERSITY
  if (isExactRole(FramebyAppRole.CUSTOMER)) return FramebyAppRole.CUSTOMER
  return 'default'
})
</script>

<template>
  <div class="w-380 mx-auto px-4 py-8 sm:py-12 
  border-2 border-emerald-400 mt-20 mb-10 rounded-2xl
  bg-linear-to-br from-emerald-300/80 via-emerald-500/70 to-emerald-600/60
  shadow-lg/30 
  ">
    <div class="mb-8 cursor-default">
      <h1 class="text-2xl sm:text-3xl font-bold text-white text-shadow-lg/30 flex items-center gap-3">
        <Icon name="ph:squares-four" size="28" class="text-emerald-500" />
        Панель управления
      </h1>
      <p v-if="currentDashboard === FramebyAppRole.ADMIN" class="mt-2 text-gray-600">
        Управление данными и профилями системы
      </p>
    </div>

    <AdminDashboard v-if="currentDashboard === FramebyAppRole.ADMIN" />
    <BrsmDashboard v-else-if="currentDashboard === FramebyAppRole.BRSM" />
    <StudentDashboard v-else-if="currentDashboard === FramebyAppRole.STUDENT" />
    <UniversityDashboard v-else-if="currentDashboard === FramebyAppRole.UNIVERSITY" />
    <CustomerDashboard v-else-if="currentDashboard === FramebyAppRole.CUSTOMER" />
    <div v-else class="rounded-2xl bg-white/80 border border-emerald-100 p-8 shadow-lg text-center">
      <Icon name="ph:info" size="48" class="text-emerald-400 mx-auto mb-4" />
      <p class="text-gray-600">Выберите профиль для заполнения данных в настройках.</p>
    </div>
    
  </div>
</template>
