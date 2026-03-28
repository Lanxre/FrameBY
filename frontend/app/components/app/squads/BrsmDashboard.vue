<script setup lang="ts">
import TabsLayout from "@/components/ui/TabsLayout/TabsLayout.vue";
import StudentSquadCreate from "./StudentSquadCreate.vue";
import StudentSquadTable from "./StudentSquadTable.vue";

const tabs = [
	{
		label: "Создать заявку",
		value: "create",
		icon: "ph:plus-circle",
		component: StudentSquadCreate,
		hasPermission: true,
	},
	{
		label: "Список наборов отрядов",
		value: "list",
		icon: "ph:users-three",
		component: StudentSquadTable,
		hasPermission: true,
	},
];
</script>

<template>
  <TabsLayout :tabs="tabs" default-tab="create" />
</template>
