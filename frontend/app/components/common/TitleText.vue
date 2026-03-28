<script setup lang="ts">
const props = defineProps<{
	title: string;
	subtitle?: string;
	center?: boolean;
}>();
</script>

<template>
  <div :class="['w-full', center ? 'text-center' : 'text-left']">
    <h1
      class="text-2xl sm:text-3xl font-semibold uppercase
             bg-linear-to-r from-emerald-400 to-green-600
             bg-clip-text text-transparent"
    >
      {{ title }}
    </h1>
    <p
      v-if="subtitle"
      class="mt-2 text-sm sm:text-base text-gray-500 max-w-6xl"
      :class="center ? 'mx-auto' : ''"
    >
      {{ subtitle }}
    </p>

  </div>
</template>