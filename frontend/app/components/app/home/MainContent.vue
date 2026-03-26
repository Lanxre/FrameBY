<script setup lang="ts">
import MainCard from "@/components/utils/card/MainCard.vue";
import MainBlock from "@/components/utils/blocks/MainBlock.vue";
import AboutCard from "@/components/common/AboutCard.vue";
</script>

<template>
    <div
        class="min-h-screen flex flex-col items-center mx-4 my-10 md:m-10 gap-8"
    >
        <MainCard class="w-full max-w-7xl" />
        <MainBlock class="w-full max-w-7xl" />
        <AboutCard class="w-full max-w-7xl" />
    </div>
</template>
