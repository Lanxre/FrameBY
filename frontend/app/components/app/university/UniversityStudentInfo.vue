<script setup lang="ts">
import type { StudentEmploymentInfo } from "@/types/frontend/university";
import { useUniversityStudents } from "~/composables/api/university/useUniversityStudents";

import UniversityStudentCard from "./UniversityStudentCard.vue";
import StudentActivityModal from "./StudentActivityModal.vue";

const { students, isLoading, fetchStudents, loadMore, hasMore, total } = useUniversityStudents();

const selectedStudent = ref<StudentEmploymentInfo | null>(null);
const isModalOpen = ref(false);

const handleScroll = async (e: Event) => {
	const target = e.target as HTMLElement;
	const scrollBottom = target.scrollHeight - target.scrollTop - target.clientHeight;
	if (scrollBottom < 200 && hasMore.value && !isLoading.value) {
		await loadMore();
	}
};

const openStudentModal = (student: StudentEmploymentInfo) => {
	selectedStudent.value = student;
	isModalOpen.value = true;
};

const closeModal = () => {
	isModalOpen.value = false;
	selectedStudent.value = null;
};

onMounted(() => {
	fetchStudents();
});
</script>

<template>
	<div class="space-y-4">
		<div class="flex items-center justify-between cursor-default">
			<div class="flex items-center gap-3">
				<div class="w-10 h-10 rounded-xl bg-linear-to-br from-emerald-400 to-green-600 flex items-center justify-center">
					<Icon name="ph:users-three" class="text-white" size="22" />
				</div>
				<div>
					<h2 class="text-xl font-bold text-gray-800">Студенты</h2>
					<p class="text-sm text-gray-500">Информация о занятости</p>
				</div>
			</div>
			<div class="flex items-center gap-2 px-3 py-1.5 bg-emerald-50 rounded-full border border-emerald-100">
				<Icon name="ph:users" size="16" class="text-emerald-600" />
				<span class="text-sm font-medium text-emerald-700">{{ students.length }} из {{ total }}</span>
			</div>
		</div>

		<div
			class="h-[550px] overflow-y-auto rounded-xl bg-white border border-emerald-100 shadow-sm custom-scrollbar"
			@scroll="handleScroll"
		>
			<div v-if="isLoading && students.length === 0" class="flex flex-col items-center justify-center h-full py-12">
				<div class="animate-spin rounded-full h-10 w-10 border-3 border-emerald-500 border-t-transparent mb-4"></div>
				<p class="text-gray-500">Загрузка студентов...</p>
			</div>

			<div v-else-if="students.length === 0" class="flex flex-col items-center justify-center h-full py-12">
				<div class="w-16 h-16 mb-4 rounded-full bg-emerald-50 flex items-center justify-center">
					<Icon name="ph:users" size="32" class="text-emerald-300" />
				</div>
				<p class="text-gray-500 font-medium">Нет студентов</p>
				<p class="text-sm text-gray-400 mt-1">В данном университете нет студентов</p>
			</div>

			<div v-else class="divide-y divide-emerald-50">
				<UniversityStudentCard
					v-for="item in students"
					:key="item.student.user_id"
					:student="item"
					@click="openStudentModal(item)"
				/>
			</div>

			<div v-if="isLoading && students.length > 0" class="flex justify-center py-4">
				<div class="animate-spin rounded-full h-6 w-6 border-2 border-emerald-500 border-t-transparent"></div>
			</div>
		</div>

		<StudentActivityModal
			:student="selectedStudent"
			:is-open="isModalOpen"
			@close="closeModal"
		/>
	</div>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
	width: 6px;
}

.custom-scrollbar::-webkit-scrollbar-track {
	background: transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
	background: #d1d5db;
	border-radius: 10px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
	background: #10b981;
}

.custom-scrollbar {
	scrollbar-width: thin;
	scrollbar-color: #d1d5db transparent;
}
</style>
