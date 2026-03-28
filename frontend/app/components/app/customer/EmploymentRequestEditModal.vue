<script setup lang="ts">
import Select from "@/components/ui/Select/Select.vue";
import ModalWindow from "@/components/common/ModalWindow.vue";
import type { EmploymentRequest } from "@/types/frontend/employment";
import { useUniversityDepartments } from "@/composables/api/useUniversityDepartments";
import { useUpdateEmploymentRequest } from "@/composables/api/employment/useEmploymentRequests";

const props = defineProps<{
	modelValue: boolean;
	request: EmploymentRequest | null;
}>();

const emit = defineEmits<{
	"update:modelValue": [value: boolean];
	saved: [value: any];
}>();

const { fetchDepartments } = useUniversityDepartments();
const { update, isLoading, errorMessage, reset } = useUpdateEmploymentRequest();

const selectedDepartment = ref<{ id: string; name: string } | null>(null);

const form = ref({
	title: "",
	description: "",
	requirements: "",
	salary: "",
	schedule: "",
	max_participants: 10,
});

watch(
	() => props.request,
	(request) => {
		if (request) {
			form.value = {
				title: request.title || "",
				description: request.description || "",
				requirements: request.requirements || "",
				salary: request.salary || "",
				schedule: request.schedule || "",
				max_participants: request.max_participants || 10,
			};
			selectedDepartment.value = {
				id: request.university_department_id,
				name: `${request.university_name} / ${request.department_name}`,
			};
		}
	},
	{ immediate: true },
);

watch(
	() => props.modelValue,
	(val) => {
		if (val) {
			fetchDepartments();
		} else {
			reset();
		}
	},
);

const close = () => {
	emit("update:modelValue", false);
};

const handleSubmit = async () => {
	if (!props.request) return;

	const body = {
		title: form.value.title,
		description: form.value.description || undefined,
		requirements: form.value.requirements || undefined,
		salary: form.value.salary || undefined,
		schedule: form.value.schedule || undefined,
		max_participants: form.value.max_participants,
	};

	const success = await update(props.request.id, body);

	if (success) {
		emit("saved", {
			id: props.request!.id,
			...body,
		});
		close();
	}
};

const errors = computed(() => ({
	title: !form.value.title?.trim() ? "Введите название вакансии" : "",
	maxParticipants:
		form.value.max_participants <= 0 ? "Укажите количество участников" : "",
}));

const isValid = computed(
	() => form.value.title?.trim() && form.value.max_participants > 0,
);
</script>

<template>
  <ModalWindow
    :model-value="modelValue"
    @update:model-value="emit('update:modelValue', $event)"
    title="Редактирование заявки"
    width="max-w-lg"
  >
    <form @submit.prevent="handleSubmit" class="space-y-4">
      <div v-if="errorMessage" class="text-sm text-red-500 bg-red-50 border border-red-200 px-4 py-3 rounded-xl">
        {{ errorMessage }}
      </div>

      <div class="space-y-1">
        <label class="text-sm font-medium text-gray-700">Название вакансии <span class="text-red-500">*</span></label>
        <div class="relative">
          <Icon name="ph:text-t" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
          <input
            v-model="form.title"
            type="text"
            placeholder="Например: Стажер-разработчик"
            class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm
                   bg-white border border-emerald-300
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
            :class="{ 'border-red-400': errors.title }"
          />
        </div>
        <p v-if="errors.title" class="text-xs text-red-500">{{ errors.title }}</p>
      </div>

      <div class="space-y-1">
        <label class="text-sm font-medium text-gray-700">Количество участников <span class="text-red-500">*</span></label>
        <div class="relative">
          <Icon name="ph:users" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
          <input
            v-model.number="form.max_participants"
            type="number"
            min="1"
            class="w-full pl-9 pr-4 py-2.5 rounded-xl text-sm
                   bg-white border border-emerald-300
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
            :class="{ 'border-red-400': errors.maxParticipants }"
          />
        </div>
        <p v-if="errors.maxParticipants" class="text-xs text-red-500">{{ errors.maxParticipants }}</p>
      </div>

      <div class="space-y-1">
        <label class="text-sm font-medium text-gray-700">Описание</label>
        <textarea
          v-model="form.description"
          rows="2"
          placeholder="Опишите обязанности и условия работы"
          class="w-full px-4 py-2.5 rounded-xl text-sm resize-none
                 bg-white border border-emerald-300
                 focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
        ></textarea>
      </div>

      <div class="space-y-1">
        <label class="text-sm font-medium text-gray-700">Требования</label>
        <textarea
          v-model="form.requirements"
          rows="2"
          placeholder="Укажите необходимые навыки и квалификацию"
          class="w-full px-4 py-2.5 rounded-xl text-sm resize-none
                 bg-white border border-emerald-300
                 focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
        ></textarea>
      </div>

      <div class="grid grid-cols-2 gap-4">
        <div class="space-y-1">
          <label class="text-sm font-medium text-gray-700">Зарплата</label>
          <input
            v-model="form.salary"
            type="text"
            class="w-full px-4 py-2.5 rounded-xl text-sm
                   bg-white border border-emerald-300
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
          />
        </div>

        <div class="space-y-1">
          <label class="text-sm font-medium text-gray-700">График</label>
          <input
            v-model="form.schedule"
            type="text"
            placeholder="5/2, 8 часов"
            class="w-full px-4 py-2.5 rounded-xl text-sm
                   bg-white border border-emerald-300
                   focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
          />
        </div>
      </div>
    </form>

    <template #footer>
      <button
        @click="close"
        class="px-4 py-2 rounded-xl text-sm font-medium
               bg-white/70 border border-emerald-100
               hover:bg-emerald-50 transition"
      >
        Отмена
      </button>

      <button
        @click="handleSubmit"
        :disabled="isLoading || !isValid"
        class="px-4 py-2 rounded-xl text-sm font-medium
               bg-linear-to-r from-emerald-400 to-green-600 text-white
               hover:opacity-90 active:scale-95 transition
               shadow-lg shadow-emerald-500/20
               disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer"
      >
        <span v-if="isLoading">Сохранение...</span>
        <span v-else>Сохранить</span>
      </button>
    </template>
  </ModalWindow>
</template>
