<script setup lang="ts">
import { useVirtualList } from "@vueuse/core";
import { useAuthStore } from "@/stores/auth";
import {
	useStudentSquads,
	useMySquads,
	useSquadStatuses,
} from "@/composables/api/squads/useStudentSquads";
import { useUpdateSquad } from "@/composables/api/squads/useUpdateSquad";
import { SQUAD_STATUS_COLORS, SQUAD_STATUS_LABELS } from "@/const/squad";
import type { StudentSquad } from "@/types/frontend/student-squad";
import { formatDate, formatTotalStudentSquads } from "@/utils/str";
import ModalConfirm from "@/components/common/ModalConfirm.vue";
import ToolTip from "@/components/ui/ToolTip.vue";
import StudentSquadEditModal from "./StudentSquadEditModal.vue";
import { FramebyAppRole } from "~/types/frontend/enums/role";

const emit = defineEmits<{ updated: [] }>();

const { user } = useAuthStore();
const { notify } = useNotificationStore();

const {
	squads,
	total,
	isLoading: isGlobalLoading,
	errorMessage,
	fetchSquads,
} = useStudentSquads();
const {
	squads: mySquads,
	fetchMySquads,
	isLoading: isMyLoading,
} = useMySquads();
const { fetchStatuses } = useSquadStatuses();
const {
	closeRecruitment,
	openRecruitment,
	isLoading: isUpdating,
} = useUpdateSquad();

const showMySquadsOnly = ref(true);
const showCloseModal = ref(false);
const squadToClose = ref<StudentSquad | null>(null);
const showEditModal = ref(false);
const squadToEdit = ref<StudentSquad | null>(null);
const limit = ref(10);
const offset = ref(0);

const isLoading = computed(() => isGlobalLoading.value || isMyLoading.value);
const totalPages = computed(() => Math.ceil(total.value / limit.value));
const currentPage = computed(() => Math.floor(offset.value / limit.value) + 1);
const isMyView = computed(() => showMySquadsOnly.value);

const sourceList = ref<StudentSquad[]>([]);

watch(
	[squads, mySquads, isMyView],
	() => {
		sourceList.value = isMyView.value ? mySquads.value : squads.value;
	},
	{ immediate: true },
);

const loadData = async () => {
	if (isMyView.value) {
		await fetchMySquads();
	} else {
		await fetchSquads({ limit: limit.value, offset: offset.value });
	}
};

watch([isMyView, offset], loadData);

onMounted(async () => {
	await Promise.all([loadData(), fetchStatuses()]);
});

const goToPage = (page: number) => {
	if (page < 1 || page > totalPages.value) return;
	offset.value = (page - 1) * limit.value;
};

const openCloseModal = (squad: StudentSquad) => {
	squadToClose.value = squad;
	showCloseModal.value = true;
};

const openEditModal = (squad: StudentSquad) => {
	squadToEdit.value = squad;
	showEditModal.value = true;
};

const confirmClose = async () => {
	if (!squadToClose.value) return;
	const success = await closeRecruitment(squadToClose.value.id);
	if (success) {
		const list = isMyView.value ? mySquads.value : squads.value;
		const squad = list.find((s) => s.id === squadToClose.value!.id);
		if (squad) squad.status = "closed";
		showCloseModal.value = false;
		squadToClose.value = null;
		notify({
			title: "Отряд успешно закрыт",
			content: "Закрытие отряда успешно выполнено",
			type: "success",
		});
	} else {
		notify({
			title: "Ошибка закрытия отряда",
			content: "Не удалось закрыть отряд",
			type: "error",
		});
	}
};

