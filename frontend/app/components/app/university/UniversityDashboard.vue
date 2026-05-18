<script setup lang="ts">
import TabsLayout from "@/components/ui/TabsLayout/TabsLayout.vue";
import UniversitySquadApproval from "./UniversitySquadApproval.vue";
import UniversityEmploymentRequests from "./UniversityEmploymentRequests.vue";
import UniversityStats from "./UniversityStats.vue";
import UniversityStudentInfo from "./UniversityStudentInfo.vue";
import StudentSquadTable from "@/components/app/squads/StudentSquadTable.vue";

const tabs = [
	{
		label: "Заявки отрядов",
		value: "squads",
		icon: "ph:users-three",
		component: UniversitySquadApproval,
		hasPermission: true,
	},
	{
		label: "Заявки на работу",
		value: "employment",
		icon: "ph:briefcase",
		component: UniversityEmploymentRequests,
		hasPermission: true,
	},
	{
		label: "Списки отрядов",
		value: "squad_list",
		icon: "ph:list-bullets",
		component: StudentSquadTable,
		hasPermission: true,
	},
	{
		label: "Статистика",
		value: "stats",
		icon: "ph:chart-bar",
		component: UniversityStats,
		hasPermission: true,
	},
	{
		label: "Информация о занятости студентов",
		value: "profile",
		icon: "ph:user",
		component: UniversityStudentInfo,
		hasPermission: true,
	}
];
</script>

<template>
  <TabsLayout :tabs="tabs" default-tab="squads" />
</template>
