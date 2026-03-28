<script setup lang="ts">
import { useBrsmForm } from "@/composables/api/forms/useBrsmForm";
import ModalConfirm from "~/components/common/ModalConfirm.vue";
const emit = defineEmits(["submit"]);

const { form, isLoading, errorMessage, showConfirm, handleSubmit, submit } =
	useBrsmForm();
</script>

<template>
<div class="max-w-xl min-w-120 mx-auto p-6
            rounded-2xl
            bg-white/80 backdrop-blur-xl
            border border-emerald-100
            shadow-lg shadow-emerald-500/10">

  <div class="flex items-center gap-2 mb-5">
    <Icon name="ph:users-three" size="22" class="text-emerald-500" />
    <h2 class="text-lg font-semibold text-gray-800 uppercase">
      Представитель БРСМ
    </h2>
  </div>

  <div v-if="errorMessage"
       class="text-sm text-red-500 bg-red-50 border border-red-200 px-3 py-2 rounded-xl mb-4">
    {{ errorMessage }}
  </div>

  <div class="space-y-4">

    <div class="space-y-1">
      <label class="text-xs text-gray-500 ml-2">ФИО</label>
      <div class="relative">
        <Icon name="ph:user" size="18"
              class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
        <input
          v-model="form.fullName"
          type="text"
          placeholder="Иванов Иван Иванович"
          class="w-full pl-9 pr-3 py-2 rounded-xl
                 bg-white/60 border border-emerald-100
                 focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                 text-sm"
        />
      </div>
    </div>

    <div class="space-y-1">
      <label class="text-xs text-gray-500 ml-2">Должность</label>
      <div class="relative">
        <Icon name="ph:briefcase" size="18"
              class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
        <input
          v-model="form.position"
          type="text"
          placeholder="Должность"
          class="w-full pl-9 pr-3 py-2 rounded-xl
                 bg-white/60 border border-emerald-100
                 focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                 text-sm"
        />
      </div>
    </div>

    <div class="space-y-1">
      <label class="text-xs text-gray-500 ml-2">Телефон</label>
      <div class="relative">
        <Icon name="ph:phone" size="18"
              class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
        <input
          v-model="form.phone"
          type="tel"
          placeholder="+375 (__) ___-__-__"
          class="w-full pl-9 pr-3 py-2 rounded-xl
                 bg-white/60 border border-emerald-100
                 focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                 text-sm"
        />
      </div>
    </div>

    <button
      @click="handleSubmit"
      :disabled="isLoading"
      class="w-full py-2.5 rounded-xl text-sm font-semibold
             bg-linear-to-r from-emerald-400 to-green-600 text-white
             hover:opacity-90 active:scale-95
             transition shadow-lg shadow-emerald-500/20
             disabled:opacity-50"
    >
      <span v-if="isLoading">Сохранение...</span>
      <span v-else>Сохранить</span>
    </button>
  </div>
  <ModalConfirm
    v-model="showConfirm"
    title="Сохранение изменений"
    description="Вы уверены, что хотите отправить данные? Они будут расмотрены в течение 24 часов."
    confirmText="Сохранить"
    cancelText="Отмена"
    :loading="isLoading"
    @confirm="submit"
  />
</div>
</template>