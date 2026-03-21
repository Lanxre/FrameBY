<script setup lang="ts">
import { computed } from 'vue'

type Position = 'top' | 'bottom' | 'left' | 'right'

const props = withDefaults(defineProps<{
  text: string
  position?: Position
}>(), {
  position: 'top'
})

const positionClasses = computed(() => {
  switch (props.position) {
    case 'bottom':
      return 'top-full mt-2 left-1/2 -translate-x-1/2'
    case 'left':
      return 'right-full mr-2 top-1/2 -translate-y-1/2'
    case 'right':
      return 'left-full ml-2 top-1/2 -translate-y-1/2'
    default:
      return 'bottom-full mb-2 left-1/2 -translate-x-1/2'
  }
})

const arrowClasses = computed(() => {
  switch (props.position) {
    case 'bottom':
      return 'top-0 left-1/2 -translate-x-1/2 -translate-y-1/2'
    case 'left':
      return 'right-0 top-1/2 translate-x-1/2 -translate-y-1/2'
    case 'right':
      return 'left-0 top-1/2 -translate-x-1/2 -translate-y-1/2'
    default:
      return 'bottom-0 left-1/2 -translate-x-1/2 translate-y-1/2'
  }
})
</script>

<template>
  <div class="relative group inline-flex">

    <slot />
    
    <div
      class="absolute z-50 px-2 py-1 rounded-md text-xs whitespace-nowrap
             bg-black/80 text-white
             opacity-0 scale-95 pointer-events-none
             group-hover:opacity-100 group-hover:scale-100
             transition duration-200 ease-out"
      :class="positionClasses"
    >
      {{ text }}
      
      <div
        class="absolute w-2 h-2 bg-black/80 rotate-45"
        :class="arrowClasses"
      />
    </div>

  </div>
</template>