<script setup lang="ts">
import type { FooterSection } from "@/types/frontend/footer";
import FooterLinkCard from "./FooterLinkCard.vue";

defineProps<{
	section: FooterSection;
}>();
</script>

<template>
  <div class="space-y-3 px-2 py-2 shadow-xl shadow-gray-800 rounded-lg bg-linear-to-r from-emerald-300 from-10% bg-emerald-400 via-20% to-emerald-500 to-90%">
      <h3 class="text-sm font-bold uppercase text-emerald-900/70">
        {{ section.title }}
      </h3>

      <ul class="space-y-2">
        <FooterLinkCard v-for="link in section.links" :key="link.label" :link="link" />
      </ul>
  </div>
</template>
