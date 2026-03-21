<script setup lang="ts">
import Dropdown from "~/components/ui/Dropdown.vue";
import {
    PROJECT_NAME,
    REQUIRE_MENU,
    REQUIRE_SECTION_MENU,
    PROFILE_MENU,
} from "~/const/index";
</script>

<template>
    <header class="glass sticky top-0 z-50">
        <div
            class="max-w-7xl mx-auto px-4 h-16 flex items-center justify-between shadow-lg shadow-emerald-600/50 rounded-lg"
        >
            <NuxtLink
                to="/"
                class="flex items-center gap-2 font-semibold text-lg group"
            >
                <div class="relative">
                    <div
                        class="absolute inset-0 bg-linear-to-r from-green-400 to-emerald-600 blur-lg opacity-30 group-hover:opacity-60 transition"
                    />
                    <Icon
                        name="mdi:triangle"
                        class="w-6 h-6 relative text-green-500"
                    />
                </div>

                <span
                    class="bg-linear-to-r text-4xl from-green-400 to-emerald-600 bg-clip-text text-transparent"
                >
                    {{ PROJECT_NAME }}
                </span>
            </NuxtLink>

            <nav class="hidden md:flex items-center gap-1">
                <Dropdown :items="REQUIRE_MENU">
                    <template #trigger>
                        <div
                            class="flex items-center rounded-xl text-sm font-medium text-gray-700 dark:text-gray-200 p-2 hover:bg-black/13 dark:hover:bg-white/10 cursor-pointer transition"
                        >
                            <span class="hover:text-emerald-600/70 uppercase"
                                >Предоставление сведений</span
                            >
                            <Icon
                                name="mdi:chevron-down"
                                class="w-4 h-4 opacity-60"
                            />
                        </div>
                    </template>
                </Dropdown>

                <Dropdown :items="REQUIRE_SECTION_MENU">
                    <template #trigger>
                        <div
                            class="flex items-center rounded-xl text-sm font-medium text-gray-700 dark:text-gray-200 p-2 hover:bg-black/13 dark:hover:bg-white/10 cursor-pointer transition"
                        >
                            <span class="hover:text-emerald-600/70 uppercase"
                                >Нормативно-справочная информация</span
                            >
                            <Icon
                                name="mdi:chevron-down"
                                class="w-4 h-4 opacity-60"
                            />
                        </div>
                    </template>
                </Dropdown>

                <NuxtLink
                    to="/about"
                    class="px-3 py-2 w-17.5 rounded-xl text-sm font-medium text-gray-700 dark:text-gray-200 hover:bg-black/5 dark:hover:bg-white/10 transition uppercase"
                >
                    <span class="hover:text-emerald-600/70">О нас</span>
                </NuxtLink>

                <NuxtLink
                    to="/contact"
                    class="px-3 py-2 rounded-xl text-sm font-medium text-gray-700 dark:text-gray-200 hover:bg-black/5 dark:hover:bg-white/10 transition uppercase"
                >
                    <span class="hover:text-emerald-600/70">Контакты</span>
                </NuxtLink>
            </nav>

            <div class="flex items-center gap-2">
                <NuxtLink
                    to="/auth/register"
                    class="hidden md:flex items-center gap-2 px-4 py-2 rounded-xl text-sm font-medium bg-linear-to-r from-green-400 to-emerald-600 text-white transition shadow-lg shadow-green-500/20"
                >
                    <Icon name="mdi:rocket-launch-outline" class="w-4 h-4" />
                    <span class="hover:text-white transition-all"
                        >Зарегистрироваться</span
                    >
                </NuxtLink>
            </div>
        </div>
    </header>
</template>
