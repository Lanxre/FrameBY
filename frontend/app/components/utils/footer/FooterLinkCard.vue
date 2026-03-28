<script setup lang="ts">
import type { FooterLink } from "@/types/frontend/footer";

defineProps<{
	link: FooterLink;
}>();
</script>

<template>
    <li :key="link.label" class="font-semibold">
    <NuxtLink
        :to="link.to"
        class="text-sm transition"
    >
        <span class="text-teal-800 hover:text-emerald-600/70">
            {{ link.label }}
        </span>
    </NuxtLink>
    </li>
</template>
