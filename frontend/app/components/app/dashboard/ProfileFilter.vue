<script setup lang="ts">
import { PROFILE_TYPES, type ProfileType } from "@/types/dashboard/profile";

const props = defineProps<{
	search: string;
	selectedType: string;
}>();

const emit = defineEmits<{
	"update:search": [value: string];
	"update:selectedType": [value: string];
}>();

const isDropdownOpen = ref(false);
const searchInput = ref(props.search);

const selectedTypeInfo = computed(
	() =>
		PROFILE_TYPES.find((t) => t.value === props.selectedType) ||
		PROFILE_TYPES[0],
);

let searchTimeout: ReturnType<typeof setTimeout>;
watch(searchInput, (value) => {
	clearTimeout(searchTimeout);
	searchTimeout = setTimeout(() => {
		emit("update:search", value);
	}, 300);
});

watch(
	() => props.search,
	(value) => {
		searchInput.value = value;
	},
);

const selectType = (type: ProfileType) => {
	emit("update:selectedType", type.value);
	isDropdownOpen.value = false;
};
</script>

<template>
  <div class="flex items-center gap-4 flex-wrap">
    <div class="relative flex-1 min-w-60">
      <Icon 
        name="ph:magnifying-glass" 
        size="18"
        class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" 
      />
      <input
        v-model="searchInput"
        type="text"
        placeholder="Поиск по имени, почте, логину..."
        class="w-full pl-9 pr-4 py-2.5 rounded-xl
               bg-white/60 border border-emerald-100
               focus:outline-none focus:ring-2 focus:ring-emerald-400/40
               text-sm"
      />
    </div>

    <div class="relative" v-click-outside="() => isDropdownOpen = false">
      <button
        @click="isDropdownOpen = !isDropdownOpen"
        class="flex items-center gap-3 px-4 py-2 rounded-xl
               bg-white/60 border border-emerald-100
               hover:bg-emerald-50 transition
               text-sm font-medium"
        :class="{ 'ring-2 ring-emerald-400/40': isDropdownOpen }"
      >
        <div 
          class="w-7 h-7 rounded-lg flex items-center justify-center" 
          :class="[selectedTypeInfo.bgColor, selectedTypeInfo.color]"
        >
          <Icon :name="selectedTypeInfo.icon" size="16" />
        </div>
        <span class="text-gray-700">{{ selectedTypeInfo.label }}</span>
        <Icon 
          name="ph:caret-down" 
          size="16" 
          class="text-gray-400 transition-transform duration-200"
          :class="{ 'rotate-180': isDropdownOpen }"
        />
      </button>

      <Transition
        enter-active-class="transition duration-100 ease-out"
        enter-from-class="transform scale-95 opacity-0"
        enter-to-class="transform scale-100 opacity-100"
        leave-active-class="transition duration-75 ease-in"
        leave-from-class="transform scale-100 opacity-100"
        leave-to-class="transform scale-95 opacity-0"
      >
        <div 
          v-if="isDropdownOpen"
          class="absolute z-50 mt-2 w-full min-w-[200px]
                 rounded-xl bg-white border border-emerald-100
                 shadow-lg shadow-emerald-500/10 overflow-hidden"
        >
          <button
            v-for="type in PROFILE_TYPES"
            :key="type.value"
            @click="selectType(type)"
            class="w-full flex items-center gap-3 px-4 py-3
                   hover:bg-emerald-50 transition text-left"
            :class="{ 'bg-emerald-50': selectedType === type.value }"
          >
            <div 
              class="w-7 h-7 rounded-lg flex items-center justify-center shrink-0" 
              :class="[type.bgColor, type.color]"
            >
              <Icon :name="type.icon" size="16" />
            </div>
            <span 
              class="text-sm font-medium"
              :class="selectedType === type.value ? 'text-emerald-700' : 'text-gray-700'"
            >
              {{ type.label }}
            </span>
            <Icon 
              v-if="selectedType === type.value"
              name="ph:check" 
              size="16" 
              class="ml-auto text-emerald-500"
            />
          </button>
        </div>
      </Transition>
    </div>
  </div>
</template>
