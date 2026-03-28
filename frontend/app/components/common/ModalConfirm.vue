<script setup lang="ts">
import ModalWindow from "./ModalWindow.vue";

const props = defineProps<{
	modelValue: boolean;
	title?: string;
	description?: string;
	confirmText?: string;
	cancelText?: string;
	loading?: boolean;
}>();

const emit = defineEmits(["update:modelValue", "confirm", "cancel"]);

const close = () => {
	emit("update:modelValue", false);
	emit("cancel");
};

const confirm = () => {
	emit("confirm");
};
</script>

<template>
  <ModalWindow
    :model-value="modelValue"
    @update:modelValue="emit('update:modelValue', $event)"
    :title="title || 'Подтверждение'"
    width="max-w-md"
    centerTitle
  >
    <div class="text-center space-y-4">
      <div class="flex justify-center">
        <div class="w-14 h-14 rounded-full
                    bg-linear-to-br from-emerald-400 to-green-600
                    flex items-center justify-center
                    shadow-lg shadow-emerald-500/20">
          <Icon name="ph:warning" size="26" class="text-white" />
        </div>
      </div>

      <div>
        <p class="text-sm text-gray-600">
          {{ description || 'Вы уверены, что хотите выполнить это действие?' }}
        </p>
      </div>

    </div>

    <template #footer>
      <button
        @click="close"
        class="px-4 py-2 rounded-xl text-sm font-medium
               bg-white/70
               border border-emerald-100
               hover:bg-emerald-50
               transition"
      >
        {{ cancelText || 'Отмена' }}
      </button>

      <button
        @click="confirm"
        :disabled="loading"
        class="px-4 py-2 rounded-xl text-sm font-medium
               bg-linear-to-r from-emerald-400 to-green-600 text-white
               hover:opacity-90 active:scale-95
               transition shadow-lg shadow-emerald-500/20
               disabled:opacity-50 disabled:cursor-not-allowed"
      >
        <span v-if="loading">Загрузка...</span>
        <span v-else class="cursor-pointer">{{ confirmText || 'Подтвердить' }}</span>
      </button>
    </template>
  </ModalWindow>
</template>