<script setup lang="ts">
import FooterCard from './FooterCard.vue';
import { FOOTER_SECTIONS } from '@/const';
</script>

<template>
    <FooterCard 
        v-for="section in FOOTER_SECTIONS" 
        :key="section.title" 
        :section="section"
    />
</template>
