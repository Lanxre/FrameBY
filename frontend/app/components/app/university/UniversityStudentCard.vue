<script setup lang="ts">
import type { StudentEmploymentInfo } from "@/types/frontend/university";
import { formatAvatar } from "@/utils/str";

const props = defineProps<{
    student: StudentEmploymentInfo;
}>();

const emit = defineEmits<{
    click: [student: StudentEmploymentInfo];
}>();

const hasSquads = computed(() => props.student.student_squads.length > 0);
const hasWorks = computed(() => props.student.works.length > 0);
const isEmployed = computed(() => hasWorks.value);
const inSquad = computed(() => hasSquads.value);
</script>

<template>
    <div
        @click="emit('click', student)"
        class="group p-4 border-b border-emerald-50 cursor-pointer hover:bg-emerald-50/50 hover:shadow-sm hover:border-emerald-100 transition-all duration-200"
    >
        <div class="flex items-center gap-4">
            <div class="relative">
                <div
                    class="w-12 h-12 rounded-full bg-gradient-to-br from-emerald-400 to-green-600 flex items-center justify-center overflow-hidden shadow-sm"
                >
                    <img
                        v-if="student.student.avatar"
                        :src="
                            formatAvatar(
                                student.student.login,
                                student.student.avatar,
                            )
                        "
                        alt="Avatar"
                        class="w-full h-full object-cover"
                    />
                    <span v-else class="text-white font-semibold text-lg">
                        {{ student.student.login.charAt(0).toUpperCase() }}
                    </span>
                </div>
                <div
                    v-if="isEmployed || inSquad"
                    class="absolute -bottom-1 -right-1 w-4 h-4 rounded-full border-2 border-white flex items-center justify-center"
                    :class="isEmployed ? 'bg-green-500' : 'bg-blue-500'"
                >
                    <Icon
                        v-if="isEmployed"
                        name="ph:briefcase"
                        size="10"
                        class="text-white"
                    />
                    <Icon
                        v-else
                        name="ph:flag-banner"
                        size="10"
                        class="text-white"
                    />
                </div>
            </div>

            <div class="flex-1 min-w-0">
                <p
                    class="font-semibold text-gray-800 truncate group-hover:text-emerald-700 transition-colors"
                >
                    {{
                        student.student.profile.full_name ||
                        student.student.login
                    }}
                </p>
                <div
                    class="flex items-center gap-2 text-sm text-gray-500 truncate"
                >
                    <Icon
                        name="ph:graduation-cap"
                        size="14"
                        class="text-emerald-500"
                    />
                    <span>{{
                        student.student.profile.specialty || "Без специальности"
                    }}</span>
                    <span
                        v-if="student.student.profile.grade"
                        class="text-emerald-600 font-medium"
                        >• {{ student.student.profile.grade }} средний балл</span
                    >
                </div>
            </div>

            <div class="flex flex-col items-end gap-1">
                <div class="flex gap-1.5">
                    <span
                        v-if="hasSquads"
                        class="flex items-center gap-1 px-2.5 py-1 text-xs font-medium bg-blue-50 text-blue-700 rounded-full border border-blue-100"
                    >
                        <Icon name="ph:flag-banner" size="12" />
                        {{ student.student_squads.length }}
                    </span>
                    <span
                        v-if="hasWorks"
                        class="flex items-center gap-1 px-2.5 py-1 text-xs font-medium bg-green-50 text-green-700 rounded-full border border-green-100"
                    >
                        <Icon name="ph:briefcase" size="12" />
                        {{ student.works.length }}
                    </span>
                </div>
                <span
                    v-if="!hasSquads && !hasWorks"
                    class="text-xs text-gray-400"
                >
                    Не трудоустроен
                </span>
            </div>
        </div>

        <div v-if="hasSquads || hasWorks" class="mt-3 flex flex-wrap gap-2">
            <div
                v-for="squad in student.student_squads.slice(0, 2)"
                :key="squad.id"
                class="flex items-center gap-1 px-2 py-1 bg-blue-50/50 rounded-lg text-xs text-blue-700 border border-blue-100/50"
            >
                <Icon name="ph:flag-banner-duotone" size="12" />
                <span class="truncate max-w-[120px]">{{ squad.name }}</span>
            </div>
            <div
                v-for="work in student.works.slice(0, 2)"
                :key="work.id"
                class="flex items-center gap-1 px-2 py-1 bg-green-50/50 rounded-lg text-xs text-green-700 border border-green-100/50"
            >
                <Icon name="ph:briefcase-duotone" size="12" />
                <span class="truncate max-w-[120px]">{{ work.company }}</span>
            </div>
            <span
                v-if="
                    student.student_squads.length > 2 ||
                    student.works.length > 2
                "
                class="px-2 py-1 text-xs text-gray-500"
            >
                + ещё
                {{ student.student_squads.length + student.works.length - 4 }}
            </span>
        </div>
    </div>
</template>
