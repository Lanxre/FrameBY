<script setup lang="ts">
import Tooltip from "@/components/ui/ToolTip.vue";
import ModalConfirm from "@/components/common/ModalConfirm.vue";
import { useProfileSettings } from "@/composables/api/profile/userProfileSettings";

const {
	form,
	errorMessage,
	avatarPreview,
	fileInputRef,
	showConfirm,
	isLoading,
	openFileDialog,
	handleAvatarChange,
	handleSubmit,
	submit,
} = useProfileSettings();
</script>

<template>
<div class="min-h-120">

  <div class="flex items-center gap-2">
    <Icon name="ph:gear" size="22" class="text-emerald-500" />
    <h2 class="text-lg font-semibold text-gray-800 uppercase">
      Настройки профиля
    </h2>
  </div>

  <div class="rounded-2xl
              bg-white/80 backdrop-blur-xl
              border border-emerald-100
              shadow-lg shadow-emerald-500/10
              p-6 space-y-6">

    <Tooltip text="Нажмите, чтобы изменить аватар" position="top">
      <div class="flex items-center gap-4 cursor-pointer" @click="openFileDialog">
        <div class="w-16 h-16 rounded-full overflow-hidden
                    bg-emerald-100 flex items-center justify-center">
                        
          <img v-if="avatarPreview && avatarPreview.length !== 1" :src="avatarPreview" class="w-full h-full object-cover" />
          <Icon v-else name="ph:user" size="26" class="text-emerald-500" />
        </div>

        <span class="text-sm text-emerald-600 hover:underline">
          Изменить аватар
        </span>
      </div>
    </Tooltip>

    <input ref="fileInputRef" type="file" class="hidden" @change="handleAvatarChange" />

    <div v-if="errorMessage"
         class="text-sm text-red-500 bg-red-50 border border-red-200 px-3 py-2 rounded-xl">
      {{ errorMessage }}
    </div>

    <div class="space-y-1">
      <label class="text-xs text-gray-500 ml-2">Логин</label>
      <p class="text-[11px] text-gray-400 ml-2">
        Отображается в системе
      </p>
    
      <div class="relative">
        <Icon name="ph:user" size="18"
              class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
    
        <input
          v-model="form.login"
          type="text"
          placeholder="Новый логин"
          class="w-full pl-9 pr-3 py-2 rounded-xl
                 bg-white/60 border border-emerald-100
                 focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                 text-sm"
        />
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
    
      <div class="space-y-1">
        <label class="text-xs text-gray-500 ml-2">Новый пароль</label>
        <p class="text-[11px] text-gray-400 ml-2">
          Минимум 6 символов
        </p>
    
        <div class="relative">
          <Icon name="ph:lock" size="18"
                class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
    
          <input
            v-model="form.password"
            type="password"
            placeholder="••••••••"
            class="w-full pl-9 pr-3 py-2 rounded-xl
                   bg-white/60 border border-emerald-100
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                   text-sm"
          />
        </div>
      </div>
    
      <div class="space-y-1">
        <label class="text-xs text-gray-500 ml-2">Подтвердите пароль</label>
        <p class="text-[11px] text-gray-400 ml-2">
          Должен совпадать с паролем
        </p>
    
        <div class="relative">
          <Icon name="ph:lock-key" size="18"
                class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
    
          <input
            v-model="form.passwordConfirm"
            type="password"
            placeholder="••••••••"
            class="w-full pl-9 pr-3 py-2 rounded-xl
                   bg-white/60 border border-emerald-100
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                   text-sm"
          />
        </div>
      </div>
    
    </div>

    <button
      @click="handleSubmit"
      class="w-full py-2.5 rounded-xl text-sm font-semibold
             bg-linear-to-r from-emerald-400 to-green-600 text-white
             hover:opacity-90 active:scale-95
             transition shadow-lg shadow-emerald-500/20"
    >
      Сохранить изменения
    </button>

  </div>

  <ModalConfirm
    v-model="showConfirm"
    title="Сохранение изменений"
    description="Вы уверены, что хотите изменить данные профиля?"
    confirmText="Сохранить"
    cancelText="Отмена"
    :loading="isLoading"
    @confirm="submit"
  />

</div>
</template>