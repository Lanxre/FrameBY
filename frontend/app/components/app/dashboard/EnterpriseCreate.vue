<script setup lang="ts">
import { useCreateEnterprise } from "@/composables/api/dashboard/useCreateEnterprise";

const emit = defineEmits<{
	created: [];
}>();

const { isLoading, errorMessage, isSuccess, create, reset } =
	useCreateEnterprise();

const name = ref("");
const address = ref("");

const handleSubmit = async () => {
	if (!name.value.trim()) {
		return;
	}

	const success = await create({
		name: name.value.trim(),
		address: address.value.trim() || undefined,
	});

	if (success) {
		name.value = "";
		address.value = "";
		emit("created");
		setTimeout(() => {
			reset();
		}, 3000);
	}
};
</script>

<template>
  <div class="p-4 rounded-xl bg-white/60 border border-emerald-100">
    <div class="flex items-center gap-2 mb-4">
      <Icon name="ph:buildings" size="20" class="text-emerald-500" />
      <h3 class="text-sm font-semibold text-gray-700">Добавить организацию</h3>
    </div>

    <div v-if="errorMessage" class="text-sm text-red-500 bg-red-50 border border-red-200 px-3 py-2 rounded-xl mb-3">
      {{ errorMessage }}
    </div>

    <div v-if="isSuccess" class="text-sm text-green-600 bg-green-50 border border-green-200 px-3 py-2 rounded-xl mb-3">
      Организация успешно создана
    </div>

    <div class="flex flex-col h-full gap-3">
      <div class="space-y-1">
        <label class="text-xs text-gray-500 ml-2">Название</label>
        <div class="relative">
          <Icon name="ph:building-office" size="16" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
          <input
            v-model="name"
            type="text"
            placeholder="ООО Яндекс"
            class="w-full pl-8 pr-3 py-2 rounded-xl text-sm
                   bg-white border border-emerald-100
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
          />
        </div>
      </div>

      <div class="space-y-1">
        <label class="text-xs text-gray-500 ml-2">Адрес</label>
        <div class="relative">
          <Icon name="ph:map-pin" size="16" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
          <input
            v-model="address"
            type="text"
            placeholder="Москва, ул. Льва Толстого 16"
            class="w-full pl-8 pr-3 py-2 rounded-xl text-sm
                   bg-white border border-emerald-100
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
          />
        </div>
      </div>

      <button
         class="py-2 mt-30 rounded-xl text-sm font-bold
                bg-linear-to-r from-emerald-400 to-green-600 text-white
                hover:opacity-90 transition
                disabled:opacity-50 disabled:cursor-not-allowed
                shadow-sm shadow-emerald-500/20"
         :disabled="isLoading || !name.trim()"
         @click="handleSubmit"
       >
        <span v-if="isLoading">Создание...</span>
        <span v-else>Создать</span>
      </button>
    </div>
  </div>
</template>
