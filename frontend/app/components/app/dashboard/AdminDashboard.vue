<script setup lang="ts">
import { useDashboardProfiles } from '@/composables/api/dashboard/useDashboardProfiles'
import type { ProfileRow } from '@/types/dashboard/profile'
import ProfileFilter from './ProfileFilter.vue'
import ProfileTable from './ProfileTable.vue'
import ProfileEditModal from './ProfileEditModal.vue'

const { profiles, total, isLoading, errorMessage, fetchProfiles } = useDashboardProfiles()

const search = ref('')
const selectedType = ref('all')
const currentPage = ref(1)
const perPage = ref(20)

const totalPages = computed(() => Math.ceil(total.value / perPage.value) || 1)

const loadProfiles = () => {
  const offset = (currentPage.value - 1) * perPage.value
  fetchProfiles({
    type: selectedType.value,
    search: search.value,
    limit: perPage.value,
    offset
  })
}

watch([selectedType, currentPage], () => {
  loadProfiles()
})

watch(search, () => {
  currentPage.value = 1
  loadProfiles()
})

onMounted(() => {
  loadProfiles()
})

const isEditModalOpen = ref(false)
const editingProfile = ref<ProfileRow | null>(null)

const openEditModal = (profile: ProfileRow) => {
  editingProfile.value = profile
  isEditModalOpen.value = true
}

const closeEditModal = () => {
  isEditModalOpen.value = false
  editingProfile.value = null
}

const handlePageChange = (page: number) => {
  currentPage.value = page
}
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
    />

    <ProfileEditModal
      v-model="isEditModalOpen"
      :profile="editingProfile"
      @update:modelValue="closeEditModal"
      @saved="loadProfiles"
    />
  </div>
</template>
