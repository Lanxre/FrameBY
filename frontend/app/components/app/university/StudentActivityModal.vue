<script setup lang="ts">
import type { StudentEmploymentInfo } from "@/types/frontend/university";
import ModalWindow from "~/components/common/ModalWindow.vue";
import { formatAvatar } from "@/utils/str";

const props = defineProps<{
	student: StudentEmploymentInfo | null;
	isOpen: boolean;
}>();

const emit = defineEmits<{
	close: [];
}>();

const formatDate = (dateStr: string) => {
	const date = new Date(dateStr);
	return date.toLocaleDateString("ru-RU", {
		day: "numeric",
		month: "long",
		year: "numeric",
	});
};

const getStatusColor = (status: string) => {
	switch (status.toLowerCase()) {
		case "active":
		case "accepted":
		case "approved":
			return "bg-green-100 text-green-700 border-green-200";
		case "pending":
		case "waiting":
			return "bg-yellow-100 text-yellow-700 border-yellow-200";
		case "rejected":
		case "cancelled":
			return "bg-red-100 text-red-700 border-red-200";
		default:
			return "bg-gray-100 text-gray-700 border-gray-200";
	}
};

const formatStatus = (status: string) => {
	switch (status.toLowerCase()) {
		case "active":
			return "Активен";
		case "accepted":
			return "Принят";
		case "approved":
			return "Утвержден";
		case "pending":
			return "В ожидании";
		case "waiting":
			return "Ожидает";
		case "rejected":
			return "Отклонен";
		case "cancelled":
			return "Отменен";
		default:
			return "Неизвестный статус";
	}
};

</script>

<template>
	<ModalWindow :model-value="isOpen" title="Деятельность студента" @update:model-value="emit('close')" @close="emit('close')">
		<div v-if="student" class="space-y-6 max-h-[70vh] overflow-y-auto pr-2 custom-scrollbar">
			<div class="flex items-start gap-4 p-4 bg-linear-to-r from-emerald-50 to-green-50 rounded-xl border border-emerald-100">
				<div class="w-16 h-16 rounded-full bg-linear-to-br from-emerald-400 to-green-600 flex items-center justify-center overflow-hidden shadow-md">
					<img
						v-if="student.student.avatar"
						:src="formatAvatar(student.student.login, student.student.avatar)"
						alt="Avatar"
						class="w-full h-full object-cover"
					/>
					<span v-else class="text-white font-bold text-2xl">
						{{ student.student.login.charAt(0).toUpperCase() }}
					</span>
				</div>
				<div class="flex-1">
					<h3 class="text-xl font-bold text-gray-800">
						{{ student.student.profile.full_name }}
					</h3>
					<span class="text-gray-600 font-bold">{{ student.student.login }} @{{ student.student.email }}</span>
					<div class="flex items-center gap-2 text-gray-600 mt-1">
						<Icon name="ph:graduation-cap" size="16" class="text-emerald-500" />
						<span>{{ student.student.profile.specialty || 'Специальность не указана' }}</span>
					</div>
					<div v-if="student.student.profile.phone" class="flex items-center gap-2 text-emerald-600 font-medium mt-1">
						<Icon name="ph:phone" size="16" />
						<span class="text-gray-600">{{ student.student.profile.phone }}</span>
					</div>
					<div v-if="student.student.profile.grade" class="flex items-center gap-2 text-emerald-600 font-medium mt-1">
						<Icon name="ph:book-open" size="16" />
						<span>{{ student.student.profile.grade }} средний балл</span>
					</div>
					
				</div>
			</div>

			<div class="grid grid-cols-2 gap-3">
				<div class="p-3 bg-blue-50 rounded-xl border border-blue-100">
					<div class="flex items-center gap-2 text-blue-700 font-semibold mb-1">
						<Icon name="ph:flag-banner" size="20" />
						<span>Отряды</span>
					</div>
					<p class="text-2xl font-bold text-blue-800">{{ student.student_squads.length }}</p>
				</div>
				<div class="p-3 bg-green-50 rounded-xl border border-green-100">
					<div class="flex items-center gap-2 text-green-700 font-semibold mb-1">
						<Icon name="ph:briefcase" size="20" />
						<span>Работа</span>
					</div>
					<p class="text-2xl font-bold text-green-800">{{ student.works.length }}</p>
				</div>
			</div>

			<div v-if="student.student_squads.length > 0">
				<h4 class="text-lg font-semibold text-gray-800 mb-3 flex items-center gap-2">
					<Icon name="ph:flag-banner-duotone" class="text-blue-500" />
					Студенческие отряды
				</h4>
				<div class="space-y-3">
					<div
						v-for="squad in student.student_squads"
						:key="squad.id"
						class="p-4 bg-white rounded-xl border border-blue-100 shadow-md"
					>
						<div class="flex justify-between items-start gap-3">
							<div class="flex-1 min-w-0">
								<p class="font-semibold text-gray-800">{{ squad.name }}</p>
								<p v-if="squad.description" class="text-sm text-gray-500 mt-1">{{ squad.description }}</p>
							</div>
							<span
								class="px-2.5 py-1 text-xs font-medium rounded-full border shrink-0 bg-green-100 text-green-700 border-green-200"
							>
								Участник
							</span>
						</div>
						<div class="flex items-center gap-1 text-xs text-gray-400 mt-2">
							<Icon name="ph:calendar" size="12" />
							<span>Участник с {{ formatDate(squad.joined_at) }}</span>
						</div>
					</div>
				</div>
			</div>

			<div v-if="student.works.length > 0">
				<h4 class="text-lg font-semibold text-gray-800 mb-3 flex items-center gap-2">
					<Icon name="ph:briefcase-duotone" class="text-green-500" />
					Трудоустройство
				</h4>
				<div class="space-y-3">
					<div
						v-for="work in student.works"
						:key="work.id"
						class="p-4 bg-white rounded-xl border border-green-100 shadow-md"
					>
						<div class="flex justify-between items-start gap-3">
							<div class="flex-1 min-w-0">
								<p class="font-semibold text-gray-800">{{ work.title }}</p>
								<div class="flex items-center gap-1 text-sm text-gray-600 mt-1">
									<Icon name="ph:buildings" size="14" class="text-green-500" />
									<span>{{ work.company }}</span>
								</div>
								<div v-if="work.position" class="flex items-center gap-1 text-sm text-gray-500 mt-0.5">
									<Icon name="ph:user" size="14" />
									<span>{{ work.position }}</span>
								</div>
							</div>
							<span
								class="px-2.5 py-1 text-xs font-medium rounded-full border shrink-0"
								:class="getStatusColor(work.status)"
							>
								{{ formatStatus(work.status) }}
							</span>
						</div>
						<div class="flex items-center gap-1 text-xs text-gray-400 mt-2">
							<Icon name="ph:calendar" size="12" />
							<span>Работает с {{ formatDate(work.started_at) }}</span>
						</div>
					</div>
				</div>
			</div>

			<div v-if="student.student_squads.length === 0 && student.works.length === 0" class="text-center py-8">
				<div class="w-16 h-16 mx-auto mb-3 rounded-full bg-gray-100 flex items-center justify-center">
					<Icon name="ph:user-circle-gear" size="32" class="text-gray-400" />
				</div>
				<p class="text-gray-500 font-medium">Нет данных о деятельности</p>
				<p class="text-sm text-gray-400 mt-1">Студент не состоит в отрядах и не трудоустроен</p>
			</div>
		</div>
	</ModalWindow>
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
