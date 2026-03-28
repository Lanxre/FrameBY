<script setup lang="ts">
import type { PropType } from "vue";

type ButtonVariant =
	| "primary"
	| "secondary"
	| "outline"
	| "border-shadow"
	| "ghost"
	| "danger"
	| "success";

const props = defineProps({
	variant: {
		type: String as PropType<ButtonVariant>,
		default: "primary",
	},
	size: {
		type: String as PropType<"sm" | "md" | "lg">,
		default: "md",
	},
	disabled: { type: Boolean, default: false },
	loading: { type: Boolean, default: false },
	type: {
		type: String as PropType<"button" | "submit" | "reset">,
		default: "button",
	},
});

const emit = defineEmits<(e: "click", event: MouseEvent) => void>();

const handleClick = (event: MouseEvent) => {
	if (!props.disabled && !props.loading) {
		emit("click", event);
	}
};
</script>

<template>
  <button
    :type="type"
    :disabled="disabled || loading"
    @click="handleClick"
    class="inline-flex items-center justify-center font-medium transition-all duration-200 rounded-2xl focus:outline-none focus:ring-2 focus:ring-offset-2 active:scale-95 cursor-pointer disabled:cursor-not-allowed"
    :class="[
      {
        primary:       'bg-emerald-600 hover:bg-emerald-700 text-white shadow-sm focus:ring-emerald-500',
        secondary:     'bg-gray-700 hover:bg-gray-800 text-white shadow-sm focus:ring-gray-500',
        outline:       'border border-gray-400 hover:bg-gray-50 text-gray-700 focus:ring-gray-500',
        
        'border-shadow': 'bg-white border-2 border-emerald-600 text-emerald-700 shadow-md hover:shadow-xl hover:-translate-y-0.5 focus:ring-emerald-500 active:shadow-md',

        ghost:         'hover:bg-gray-100 text-gray-700 focus:ring-gray-500',
        danger:        'bg-red-600 hover:bg-red-700 text-white focus:ring-red-500',
        success:       'bg-green-600 hover:bg-green-700 text-white focus:ring-green-500',
      }[variant],

      {
        sm: 'px-5 py-2.5 text-sm',
        md: 'px-7 py-3.5 text-base',
        lg: 'px-9 py-4 text-lg',
      }[size],

      (disabled || loading) && 'opacity-50'
    ]"
  >
    <span v-if="loading" class="animate-spin w-5 h-5 border-2 border-current border-t-transparent rounded-full mr-2" />

    <slot />
  </button>
</template>