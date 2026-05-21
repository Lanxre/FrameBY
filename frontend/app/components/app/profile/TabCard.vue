<script setup lang="ts" >

interface Props {
  title: string;
  tooltipText?: string;
  text?: string;
  icon: string;
  uppercase?: boolean;
}

withDefaults(defineProps<Props>(), {
  uppercase: false,
});

</script>

<template>
<div class="h-full flex flex-col justify-between gap-1 p-4 rounded-xl bg-white/70 border border-emerald-100 relative group">
  <p class="text-gray-400 text-xs flex items-center gap-1">
    <Icon :name="icon" size="14" />
    {{ title }}
  </p>
  <p class="font-semibold wrap-break-word" :class="[uppercase ? 'uppercase' : '', tooltipText ? 'text-emerald-700' : 'text-emerald-600']">
    {{ text }}
  </p>

  <div v-if="tooltipText"
       class="absolute z-50 px-2 py-1 rounded-md text-xs whitespace-nowrap
              bg-black/80 text-white
              opacity-0 scale-95 pointer-events-none
              group-hover:opacity-100 group-hover:scale-100
              transition duration-200 ease-out
              bottom-full mb-2 left-1/2 -translate-x-1/2">
    {{ tooltipText }}
    <div class="absolute w-2 h-2 bg-black/80 rotate-45 bottom-0 left-1/2 -translate-x-1/2 translate-y-1/2" />
  </div>
</div>
</template>
