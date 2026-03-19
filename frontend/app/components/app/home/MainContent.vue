<script setup lang="js">
import MainCard from '~/components/utils/card/MainCard.vue';
import MainBlock from '~/components/utils/blocks/MainBlock.vue';
import AboutCard from '~/components/common/AboutCard.vue';
</script>

<template>
  <div class="min-h-100 flex flex-col items-center m-10 gap-6">
    <MainCard/>
    <MainBlock/>    
    <AboutCard class="w-300"/>
  </div>
</template>