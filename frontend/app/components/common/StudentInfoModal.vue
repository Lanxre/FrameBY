<script setup lang="ts">
import ModalWindow from "./ModalWindow.vue";
import type { EmploymentParticipant } from "@/types/frontend/employment";
import { $api } from "@/composables/api/useApi";

const props = defineProps<{
	modelValue: boolean;
	participant: EmploymentParticipant | null;
}>();

const emit = defineEmits<{
	(e: "update:modelValue", value: boolean): void;
}>();

const isOpen = computed({
	get: () => props.modelValue,
	set: (value) => emit("update:modelValue", value),
});

interface StudentProfile {
	full_name: string;
	specialty?: string;
	grade?: number;
	position?: string;
	phone?: string;
	university_name?: string;
	department_name?: string;
}

const studentProfile = ref<StudentProfile | null>(null);
const isLoading = ref(false);

watch(
	[() => props.modelValue, () => props.participant],
	async ([open, participant]) => {
		if (open && participant) {
			await fetchStudentProfile(participant.user_id);
		} else if (!open) {
			studentProfile.value = null;
		}
	},
	{ immediate: true }
);

const fetchStudentProfile = async (userId: string) => {
	isLoading.value = true;
	studentProfile.value = null;
	try {
		const response = await $api<StudentProfile>(
			`/profile/student/by-user?user_id=${userId}`,
			{ method: "GET" }
		);
		studentProfile.value = response;
	} catch (e) {
		console.error("Error fetching student profile:", e);
		studentProfile.value = {
			full_name: props.participant?.student_name || "Неизвестно",
		};
	} finally {
		isLoading.value = false;
	}
};
</script>

<template>
	<ModalWindow v-model="isOpen" title="Информация о студенте">
		<div v-if="isLoading" class="flex justify-center py-8">
			<div class="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
		</div>

		<div v-else-if="studentProfile">
			<div class="rounded-2xl border border-emerald-100 bg-gradient-to-br from-emerald-50/50 to-green-50/50 p-6">
				<div class="flex items-start gap-4 mb-6">
					<div class="w-16 h-16 rounded-full bg-gradient-to-br from-emerald-400 to-green-500 flex items-center justify-center shadow-lg">
						<Icon name="ph:user" size="32" class="text-white" />
					</div>
					<div class="flex-1 min-w-0">
						<h3 class="text-xl font-bold text-gray-800 truncate">
							{{ studentProfile.full_name }}
						</h3>
						<div class="flex items-center gap-2 mt-1">
							<span class="px-2.5 py-1 text-xs font-medium rounded-full bg-emerald-100 text-emerald-700">
								Студент
							</span>
							<span v-if="studentProfile.grade" class="px-2.5 py-1 text-xs font-medium rounded-full bg-green-100 text-green-700">
								{{ studentProfile.grade }} средний балл
							</span>
						</div>
					</div>
				</div>

				<div class="space-y-3">
					<div v-if="studentProfile.university_name" class="flex items-center gap-3 p-3 rounded-xl bg-white/70 border border-emerald-100">
						<div class="w-8 h-8 rounded-lg bg-emerald-100 flex items-center justify-center shrink-0">
							<Icon name="ph:graduation-cap" size="16" class="text-emerald-600" />
						</div>
						<div class="min-w-0">
							<p class="text-xs text-gray-500">Университет</p>
							<p class="text-sm font-medium text-gray-800 truncate">{{ studentProfile.university_name }}</p>
						</div>
					</div>

					<div v-if="studentProfile.department_name" class="flex items-center gap-3 p-3 rounded-xl bg-white/70 border border-emerald-100">
						<div class="w-8 h-8 rounded-lg bg-green-100 flex items-center justify-center shrink-0">
							<Icon name="ph:buildings" size="16" class="text-green-600" />
						</div>
						<div class="min-w-0">
							<p class="text-xs text-gray-500">Факультет</p>
							<p class="text-sm font-medium text-gray-800 truncate">{{ studentProfile.department_name }}</p>
						</div>
					</div>

					<div v-if="studentProfile.specialty" class="flex items-center gap-3 p-3 rounded-xl bg-white/70 border border-emerald-100">
						<div class="w-8 h-8 rounded-lg bg-teal-100 flex items-center justify-center shrink-0">
							<Icon name="ph:book" size="16" class="text-teal-600" />
						</div>
						<div class="min-w-0">
							<p class="text-xs text-gray-500">Специальность</p>
							<p class="text-sm font-medium text-gray-800 truncate">{{ studentProfile.specialty }}</p>
						</div>
					</div>

					<div v-if="studentProfile.position" class="flex items-center gap-3 p-3 rounded-xl bg-white/70 border border-emerald-100">
						<div class="w-8 h-8 rounded-lg bg-amber-100 flex items-center justify-center shrink-0">
							<Icon name="ph:briefcase" size="16" class="text-amber-600" />
						</div>
						<div class="min-w-0">
							<p class="text-xs text-gray-500">Должность</p>
							<p class="text-sm font-medium text-gray-800 truncate">{{ studentProfile.position }}</p>
						</div>
					</div>

					<div v-if="studentProfile.phone" class="flex items-center gap-3 p-3 rounded-xl bg-white/70 border border-emerald-100">
						<div class="w-8 h-8 rounded-lg bg-blue-100 flex items-center justify-center shrink-0">
							<Icon name="ph:phone" size="16" class="text-blue-600" />
						</div>
						<div class="min-w-0">
							<p class="text-xs text-gray-500">Телефон</p>
							<p class="text-sm font-medium text-gray-800">{{ studentProfile.phone }}</p>
						</div>
					</div>
				</div>
			</div>
		</div>

		<div v-else class="text-center py-8 text-gray-500">
			<Icon name="ph:user-circle-dashed" size="48" class="mx-auto mb-3 text-gray-300" />
			<p>Информация о студенте недоступна</p>
		</div>
	</ModalWindow>
</template>
