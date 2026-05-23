<script setup lang="ts">
import type { ProfileRow } from "@/types/dashboard/profile";
import Select from "@/components/ui/Select/Select.vue";
import { useProfileEdit } from "@/composables/api/dashboard/useProfileEdit";
import { formatAvatar } from "@/utils/str";
import { FramebyAppRole } from "~/types/frontend/enums/role";

const props = defineProps<{
	modelValue: boolean;
	profile: ProfileRow | null;
}>();

const emit = defineEmits<{
	"update:modelValue": [value: boolean];
	saved: [];
}>();

const {
	roleOptions,
	departments,
	specialties,
	selectedRole,
	selectedUniversity,
	fullName,
	subrole,
	specialtyId,
	specialtyName,
	grade,
	position,
	phone,
	isLoading,
	errorMessage,
	populateForm,
	resetForm,
	save,
	fetchByUniversityDepartment,
} = useProfileEdit();

const selectedSpecialty = ref<{ id: string; name: string } | null>(null);

watch(selectedSpecialty, (val) => {
	specialtyId.value = val?.id || null;
});

watch(
	() => selectedUniversity.value?.id,
	(val) => {
		selectedSpecialty.value = null;
		specialtyId.value = null;
		if (val) {
			fetchByUniversityDepartment(val);
		}
	},
);

watch(
	[() => props.modelValue, () => specialties.value],
	([open]) => {
		if (open && specialtyId.value && specialties.value.length > 0) {
			const match = specialties.value.find((s) => s.id === specialtyId.value);
			if (match) selectedSpecialty.value = { id: match.id, name: match.name };
		}
	},
	{ immediate: true },
);

watch(
	() => props.profile,
	(newProfile) => {
		if (newProfile) {
			populateForm(newProfile);
		}
	},
	{ immediate: true },
);

watch(
	() => props.modelValue,
	(isOpen) => {
		if (!isOpen) {
			resetForm();
		}
	},
);

const specialtyOptions = computed(() =>
	specialties.value.map((s) => ({ id: s.id, name: s.name })),
);

const close = () => {
	emit("update:modelValue", false);
};

const handleSave = async () => {
	if (!props.profile) return;

	const success = await save(props.profile.user_id);
	if (success) {
		emit("saved");
		close();
	}
};
</script>