const openRecruitmentFor = async (squad: StudentSquad) => {
	const success = await openRecruitment(squad.id);
	if (success) {
		const list = isMyView.value ? mySquads.value : squads.value;
		const item = list.find((s) => s.id === squad.id);

		if (user!.role === FramebyAppRole.BRSM && item) {
			item.status = "recruitment_open";
		} else if (item) {
			item.status = "pending";
		}

		notify({
			title: "Успех",
			content: "Отряд успешно открыт для набора",
			type: "success",
		});
	} else {
		notify({
			title: "Ошибка открытия отряда для набора",
			content: "Не удалось открыть отряд для набора",
			type: "error",
		});
	}
};

const handleSquadSaved = (data: {
	id: string;
	title: string;
	description: string | null;
	profile: string | null;
	max_participants: number;
}) => {
	const list = isMyView.value ? mySquads.value : squads.value;
	const squad = list.find((s) => s.id === data.id);

	const updatedTime = new Date(Date.now());
	updatedTime.setHours(updatedTime.getHours() + 3);

	if (squad) {
		squad.title = data.title;
		squad.description = data.description;
		squad.profile = data.profile;
		squad.max_participants = data.max_participants;
		squad.updated_at = updatedTime;
	}
	showEditModal.value = false;
	squadToEdit.value = null;

	notify({
		title: "Отряд успешно обновлен",
		content: "Изменения сохранены",
		type: "success",
	});
};

const { list, containerProps, wrapperProps } = useVirtualList(sourceList, {
	itemHeight: 220,
	overscan: 8,
});
</script>

