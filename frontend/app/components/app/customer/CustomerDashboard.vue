<script setup lang="ts">
import TabsLayout from "@/components/ui/TabsLayout/TabsLayout.vue";
import StudentSquadCreate from "@/components/app/squads/StudentSquadCreate.vue";
import CustomerEmploymentCreate from "./CustomerEmploymentCreate.vue";
import CustomerMyRequests from "./CustomerMyRequests.vue";
import CustomerJobResponses from "./CustomerJobResponses.vue";

const tabs = [
	{
		label: "Заявка на создание отряда",
		value: "squad",
		icon: "ph:users-three",
		component: StudentSquadCreate,
		hasPermission: true,
	},
	{
		label: "Заявка на работу",
		value: "employment",
		icon: "ph:briefcase",
		component: CustomerEmploymentCreate,
		hasPermission: true,
	},
	{
		label: "Мои заявки",
		value: "list",
		icon: "ph:list-bullets",
		component: CustomerMyRequests,
		hasPermission: true,
	},
	{
		label: "Отклики",
		value: "responses",
		icon: "ph:paper-plane-tilt",
		component: CustomerJobResponses,
		hasPermission: true,
	},
];
</script>

<template>
	<TabsLayout :tabs="tabs" default-tab="squad" />
</template>
