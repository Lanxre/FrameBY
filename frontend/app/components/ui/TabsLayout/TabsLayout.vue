<script setup lang="ts">
import { ref, computed } from 'vue'
import type { TabItem } from '@/types/frontend/tabs-layout';

const props = defineProps<{
  tabs: TabItem[]
  defaultTab?: string
  activeTabClass?: string
}>()

const activeTab = ref(props.defaultTab || props.tabs[0]?.value)

const activeComponent = computed(() => {
  return props.tabs.find(t => t.value === activeTab.value)?.component
})
</script>

<template>
<div class="w-full max-w-7xl mx-auto px-4 py-6">
  
  <div class="grid grid-cols-10 gap-6">
    <div class="col-span-3
                rounded-2xl
                bg-white/80 backdrop-blur-xl
                border border-emerald-100
                shadow-lg shadow-emerald-500/10
                p-3 space-y-1">

      <div
        v-for="tab in tabs"
        :key="tab.value"
        @click="activeTab = tab.value"
        class="flex items-center gap-2 px-3 py-2 rounded-xl
               text-sm font-semibold
               cursor-pointer transition"
        :class="activeTab === tab.value
          ? (activeTabClass || 'bg-emerald-50 text-emerald-600')
          : 'text-gray-600 hover:bg-emerald-50/60 hover:text-gray-700'"
      >
        <div v-if="tab.hasPermission" class="flex items-center gap-2">
            <Icon v-if="tab.icon" :name="tab.icon" size="18" />
            <span>{{ tab.label }}</span>
        </div>
      </div>

    </div>

    <div class="col-span-7
                rounded-2xl
                bg-white/80 backdrop-blur-xl
                border border-emerald-100
                shadow-lg shadow-emerald-500/10
                p-6">

      <component :is="activeComponent" />

    </div>

  </div>

</div>
</template>