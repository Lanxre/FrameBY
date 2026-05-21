<script setup lang="ts" >
import Tooltip from "@/components/ui/ToolTip.vue";

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
<div v-if="tooltipText">
    <Tooltip :text="tooltipText">
        <div class="w-full h-full flex flex-col p-4 rounded-xl bg-white/70 border border-emerald-100">

        <p class="text-gray-400 text-xs mb-1 flex items-center gap-1">
          <Icon :name="icon" size="14" />
          {{ title }}
        </p>
        <p class="flex items-end text-emerald-700 font-semibold w-56 h-12 overflow-hidden" :class="uppercase ? 'uppercase' : ''">
          {{ text }}
        </p>
      </div>
    </Tooltip>
</div>
<div v-else class="w-full h-full flex flex-col justify-between p-4 rounded-xl bg-white/70 border border-emerald-100">
  <p class="text-gray-400 text-xs mb-1 flex items-center gap-1">
    <Icon :name="icon" size="14" />
    {{ title }}
  </p>
  <p class="text-emerald-600 font-semibold" :class="uppercase ? 'uppercase' : ''">
    {{ text }}
  </p>
</div>
</template>
