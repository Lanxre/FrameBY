<script setup lang="ts">
import { ref } from 'vue'
import { useLogin } from '@/composables/api/auth/useLogin'

const {
  form,
  isLoading,
  errorMessage,
  handleManualLogin
} = useLogin();

const showPassword = ref(false);
</script>

<template>
<div class="min-h-screen flex items-center justify-center p-4 sm:p-6 relative overflow-hidden">

  <div class="absolute inset-0 pointer-events-none">
    <div class="absolute top-30 left-1/2 -translate-x-1/2 w-200 h-150
                bg-emerald-400/20 blur-2xl rounded-full" />
  </div>

  <div class="relative w-full max-w-md p-6 sm:p-8
              rounded-2xl
              bg-white/70 backdrop-blur
              border border-black/5
              shadow-lg">

    <div class="text-center mb-6 flex items-center justify-center gap-2">
      <Icon name="mdi:login" size="28" class="text-emerald-500" />
      <h1 class="text-3xl font-semibold uppercase
                 bg-linear-to-r from-emerald-400 to-green-600
                 bg-clip-text text-transparent">
        Вход
      </h1>
    </div>

    <div
      v-if="errorMessage"
      class="text-sm text-red-500 bg-red-50 border border-red-200
             px-3 py-2 rounded-xl mb-4"
    >
      {{ errorMessage }}
    </div>

    <form @submit.prevent="handleManualLogin" class="space-y-4">

      <div class="space-y-1">
        <label class="text-xs text-gray-500 ml-1">Почта</label>
        <div class="relative">
          <Icon name="mdi:email-outline"
                size="18"
                class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />

          <input
            v-model="form.email"
            type="email"
            placeholder="example@mail.com"
            class="w-full pl-9 pr-3 py-2 rounded-xl
                   bg-white/60
                   border border-black/10
                   focus:outline-none
                   focus:ring-2 focus:ring-emerald-400/40
                   text-sm"
          />
        </div>
      </div>

      <div class="space-y-1">
        <label class="text-xs text-gray-500 ml-1">Пароль</label>
        <div class="relative">
          <Icon name="mdi:lock-outline"
                size="18"
                class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />

          <input
            v-model="form.password"
            :type="showPassword ? 'text' : 'password'"
            placeholder="••••••••"
            class="w-full pl-9 pr-10 py-2 rounded-xl
                   bg-white/60
                   border border-black/10
                   focus:outline-none
                   focus:ring-2 focus:ring-emerald-400/40
                   text-sm"
          />

          <button
            type="button"
            @click="showPassword = !showPassword"
            class="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-emerald-500 transition"
          >
            <Icon
              :name="showPassword ? 'mdi:eye-off-outline' : 'mdi:eye-outline'"
              size="18"
            />
          </button>
        </div>
      </div>

      <button
        type="submit"
        :disabled="isLoading"
        class="w-full py-2.5 rounded-xl text-sm font-medium
               bg-linear-to-r from-emerald-400 to-green-600 text-white
               hover:opacity-90 active:scale-95
               transition shadow-lg shadow-emerald-500/20
               disabled:opacity-50 disabled:cursor-not-allowed"
      >
        <span v-if="isLoading">Вход...</span>
        <span v-else>Войти</span>
      </button>

      <div class="text-center text-sm text-gray-500">
        Нет аккаунта?
        <NuxtLink to="/register" class="text-emerald-600 hover:underline">
          Зарегистрироваться
        </NuxtLink>
      </div>

    </form>
  </div>
</div>
</template>