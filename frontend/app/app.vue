<script setup lang="ts">
import Notifications from '@/components/ui/Notifications/Notifications.vue';
</script>

<template>
  <div class="min-h-screen flex flex-col">
    <Notifications />
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </div>
</template>