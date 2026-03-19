<script setup lang="ts">
</script>

<template>
<div class="min-h-screen flex items-center justify-center p-4 sm:p-6 relative overflow-hidden">

  <div class="absolute inset-0 pointer-events-none">
    <div class="absolute -top-30 left-1/2 -translate-x-1/2 w-175 h-200
                bg-emerald-400/20 blur-3xl rounded-full" />
  </div>

  <div class="relative text-center">
    <div class="text-7xl sm:text-8xl font-bold
                bg-linear-to-r from-emerald-400 to-green-600
                bg-clip-text text-transparent mb-4">
      404
    </div>

    <div class="text-lg sm:text-xl text-gray-700 dark:text-gray-300 uppercase mb-3">
      страница не найдена
    </div>

    <p class="text-sm text-gray-500 font-semibold max-w-md mx-auto mb-6">
      Возможно, страница была удалена или вы перешли по неверной ссылке.
    </p>

    <NuxtLink
      to="/"
      class="inline-flex items-center gap-2 px-5 py-2.5 rounded-xl text-sm font-medium
             bg-linear-to-r from-emerald-400 to-green-600 text-white
             hover:opacity-90 active:scale-95
             transition shadow-lg shadow-emerald-500/20"
    >
      <Icon name="mdi:home-outline" class="w-4 h-4" />
      На главную
    </NuxtLink>

  </div>
</div>
</template>