<script setup lang="ts">
import { $api } from "@/composables/api/useApi";

interface UniversityStats {
	university_name: string;
	department_name: string;
	total_students: number;
	students_in_squads: number;
	students_employed: number;
	total_squads: number;
	job_invitations: number;
}

const stats = ref<UniversityStats | null>(null);
const isLoading = ref(true);

const fetchStats = async () => {
	isLoading.value = true;
	try {
		stats.value = await $api<UniversityStats>("/profile/university/stats", {
			method: "GET",
		});
	} catch (e) {
		console.error("Error fetching university stats:", e);
	} finally {
		isLoading.value = false;
	}
};

onMounted(() => {
	fetchStats();
});

const statCards = computed(() => {
	if (!stats.value) return [];
	return [
		{
			value: stats.value.total_students,
			label: "Всего студентов",
			icon: "ph:users-three",
			gradient: "from-emerald-500 to-teal-600",
			iconBg: "from-emerald-400 to-teal-500",
		},
		{
			value: stats.value.students_in_squads,
			label: "В отрядах",
			subtext: `${Math.round((stats.value.students_in_squads / stats.value.total_students) * 100)}% от общего числа`,
			icon: "ph:flag-banner",
			gradient: "from-green-500 to-emerald-600",
			iconBg: "from-green-400 to-emerald-500",
		},
		{
			value: stats.value.students_employed,
			label: "С вторичной занятостью",
			subtext: `${Math.round((stats.value.students_employed / stats.value.total_students) * 100)}% трудоустроено`,
			icon: "ph:briefcase",
			gradient: "from-teal-500 to-cyan-600",
			iconBg: "from-teal-400 to-cyan-500",
		},
		{
			value: stats.value.total_squads,
			label: "Студенческих отрядов",
			icon: "ph:users",
			gradient: "from-lime-500 to-green-600",
			iconBg: "from-lime-400 to-green-500",
		},
		{
			value: stats.value.job_invitations,
			label: "Активных вакансий",
			icon: "ph:envelope-simple",
			gradient: "from-emerald-400 to-cyan-500",
			iconBg: "from-emerald-300 to-cyan-400",
		},
	];
});
</script>

<template>
	<div class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
		<div class="px-6 py-4 border-b border-gray-100 bg-gradient-to-r from-emerald-50 to-teal-50">
			<div class="flex items-center gap-3">
				<div class="w-10 h-10 rounded-xl bg-gradient-to-br from-emerald-400 to-green-600 flex items-center justify-center shadow-lg shadow-emerald-200">
					<Icon name="ph:chart-bar" class="text-white" size="22" />
				</div>
				<div>
					<h2 class="text-lg font-bold text-gray-800">Статистика факультета</h2>
					<p v-if="stats" class="text-sm text-gray-500">
						{{ stats.university_name }} — {{ stats.department_name }}
					</p>
				</div>
			</div>
		</div>

		<div class="p-6">
			<div v-if="isLoading" class="flex justify-center py-12">
				<div class="animate-spin rounded-full h-10 w-10 border-b-2 border-emerald-600"></div>
			</div>

			<div v-else-if="stats" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-5 gap-4">
				<div
					v-for="(card, index) in statCards"
					:key="index"
					class="group relative overflow-hidden rounded-xl p-5 bg-gradient-to-br shadow-md hover:shadow-xl transition-all duration-300 hover:-translate-y-1"
					:class="card.gradient"
				>
					<div class="absolute top-0 right-0 w-24 h-24 bg-white/10 rounded-full -translate-y-8 translate-x-8"></div>
					<div class="absolute bottom-0 left-0 w-16 h-16 bg-white/5 rounded-full translate-y-4 -translate-x-4"></div>
					
					<div class="relative z-10">
						<div class="flex items-center justify-between mb-3">
							<span class="text-3xl font-bold text-white drop-shadow-sm">{{ card.value }}</span>
							<div class="w-10 h-10 rounded-lg bg-white/20 flex items-center justify-center backdrop-blur-sm">
								<Icon :name="card.icon" class="text-white" size="22" />
							</div>
						</div>
						<p class="text-white/90 font-medium text-sm">{{ card.label }}</p>
						<p v-if="card.subtext" class="text-white/70 text-xs mt-1">{{ card.subtext }}</p>
					</div>
				</div>
			</div>

			<div v-else class="text-center text-gray-500 py-12">
				<Icon name="ph:warning-circle" size="48" class="text-gray-300 mb-2 mx-auto" />
				<p>Не удалось загрузить статистику</p>
			</div>
		</div>
	</div>
</template>
