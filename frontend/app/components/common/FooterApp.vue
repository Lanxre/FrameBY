<script setup lang="ts">
import { PROJECT_NAME, GITHUB_URL, TELEGRAM_URL, TWITTER_URL, FOOTER_SECTIONS } from '@/const';
import FooterList from '@/components/utils/footer/FooterList.vue';
</script>

<template>
  <footer class="mt-auto bg-radial-[at_25%_25%] from-white to-emerald-600 to-75% border-t border-emerald-300">
    <div class="max-w-7xl mx-auto px-4 py-12">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-10">
          
        <div class="space-y-4">
          <NuxtLink
            to="/"
            class="flex items-center gap-2 font-semibold text-lg group"
          >
            <div class="relative">
              <div
                class="absolute inset-0 blur-lg opacity-30 group-hover:opacity-60 transition"
                style="background: var(--color-primary-gradient)"
              />
              <Icon
                name="mdi:triangle"
                class="w-6 h-6 relative text-(--color-primary)"
              />
            </div>

            <span
              class="bg-clip-text text-transparent"
              style="background-image: var(--color-primary-gradient)"
            >
                {{ PROJECT_NAME }}
            </span>
          </NuxtLink>

          <p class="text-sm max-w-xs">
            Современное приложение для управления заказами для поготовки кадров и подбора персонала.
          </p>

          <div class="flex items-center gap-3">
            <a :href="GITHUB_URL" class="p-2 rounded-xl hover:bg-white/5 transition">
              <Icon name="mdi:github" class="w-5 h-5" />
            </a>
            <a :href="TWITTER_URL" class="p-2 rounded-xl hover:bg-white/5 transition">
              <Icon name="mdi:twitter" class="w-5 h-5" />
            </a>
            <a :href="TELEGRAM_URL" class="p-2 rounded-xl hover:bg-white/5 transition">
              <Icon name="mdi:telegram" class="w-5 h-5" />
            </a>
          </div>
        </div>

        <FooterList/>

      </div>

      <div class="mt-10 h-px bg-white/10" />

      <div class="mt-6 flex flex-col md:flex-row items-center justify-between gap-4 ">

        <p class="text-sm">
          © {{ new Date().getFullYear() }} FrameBy. Все права защищены.
        </p>

      </div>
    </div>
  </footer>
</template>