<template>
  <Teleport to="body">
    <Transition
      enter-active-class="transition duration-200 ease-out"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition duration-150 ease-in"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div 
        v-if="modelValue"
        class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/30 backdrop-blur-sm"
        @click.self="close"
      >
        <Transition
          enter-active-class="transition duration-200 ease-out"
          enter-from-class="transform scale-95 opacity-0"
          enter-to-class="transform scale-100 opacity-100"
          leave-active-class="transition duration-150 ease-in"
          leave-from-class="transform scale-100 opacity-100"
          leave-to-class="transform scale-95 opacity-0"
        >
          <div 
            v-if="modelValue"
            class="w-full max-w-lg rounded-2xl bg-white/95 backdrop-blur-xl 
                   border border-emerald-100 shadow-xl shadow-emerald-500/10"
          >
            <div class="flex items-center justify-between px-6 py-4 border-b border-emerald-100">
              <h3 class="text-lg font-semibold text-gray-800">
                Редактирование профиля
              </h3>
              <button 
                @click="close"
                class="w-9 h-9 flex items-center justify-center rounded-xl hover:bg-emerald-50 transition cursor-pointer"
              >
                <Icon name="ph:x" size="20" class="text-gray-500" />
              </button>
            </div>

            <div class="p-6 space-y-5 min-h-100 max-h-[70vh] overflow-y-auto" v-if="profile">
              <div class="flex items-center gap-4 p-4 rounded-xl bg-emerald-50/50 border border-emerald-100">
                <div class="w-12 h-12 rounded-full bg-linear-to-r from-emerald-400 to-green-600 flex items-center justify-center text-white font-semibold text-lg shrink-0">
                    <img v-if="profile.avatar" :src="formatAvatar(profile.login, profile.avatar)" alt="Avatar" class="w-full h-full object-cover rounded-full" />
                    <span v-else>{{ profile.login.charAt(0).toUpperCase() }}</span>
                </div>
                <div class="min-w-0">
                  <p class="font-medium text-gray-800 truncate">{{ profile.login }}</p>
                  <p class="text-sm text-gray-500 truncate">{{ profile.email }}</p>
                </div>
              </div>

              <div v-if="errorMessage" class="text-sm text-red-500 bg-red-50 border border-red-200 px-4 py-3 rounded-xl">
                {{ errorMessage }}
              </div>

              <div class="space-y-2">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Роль</label>
                <Select
                  v-model="selectedRole"
                  :options="roleOptions"
                  placeholder="Выберите роль"
                  icon="ph:identification-card"
                />
              </div>

              <div class="space-y-2" v-if="selectedRole?.id !== FramebyAppRole.USER">
                <label class="text-sm font-medium text-gray-700 ml-1.5">ФИО</label>
                <div class="relative">
                  <Icon name="ph:user" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
                  <input
                    v-model="fullName"
                    type="text"
                    placeholder="Иванов Иван Иванович"
                    class="w-full pl-9 pr-4 py-2.5 rounded-xl
                           bg-white/60 border border-emerald-100
                           focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                           text-sm cursor-text"
                  />
                </div>
              </div>

              <div class="space-y-2" v-if="selectedRole?.id === FramebyAppRole.BRSM || selectedRole?.id === FramebyAppRole.UNIVERSITY || selectedRole?.id === FramebyAppRole.CUSTOMER">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Суброль</label>
                <div class="relative">
                  <Icon name="ph:address-book" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
                  <input
                    v-model="subrole"
                    type="text"
                    placeholder="Например: координатор, декан..."
                    class="w-full pl-9 pr-4 py-2.5 rounded-xl
                           bg-white/60 border border-emerald-100
                           focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                           text-sm cursor-text"
                  />
                </div>
              </div>
              
              <div v-if="selectedRole?.id === FramebyAppRole.UNIVERSITY" class="space-y-2">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Университет / Кафедра</label>
                <Select
                  v-model="selectedUniversity"
                  :options="departments.map(dp => ({ id: dp.id, name: `${dp.university_name} / ${dp.department_name}`}))"
                  placeholder="Выберите университет и кафедру"
                  icon="ph:buildings"
                />
              </div>
              
              <div v-if="selectedRole?.id === FramebyAppRole.STUDENT" class="space-y-2">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Университет / Кафедра</label>
                <Select
                  v-model="selectedUniversity"
                  :options="departments.map(dp => ({ id: dp.id, name: `${dp.university_name} / ${dp.department_name}`}))"
                  placeholder="Выберите университет и кафедру"
                  icon="ph:graduation-cap"
                />
              </div>

              <div v-if="selectedRole?.id === FramebyAppRole.STUDENT" class="space-y-2">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Специальность</label>
                <Select
                  v-model="selectedSpecialty"
                  :options="specialtyOptions"
                  placeholder="Выберите специальность"
                  icon="ph:book-open"
                  :disabled="!selectedUniversity"
                />
              </div>

              <div v-if="selectedRole?.id === 'student'" class="space-y-2">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Оценка (1-10)</label>
                <div class="relative">
                  <Icon name="ph:star" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
                  <input
                    v-model="grade"
                    type="number"
                    min="1"
                    max="10"
                    step="0.1"
                    placeholder="9.5"
                    class="w-full pl-9 pr-4 py-2.5 rounded-xl
                           bg-white/60 border border-emerald-100
                           focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                           text-sm cursor-text"
                  />
                </div>
              </div>

              

              <div class="space-y-2" v-if="selectedRole?.id !== FramebyAppRole.USER">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Должность</label>
                <div class="relative">
                  <Icon name="ph:briefcase" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
                  <input
                    v-model="position"
                    type="text"
                    placeholder="Должность"
                    class="w-full pl-9 pr-4 py-2.5 rounded-xl
                           bg-white/60 border border-emerald-100
                           focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                           text-sm cursor-text"
                  />
                </div>
              </div>

              <div class="space-y-2" v-if="selectedRole?.id !== FramebyAppRole.USER">
                <label class="text-sm font-medium text-gray-700 ml-1.5">Телефон</label>
                <div class="relative">
                  <Icon name="ph:phone" size="18" class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" />
                  <input
                    v-model="phone"
                    type="tel"
                    placeholder="+375 XX XXX-XX-XX"
                    class="w-full pl-9 pr-4 py-2.5 rounded-xl
                           bg-white/60 border border-emerald-100
                           focus:outline-none focus:ring-2 focus:ring-emerald-400/40
                           text-sm cursor-text"
                  />
                </div>
              </div>
            </div>

            <div class="flex items-center justify-end gap-3 px-6 py-4 border-t border-emerald-100">
              <button
                @click="close"
                class="px-5 py-2.5 rounded-xl text-sm font-medium
                       bg-white border border-emerald-200
                       text-gray-700 hover:bg-emerald-50 transition cursor-pointer"
              >
                Отмена
              </button>
              <button
                @click="handleSave"
                :disabled="isLoading || !selectedRole"
                class="px-5 py-2.5 rounded-xl text-sm font-semibold
                       bg-linear-to-r from-emerald-400 to-green-600 text-white
                       hover:opacity-90 transition cursor-pointer
                       disabled:opacity-50 disabled:cursor-not-allowed
                       shadow-lg shadow-emerald-500/20"
              >
                <span v-if="isLoading">Сохранение...</span>
                <span v-else>Сохранить</span>
              </button>
            </div>
          </div>
        </Transition>
      </div>
    </Transition>
  </Teleport>
</template>
