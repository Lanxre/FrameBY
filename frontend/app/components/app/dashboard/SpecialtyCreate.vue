<script setup lang="ts">
import { useCreateSpecialty } from "@/composables/api/dashboard/useCreateSpecialty";
import { useUniversityDepartments } from "@/composables/api/useUniversityDepartments";
import Select from "@/components/ui/Select/Select.vue";

const emit = defineEmits<{
	created: [];
}>();

const { isLoading, errorMessage, isSuccess, create, reset } =
	useCreateSpecialty();

const { departments, isLoading: isLoadingDepartments, fetchDepartments } =
	useUniversityDepartments();

const selectedDepartment = ref<{ id: string; departmentId: string; name: string } | null>(null);
const specialtyName = ref("");

onMounted(() => {
	fetchDepartments();
});

const handleSubmit = async () => {
	if (!specialtyName.value.trim() || !selectedDepartment.value) {
		return;
	}

	const success = await create({
		name: specialtyName.value.trim(),
		department_id: selectedDepartment.value.departmentId,
	});

	if (success) {
		specialtyName.value = "";
		selectedDepartment.value = null;
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
      <Icon name="ph:book-open" size="20" class="text-emerald-500" />
      <h3 class="text-sm font-semibold text-gray-700">Добавить специальность</h3>
    </div>

    <div v-if="errorMessage" class="text-sm text-red-500 bg-red-50 border border-red-200 px-3 py-2 rounded-xl mb-3">
      {{ errorMessage }}
    </div>

    <div v-if="isSuccess" class="text-sm text-green-600 bg-green-50 border border-green-200 px-3 py-2 rounded-xl mb-3">
      Специальность успешно создана
    </div>

    <div class="flex flex-col gap-3">
      <div class="space-y-1">
        <label class="text-xs text-gray-500 ml-2">Университет / Кафедра</label>
        <Select
          v-model="selectedDepartment"
          :options="departments.map(dp => ({ id: dp.id, departmentId: dp.department_id, name: `${dp.university_name} / ${dp.department_name}` }))"
          placeholder="Выберите кафедру"
          icon="ph:graduation-cap"
          :disabled="isLoadingDepartments"
        />
      </div>

      <div class="space-y-1">
        <label class="text-xs text-gray-500 ml-2">Название специальности</label>
        <div class="relative">
          <Icon name="ph:book-open" size="16" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
          <input
            v-model="specialtyName"
            type="text"
            placeholder="Прикладная информатика"
            class="w-full pl-8 pr-3 py-2 rounded-xl text-sm
                   bg-white border border-emerald-100
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
          />
        </div>
      </div>

      <button
        class="py-2 rounded-xl text-sm font-bold
               bg-linear-to-r from-emerald-400 to-green-600 text-white
               hover:opacity-90 transition
               disabled:opacity-50 disabled:cursor-not-allowed
               shadow-sm shadow-emerald-500/20"
        :disabled="isLoading || !specialtyName.trim() || !selectedDepartment"
        @click="handleSubmit"
      >
        <span v-if="isLoading">Создание...</span>
        <span v-else>Создать</span>
      </button>
    </div>
  </div>
</template>
