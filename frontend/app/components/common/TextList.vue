<script setup lang="ts">
interface ListItem {
  title?: string
  content: string
}

const props = defineProps<{
  items: ListItem[]
}>()
</script>

<template>
<div class="space-y-6">

  <div
    v-for="(item, index) in items"
    :id="`item-${index + 1}`"
    :key="index"
    class="group p-5 rounded-2xl
           bg-white/60 dark:bg-white/5
           backdrop-blur
           border border-black/5 dark:border-white/10
           hover:border-emerald-400/40
           hover:shadow-md
           transition"
  >

    <div class="flex items-start gap-3 mb-3">

      <div class="flex items-center justify-center min-w-8 h-8 rounded-lg
                  bg-linear-to-r from-emerald-400 to-green-600
                  text-white text-sm font-semibold shadow">
        {{ index + 1 }}
      </div>

      <div v-if="item.title"
           class="font-medium text-gray-800 dark:text-gray-200
                  group-hover:text-emerald-600 transition">
        {{ item.title }}
      </div>

    </div>

    <p class="text-sm leading-relaxed text-gray-600 dark:text-gray-300 whitespace-pre-line wrap-break-words">
      {{ item.content }}
    </p>

  </div>

</div>
</template>