<script setup lang="ts">
import { ref } from 'vue'
import { onClickOutside } from '@vueuse/core'

interface DropdownItem {
  label: string
  to: string
  icon?: string
}

const props = defineProps<{
  items: DropdownItem[]
}>()

const isOpen = ref(false)
const dropdownRef = ref<HTMLElement | null>(null)

const toggle = () => (isOpen.value = !isOpen.value)
const close = () => (isOpen.value = false)

onClickOutside(dropdownRef, close)
</script>

<template>
  <div ref="dropdownRef" class="relative">
    <button
      @click="toggle"
      class="px-3 py-2 rounded-xl text-sm font-medium
             text-gray-700 dark:text-gray-200
             transition"
    >
      <slot name="trigger" />
    </button>

    <transition
      enter-active-class="transition duration-150 ease-out"
      enter-from-class="opacity-0 translate-y-2 scale-95"
      enter-to-class="opacity-100 translate-y-0 scale-100"
      leave-active-class="transition duration-100 ease-in"
      leave-from-class="opacity-100 translate-y-0 scale-100"
      leave-to-class="opacity-0 translate-y-2 scale-95"
    >
      <div
        v-if="isOpen"
        class="absolute left-0 mt-2 w-100 rounded-xl z-50
               glass-strong shadow-xl"
      >
        <ul class="py-2">
          <li v-for="item in items" :key="item.label">
            <NuxtLink
              :to="item.to"
              @click="close"
              class="flex items-center gap-2 px-4 py-2 text-sm"
            >
              <Icon
                v-if="item.icon"
                :name="item.icon"
      ~          class="w-6 h-6 opacity-70 shrink-0"
              />
              <span class="hover:text-emerald-600/70">
                  {{ item.label }}
              </span>
            </NuxtLink>
          </li>
        </ul>
      </div>
    </transition>
  </div>
</template>