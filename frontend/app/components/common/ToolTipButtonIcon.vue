<script setup lang="ts">
import ToolTip from '../ui/ToolTip.vue';

interface Props {
  text: string;
  position: 'top' | 'bottom' | 'left' | 'right';
  
  icon: string;
  iconSize: string;

  onClick?: () => void;
}

defineProps<Props>();
</script>

<template>
  <ToolTip :text="text" :position="position">
    <button @click="onClick" class="p-1.5 text-gray-400 hover:text-zinc-600 transition rounded-lg cursor-pointer">
        <Icon :name="icon" :size="iconSize" />
    </button>
  </ToolTip>
</template>