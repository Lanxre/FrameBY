<script lang="ts" setup>
import { toRef, computed } from "vue";
import { useModalLogic } from "@/composables/ui/useModal";

interface Props {
	modelValue: boolean;
	title?: string;

	width?: string;
	height?: string;

	customWidth?: string;
	customHeight?: string;

	centerTitle?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
	modelValue: false,
	width: "max-w-lg",
	centerTitle: false,
});

const emit = defineEmits(["update:modelValue", "close"]);

const close = () => {
	emit("update:modelValue", false);
	emit("close");
};

useModalLogic(toRef(props, "modelValue"), close);

/**
 * 🔥 inline style (если передан кастом)
 */
const modalStyle = computed(() => ({
	width: props.customWidth || undefined,
	height: props.customHeight || undefined,
}));
</script>

<template>
<Teleport to="body">
  <Transition name="modal">
    <div v-if="modelValue" class="relative z-100">
      <div
        class="fixed inset-0 bg-black/20 backdrop-blur-sm"
        @click="close"
      />

      <div class="fixed inset-0 flex items-center justify-center p-4">
        
        <div
          class="relative w-full transform transition-all"
          :class="!customWidth ? width : ''"
          :style="modalStyle"
          @click.stop
        >

          <div class="rounded-2xl
                      bg-white/80 backdrop-blur-xl
                      border border-emerald-100
                      shadow-xl shadow-emerald-500/10
                      flex flex-col">

            <div
              class="flex items-center px-6 py-4 border-b border-emerald-100 relative"
              :class="centerTitle ? 'justify-center' : 'justify-between'"
            >
              <h3 class="text-lg font-semibold text-gray-800 cursor-default">
                {{ title }}
              </h3>

              <button
                @click="close"
                class="text-gray-400 hover:text-emerald-500 transition cursor-pointer"
                :class="centerTitle ? 'absolute right-6' : ''"
              >
                <Icon name="ph:x" size="20" />
              </button>
            </div>


            <div
              class="px-6 py-5 text-sm text-gray-700 overflow-auto flex-1"
              :class="height"
            >
              <slot />
            </div>

            <div
              v-if="$slots.footer"
              class="px-6 py-4 border-t border-emerald-100 flex justify-end gap-2"
            >
              <slot name="footer" />
            </div>

          </div>

        </div>

      </div>
    </div>
  </Transition>
</Teleport>
</template>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.25s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}

.modal-enter-active .transform,
.modal-leave-active .transform {
  transition: all 0.25s ease;
}

.modal-enter-from .transform {
  opacity: 0;
  transform: translateY(8px) scale(0.96);
}

.modal-leave-to .transform {
  opacity: 0;
  transform: translateY(8px) scale(0.96);
}
</style>