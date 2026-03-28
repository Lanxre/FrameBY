<script setup lang="ts">
import { useDashboardProfiles } from "@/composables/api/dashboard/useDashboardProfiles";
import type { ProfileRow } from "@/types/dashboard/profile";
import ProfileFilter from "./ProfileFilter.vue";
import ProfileTable from "./ProfileTable.vue";
import ProfileEditModal from "./ProfileEditModal.vue";
import ModalConfirm from "@/components/common/ModalConfirm.vue";
import EnterpriseCreate from "./EnterpriseCreate.vue";
import UniversityCreate from "./UniversityCreate.vue";

const {
	profiles,
	total,
	isLoading,
	errorMessage,
	fetchProfiles,
	deleteProfile,
} = useDashboardProfiles();

const PAGE_SIZE = 20;
const CURRENT_PAGE = 1;
const INITIAL_SELECTION = "all";

const search = ref("");

const selectedType = ref<string>(INITIAL_SELECTION);
const currentPage = ref<number>(CURRENT_PAGE);
const perPage = ref<number>(PAGE_SIZE);

const totalPages = computed(() => Math.ceil(total.value / perPage.value) || 1);

const loadProfiles = () => {
	const offset = (currentPage.value - 1) * perPage.value;
	fetchProfiles({
		type: selectedType.value,
		search: search.value,
		limit: perPage.value,
		offset,
	});
};

watch([selectedType, currentPage], () => {
	loadProfiles();
});

watch(search, () => {
	currentPage.value = 1;
	loadProfiles();
});

onMounted(() => {
	loadProfiles();
});

const isEditModalOpen = ref(false);
const editingProfile = ref<ProfileRow | null>(null);

const openEditModal = (profile: ProfileRow) => {
	editingProfile.value = profile;
	isEditModalOpen.value = true;
};

const closeEditModal = () => {
	isEditModalOpen.value = false;
	editingProfile.value = null;
};

const handlePageChange = (page: number) => {
	currentPage.value = page;
};

const isDeleteModalOpen = ref(false);
const deletingProfile = ref<ProfileRow | null>(null);
const isDeleting = ref(false);

const openDeleteModal = (profile: ProfileRow) => {
	deletingProfile.value = profile;
	isDeleteModalOpen.value = true;
};

const closeDeleteModal = () => {
	isDeleteModalOpen.value = false;
	deletingProfile.value = null;
};

const handleConfirmDelete = async () => {
	if (!deletingProfile.value) return;

	isDeleting.value = true;
	const success = await deleteProfile(deletingProfile.value.user_id);
	isDeleting.value = false;

	if (success) {
		closeDeleteModal();
		loadProfiles();
	}
};
</script>

<template>
  <div class="space-y-4">
    <ProfileFilter 
      v-model:search="search"
      v-model:selectedType="selectedType"
    />

    <div v-if="errorMessage" class="text-sm text-red-500 bg-red-50 border border-red-200 px-4 py-3 rounded-xl">
      {{ errorMessage }}
    </div>

    <ProfileTable
      :profiles="profiles"
      :isLoading="isLoading"
      :total="total"
      :currentPage="currentPage"
      :totalPages="totalPages"
      @edit="openEditModal"
      @page-change="handlePageChange"
      @delete="openDeleteModal"
    />

    <ProfileEditModal
      v-model="isEditModalOpen"
      :profile="editingProfile"
      @update:modelValue="closeEditModal"
      @saved="loadProfiles"
    />

    <ModalConfirm
      v-model="isDeleteModalOpen"
      title="Удалить профиль"
      :description="`Вы уверены, что хотите удалить профиль пользователя ${deletingProfile?.login}? Это действие необратимо.`"
      confirmText="Удалить"
      cancelText="Отмена"
      :loading="isDeleting"
      @confirm="handleConfirmDelete"
    />
    
    <div class="mb-8 mt-10 cursor-default">
      <h1 class="text-2xl sm:text-3xl font-bold text-white text-shadow-lg/30 flex items-center gap-3">
        <Icon name="ph:squares-four" size="28" class="text-emerald-500" />
        Добавление данных
      </h1>
      <p class="mt-2 text-gray-600">
          Добавит данные в систему
      </p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <EnterpriseCreate @created="loadProfiles" />
      <UniversityCreate @created="loadProfiles" />
    </div>
  </div>
</template>
