<script lang="ts" setup>
import { computed, onUnmounted, ref } from "vue";
import type { NotificationItem } from "@/types/frontend/notification";

const props = defineProps<{
	notification: NotificationItem;
}>();

const emit = defineEmits(["remove"]);

const el = ref<HTMLElement | null>(null);
const startX = ref(0);
const currentX = ref(0);
const isDragging = ref(false);

const getIconColor = (type?: string) => {
	switch (type) {
		case "success":
			return "text-emerald-500";
		case "error":
			return "text-red-500";
		case "warning":
			return "text-amber-500";
		default:
			return "text-blue-500";
	}
};

const onStart = (clientX: number) => {
	startX.value = clientX;
	isDragging.value = true;
};

const onMove = (clientX: number) => {
	if (!isDragging.value) return;
	const delta = clientX - startX.value;
	if (delta > 0) {
		currentX.value = delta;
	}
};

const onEnd = () => {
	if (!isDragging.value) return;
	isDragging.value = false;

	if (currentX.value > 100) {
		emit("remove", props.notification.id);
	} else {
		currentX.value = 0;
	}
};

const handleTouchStart = (e: TouchEvent) => onStart(e.touches[0]!.clientX);
const handleTouchMove = (e: TouchEvent) => onMove(e.touches[0]!.clientX);
const handleTouchEnd = () => onEnd();

const handleMouseDown = (e: MouseEvent) => {
	onStart(e.clientX);
	window.addEventListener("mousemove", handleMouseMove);
	window.addEventListener("mouseup", handleMouseUp);
};

const handleMouseMove = (e: MouseEvent) => onMove(e.clientX);
const handleMouseUp = () => {
	onEnd();
	window.removeEventListener("mousemove", handleMouseMove);
	window.removeEventListener("mouseup", handleMouseUp);
};

onUnmounted(() => {
	window.removeEventListener("mousemove", handleMouseMove);
	window.removeEventListener("mouseup", handleMouseUp);
});

const style = computed(() => ({
	transform: `translateX(${currentX.value}px)`,
	opacity: 1 - currentX.value / 300,
	transition: isDragging.value
		? "none"
		: "transform 0.3s ease-out, opacity 0.3s ease-out",
	touchAction: "pan-y",
}));
</script>

<template>
<div 
  ref="el"
  class="pointer-events-auto w-full
         bg-white backdrop-blur-xl
         border border-emerald-100
         rounded-xl
         shadow-lg shadow-emerald-500/10
         p-4 flex items-start gap-3
         cursor-grab active:cursor-grabbing select-none
         transition"
  :style="style"
  @touchstart="handleTouchStart"
  @touchmove="handleTouchMove"
  @touchend="handleTouchEnd"
  @mousedown="handleMouseDown"
>
  <div class="shrink-0 mt-0.5">
    <svg v-if="notification.type === 'success'"
         :class="getIconColor(notification.type)"
         class="w-5 h-5"
         fill="none" viewBox="0 0 24 24" stroke="currentColor">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
            d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
    </svg>

    <svg v-else-if="notification.type === 'error'"
         :class="getIconColor(notification.type)"
         class="w-5 h-5"
         fill="none" viewBox="0 0 24 24" stroke="currentColor">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
            d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
    </svg>

    <svg v-else-if="notification.type === 'warning'"
         :class="getIconColor(notification.type)"
         class="w-5 h-5"
         fill="none" viewBox="0 0 24 24" stroke="currentColor">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
            d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
    </svg>

    <svg v-else
         :class="getIconColor(notification.type)"
         class="w-5 h-5"
         fill="none" viewBox="0 0 24 24" stroke="currentColor">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
            d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
    </svg>
  </div>

  <div class="flex-1 w-0">
    <p class="text-sm font-semibold text-black">
      {{ notification.title }}
    </p>

    <p v-if="notification.content"
       class="mt-1 text-sm text-black">
      {{ notification.content }}
    </p>
  </div>

  <div class="shrink-0 flex ml-4">
    <button 
      @mousedown.stop
      @click.stop="emit('remove', notification.id)"
      class="rounded-lg p-1
             text-gray-400 hover:text-emerald-600
             hover:bg-emerald-50
             transition cursor-pointer"
    >
      <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M6 18L18 6M6 6l12 12" />
      </svg>
    </button>
  </div>
</div>
</template>