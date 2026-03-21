<script setup lang="ts">
import type { UserEntity } from '@/types/backend/user'
import { formatDate, formatRole } from '@/utils/str'
import type { FramebyAppRole } from '~/types/frontend/enums/role';

import Tooltip from '@/components/ui/ToolTip.vue'
import ModalWindow from '@/components/common/ModalWindow.vue';
import TabsLayout from '@/components/ui/TabsLayout/TabsLayout.vue';
import SettingsTabs from '@/components/app/profile-tabs/SettingsTabs.vue';

const props = defineProps<{
  user: UserEntity
}>()

const authStore = useAuthStore();

const isModalSettingsOpen = ref(false);

const tabs = [
  {
    label: 'Настройки',
    value: 'settings',
    icon: 'ph:gear',
    component: SettingsTabs
  },
]

const handleLogout = async () => {
  await authStore.logout();
  navigateTo('/auth/login');
};


</script>

<template>
<div class="w-full max-w-xl mx-auto p-6
            rounded-3xl
            bg-white/80 backdrop-blur-xl
            border border-emerald-100
            shadow-xl shadow-emerald-500/10">

  <div class="flex items-center justify-between mb-6">

    <div class="flex items-center gap-4">
      <div class="relative">
        <div class="w-16 h-16 rounded-full
                    bg-linear-to-r from-emerald-400 to-green-600
                    flex items-center justify-center
                    text-white text-xl font-semibold">
          {{ user.login.charAt(0).toUpperCase() }}
        </div>

        <span class="absolute bottom-0 right-0 w-3 h-3
                     bg-emerald-400 rounded-full border-2 border-white" />
      </div>

      <div>
        <h2 class="text-lg font-semibold text-gray-800 flex items-center gap-2">
          <Icon name="ph:user" size="18" class="text-emerald-500" />
          {{ user.login }}
        </h2>

        <p class="text-sm text-gray-500 flex items-center gap-2 font-semibold">
          <Icon name="ph:envelope" size="16" />
          {{ user.email }}
        </p>
      </div>
    </div>
    <div class="flex items-center gap-2">
        <Tooltip text="Настройки">
            <button
                @click="isModalSettingsOpen = true"
                class="p-2 rounded-xl
                        bg-white/70
                        border border-emerald-100
                        hover:bg-emerald-50
                        transition flex items-center cursor-pointer"
            >
                <Icon name="ph:gear" size="20" class="text-emerald-600" />
            </button>
        </Tooltip>
        <Tooltip text="Выход из аккаунта">
            <button
                @click="handleLogout"
                class="p-2 rounded-xl
                        bg-white/70
                        border border-emerald-100
                        hover:bg-emerald-50
                        transition flex items-center cursor-pointer"
            >
                <Icon name="ph:sign-out" size="20" class="text-emerald-600" />
            </button>
        </Tooltip>
    </div>
  </div>

  <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 auto-rows-fr">
    <Tooltip text="Уникальный идентификатор пользователя">
        <div class="w-full h-full flex flex-col justify-between p-4 rounded-xl bg-white/70 border border-emerald-100">

        <p class="text-gray-400 text-xs mb-1 flex items-center gap-1">
          <Icon name="ph:hash" size="14" />
          ID
        </p>
        <p class="text-gray-800 font-normal">
          {{ user.id }}
        </p>
      </div>
    </Tooltip>

    <Tooltip text="Роль пользователя в системе">
        <div class="w-full h-full flex flex-col justify-between p-4 rounded-xl bg-white/70 border border-emerald-100">

        <p class="text-gray-400 text-xs mb-1 flex items-center gap-1">
          <Icon name="ph:shield-check" size="14" />
          Роль
        </p>
        <p class="text-emerald-600 font-semibold uppercase">
          {{ formatRole(user.role as FramebyAppRole) }}
        </p>
      </div>
    </Tooltip>

    <Tooltip text="Дата регистрации">
        <div class="w-full h-full flex flex-col justify-between p-4 rounded-xl bg-white/70 border border-emerald-100">

        <p class="text-gray-400 text-xs mb-1 flex items-center gap-1">
          <Icon name="ph:calendar" size="14" />
          Создан
        </p>
        <p class="text-gray-800 font-semibold">
          {{ formatDate(user.created_at) }}
        </p>
      </div>
    </Tooltip>

    <Tooltip text="Последнее обновление профиля">
        <div class="w-full h-full flex flex-col justify-between p-4 rounded-xl bg-white/70 border border-emerald-100">
        <p class="text-gray-400 text-xs mb-1 flex items-center gap-1">
          <Icon name="ph:clock-clockwise" size="14" />
          Обновлён
        </p>
        <p class="text-gray-800 font-semibold">
          {{ formatDate(user.updated_at) }}
        </p>
      </div>
    </Tooltip>
    
  </div>
  <ModalWindow
    v-model="isModalSettingsOpen"
    title="Настройки"
    customWidth="150vh"
    @close="isModalSettingsOpen = false"
  >
    <TabsLayout :tabs="tabs"/>
  </ModalWindow>
</div>
</template>