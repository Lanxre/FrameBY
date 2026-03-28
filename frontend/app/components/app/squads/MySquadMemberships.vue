<script setup lang="ts">
import { useVirtualList } from "@vueuse/core";
import StudentSquadCard from "./StudentSquadCard.vue";
import { useMySquads } from "@/composables/api/squads/useStudentSquads";
import { useLeaveSquad } from "@/composables/api/squads/useJoinLeaveSquad";
import { useNotificationStore } from "#imports";

const { squads, isLoading, errorMessage, fetchMySquads } = useMySquads();
const {
	leave,
	isLoading: isLeaving,
	errorMessage: leaveError,
	reset: resetLeave,
	isSuccess: leaveSuccess,
} = useLeaveSquad();

const { notify } = useNotificationStore();

const loadSquads = () => {
	fetchMySquads();
};

const handleLeave = async (id: string) => {
	const success = await leave(id);
	if (success) {
		resetLeave();
		squads.value = squads.value.filter((squad) => squad.id !== id);
		notify({
			title: "Успех",
			content: "Вы успешно покинули отряд!",
			type: "success",
		});
	} else {
		notify({
			title: "Ошибка",
			content: "Не удалось покинуть отряд!",
			type: "error",
		});
	}
};

const { list, containerProps, wrapperProps } = useVirtualList(squads, {
	itemHeight: 200,
	overscan: 8,
});

onMounted(loadSquads);
</script>

<template>
  <div class="space-y-4">
    <div v-if="errorMessage" class="p-4 bg-red-50 border border-red-200 rounded-lg text-red-600 text-sm">
      {{ errorMessage }}
    </div>

    <div v-if="leaveError" class="p-4 bg-red-50 border border-red-200 rounded-lg text-red-600 text-sm">
      {{ leaveError }}
    </div>

    <div v-if="leaveSuccess" class="p-4 bg-green-50 border border-green-200 rounded-lg text-green-600 text-sm">
      Вы успешно покинули отряд!
    </div>

    <div v-if="isLoading" class="flex justify-center py-8">
      <Icon name="ph:circle-notch" size="32" class="animate-spin text-emerald-500" />
    </div>

    <div v-else-if="squads.length === 0" class="text-center py-8 text-gray-500">
      <Icon name="ph:users-three" size="48" class="mx-auto mb-2 text-gray-300" />
      <p>Вы пока не состоите в отрядах</p>
    </div>

    <div
      v-else
      v-bind="containerProps"
      class="h-125 overflow-y-auto rounded-2xl border border-emerald-100 bg-white/60 backdrop-blur"
    >
      <div v-bind="wrapperProps" class="p-4 space-y-3">
        <StudentSquadCard
          v-for="{ data: squad } in list"
          :key="squad.id"
          :squad="squad"
          :isLeaving="isLeaving"
          :onLeave="handleLeave"
        />
      </div>
    </div>
  </div>
</template>