<template>
  <div class="space-y-5 min-h-100">
    <div class="flex items-center justify-between">
      <div class="flex items-center gap-2">
        <Icon name="ph:list-bullets" size="22" class="text-emerald-500" />
        <h2 class="text-lg font-semibold text-gray-800">Список наборов отрядов</h2>
      </div>

      <div class="flex items-center gap-4">
        <label class="flex items-center gap-2 cursor-pointer">
          <input type="checkbox" v-model="showMySquadsOnly" class="sr-only peer" />
          <div class="w-10 h-6 bg-gray-200 rounded-full peer-checked:bg-emerald-500 transition relative">
            <div class="absolute top-1 left-1 w-4 h-4 bg-white rounded-full transition peer-checked:translate-x-4"></div>
          </div>
          <span class="text-sm text-gray-600">Только мои</span>
        </label>
        <span class="text-sm text-gray-500">Всего: {{ isMyView ? formatTotalStudentSquads(mySquads.length) : formatTotalStudentSquads(total) }}</span>
      </div>
    </div>

    <div v-if="errorMessage" class="text-sm text-red-600 bg-red-50 border border-red-200 px-4 py-3 rounded-xl">
      {{ errorMessage }}
    </div>

    <div
      v-bind="containerProps"
      class="h-150 overflow-y-auto rounded-2xl border border-emerald-100 bg-white/60 backdrop-blur"
    >
      <div v-bind="wrapperProps" class="p-4 space-y-3">
        <div
          v-for="{ data: squad } in list"
          :key="squad.id"
          class="p-5 rounded-2xl border border-emerald-100 bg-white shadow-lg"
        >
          <div class="flex justify-between gap-4">
            <div class="flex-1">
              <div class="flex items-center gap-2 mb-2">
                <h3 class="font-semibold text-gray-800">{{ squad.title }}</h3>
                <span
                  class="px-2 py-0.5 text-xs rounded-full"
                  :class="[SQUAD_STATUS_COLORS[squad.status]?.bg, SQUAD_STATUS_COLORS[squad.status]?.text]"
                >
                  {{ SQUAD_STATUS_LABELS[squad.status] }}
                </span>
              </div>

              <div class="grid grid-cols-1 sm:grid-cols-2 gap-3 text-sm text-gray-600 mb-3">
                <div class="flex items-center gap-1">
                  Необходимо участников: {{ squad.current_count }} / {{ squad.max_participants }}
                  <Icon name="ph:users" size="14" />
                </div>

                <div class="flex flex-col gap-1">
                  <div class="flex items-center gap-1">
                    <Icon name="ph:calendar" size="14" />
                    <span>{{ formatDate(squad.created_at) }}</span>
                  </div>
                  <div class="flex items-center gap-1">
                    <Icon name="ph:clock" size="14" />
                    <span>{{ formatDate(squad.updated_at) }}</span>
                  </div>
                </div>

                <div class="col-span-2 flex flex-col gap-2">
                  <div class="flex gap-1 items-center">
                    <Icon name="ph:user" size="14" />
                    {{ squad.organizer.name || 'Неизвестно' }}
                  </div>
                  <div class="flex gap-1 items-center">
                    <Icon name="ph:briefcase" size="14" />
                    {{ squad.organizer.position || 'Неизвестно' }}
                  </div>
                  <div class="flex gap-1 items-center">
                    <Icon name="ph:phone" size="14" />
                    {{ squad.organizer.phone || 'Неизвестно' }}
                  </div>
                </div>

                <div v-if="squad.profile" class="col-span-2 text-xs text-emerald-600 font-medium">
                  {{ squad.profile }}
                </div>
              </div>

              <p v-if="squad.description" class="text-sm text-gray-500 line-clamp-2">
                {{ squad.description }}
              </p>
            </div>

            <div v-if="isMyView || squad.organizer.id === user?.id" class="flex flex-col gap-2">
              <ToolTip text="Редактировать">
                <button @click="openEditModal(squad)" class="p-1.5 text-gray-400 hover:text-emerald-500 transition rounded-lg cursor-pointer">
                  <Icon name="ph:pencil" size="18" />
                </button>
              </ToolTip>

              <ToolTip v-if="squad.status === 'recruitment_open'" text="Закрыть">
                <button @click="openCloseModal(squad)" class="p-1.5 text-gray-400 hover:text-red-500 transition rounded-lg cursor-pointer">
                  <Icon name="ph:x" size="18" />
                </button>
              </ToolTip>

              <ToolTip v-if="squad.status === 'closed'" text="Открыть">
                <button @click="openRecruitmentFor(squad)" class="p-1.5 text-gray-400 hover:text-green-500 transition rounded-lg cursor-pointer">
                  <Icon name="ph:plus" size="18" />
                </button>
              </ToolTip>
            </div>
          </div>

          <div class="mt-4">
            <div class="w-full h-2 bg-gray-100 rounded-full overflow-hidden">
              <div
                class="h-full bg-emerald-500 transition-all"
                :style="{ width: `${Math.min((squad.current_count / squad.max_participants) * 100, 100)}%` }"
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="!isMyView && totalPages > 1" class="flex justify-center items-center gap-3 pt-4">
      <button
        @click="goToPage(currentPage - 1)"
        :disabled="currentPage === 1 || isLoading"
        class="w-10 h-10 rounded-xl border border-emerald-100 bg-white/70 hover:bg-emerald-50 disabled:opacity-50 flex items-center justify-center"
      >
        <Icon name="ph:caret-left" size="18" />
      </button>

      <span class="text-sm text-gray-600">{{ currentPage }} / {{ totalPages }}</span>

      <button
        @click="goToPage(currentPage + 1)"
        :disabled="currentPage === totalPages || isLoading"
        class="w-10 h-10 rounded-xl border border-emerald-100 bg-white/70 hover:bg-emerald-50 disabled:opacity-50 flex items-center justify-center"
      >
        <Icon name="ph:caret-right" size="18" />
      </button>
    </div>

    <ModalConfirm
      v-model="showCloseModal"
      title="Закрыть набор"
      :description="`Закрыть набор '${squadToClose?.title}'?`"
      confirmText="Закрыть"
      :loading="isUpdating"
      @confirm="confirmClose"
    />

    <StudentSquadEditModal
      v-model="showEditModal"
      :squad="squadToEdit"
      @saved="(data) => handleSquadSaved(data)"
    />
  </div>
</template>