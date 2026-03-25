<script setup lang="ts">
defineProps<{
  currentPage: number
  totalPages: number
}>()

const emit = defineEmits<{
  pageChange: [page: number]
}>()
</script>

<template>
  <div class="flex justify-center items-center gap-3">
    <button
      @click="emit('pageChange', currentPage - 1)"
      :disabled="currentPage === 1"
      class="w-10 h-10 rounded-xl border border-emerald-100 bg-white/70 hover:bg-emerald-50 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center transition-colors"
    >
      <Icon name="ph:caret-left" size="18" />
    </button>

    <span class="text-sm text-gray-600">
      {{ currentPage }} / {{ totalPages }}
    </span>

    <button
      @click="emit('pageChange', currentPage + 1)"
      :disabled="currentPage === totalPages"
      class="w-10 h-10 rounded-xl border border-emerald-100 bg-white/70 hover:bg-emerald-50 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center transition-colors"
    >
      <Icon name="ph:caret-right" size="18" />
    </button>
  </div>
</template>
