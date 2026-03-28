<script setup lang="ts">
import { useAuthStore } from "@/stores/auth";

import UserProfileCard from "~/components/app/profile/UserProfileCard.vue";

definePageMeta({
	middleware: ["auth"],
});

const authStore = useAuthStore();
const { user } = storeToRefs(authStore);
</script>
<template>
<div class="max-w-5xl mx-auto px-4 py-10 sm:py-14" v-if="user">
    <UserProfileCard :user="user!" />
</div>
<div v-else>
    <div class="max-w-5xl mx-auto px-4 py-10 sm:py-14">
        <div class="text-center">
            <h1 class="text-2xl font-bold">Профиль</h1>
            <p class="mt-4">Вы не авторизованы.</p>
        </div>
    </div>
</div>
</template>