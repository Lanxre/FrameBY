<script setup lang="ts">
import CardLogo from "@/assets/images/main_card_bg.png";
import Button from "@/components/ui/Button.vue";

const handleButtonClick = () => {
	navigateTo("/dashboard");
};
</script>

<template>
  <div 
    class="min-h-125 w-full max-w-7xl mx-auto flex flex-col px-8 py-12 m-10 gap-8 
           shadow-2xl rounded-3xl overflow-hidden relative
           bg-cover bg-center text-white"
    :style="{ backgroundImage: `url(${CardLogo})` }"
  >
    <div class="absolute inset-0 bg-linear-to-r from-emerald-600/70 via-emerald-700/60 to-emerald-800/40" />

    <div class="flex flex-col relative z-10 flex-1">
      <h1 class="text-5xl md:text-6xl wrap-break-word w-190 font-bold uppercase leading-tight text-shadow-lg">
        Автоматизированная система
      </h1>
      <p class="text-2xl font-bold uppercase mt-4 text-shadow-lg">
        На заказ кадров и их подбора
      </p>
      <p class="text-xl mt-10 font-medium text-shadow-lg">
        Быстрая и эффективная система
      </p>
    </div>

    <div class="relative z-10 mt-auto">
      <Button 
        variant="border-shadow" 
        size="lg"
        class="w-full sm:w-auto px-10"
        @click="handleButtonClick"
      >
        Начать использовать
      </Button>
    </div>
  </div>
</template>