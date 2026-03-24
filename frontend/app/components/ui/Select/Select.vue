<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'

interface Option {
  id: string | number
  name: string
}

const props = defineProps<{
  modelValue: Option | null
  options: Option[]
  placeholder?: string
  icon?: string
  disabled?: boolean
}>()

const emit = defineEmits(['update:modelValue'])

const isOpen = ref(false)
const rootRef = ref<HTMLElement | null>(null)

const selectedLabel = computed(() => {
  return props.modelValue?.name || props.placeholder || 'Выберите значение'
})

const select = (option: Option) => {
  emit('update:modelValue', option)
  isOpen.value = false
}

const handleClickOutside = (e: MouseEvent) => {
  if (!rootRef.value?.contains(e.target as Node)) {
    isOpen.value = false
  }
}

onMounted(() => {
  window.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  window.removeEventListener('click', handleClickOutside)
})
</script>

<template>
<div ref="rootRef" class="relative w-full">

  <div
    @click="!props.disabled && (isOpen = !isOpen)"
    class="w-full px-3 py-2 rounded-xl
           bg-white/70 backdrop-blur
           border border-emerald-100
           flex items-center
           cursor-pointer
           transition"
    :class="[
      icon ? 'gap-3' : 'justify-between',
      props.disabled ? 'opacity-60 cursor-not-allowed' : 'hover:bg-emerald-50'
    ]"
  >
    <div class="flex items-center gap-3 min-w-0 flex-1">
      <Icon
        v-if="icon"
        :name="icon"
        size="18"
        class="text-gray-400 shrink-0"
      />
      <span class="text-sm text-gray-700 truncate">
        {{ selectedLabel }}
      </span>
    </div>

    <Icon
      name="ph:caret-down"
      size="18"
      class="text-gray-400 transition shrink-0"
      :class="{ 'rotate-180': isOpen }"
    />
  </div>

  <Transition name="fade">
    <div
      v-if="isOpen"
      class="absolute z-50 mt-2 w-full
             rounded-xl
             bg-white/90 backdrop-blur-xl
             border border-emerald-100
             shadow-lg shadow-emerald-500/10
             max-h-60 overflow-y-auto"
    >
      <div
        v-for="option in options"
        :key="option.id"
        @click="select(option)"
        class="px-3 py-2 text-sm
               cursor-pointer
               hover:bg-emerald-50
               transition
               flex items-center justify-between"
      >
        <span class="text-gray-700">
          {{ option.name }}
        </span>

        <Icon
          v-if="modelValue?.id === option.id"
          name="ph:check"
          size="16"
          class="text-emerald-500"
        />
      </div>
    </div>
  </Transition>

</div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: all 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-5px);
}
</style>