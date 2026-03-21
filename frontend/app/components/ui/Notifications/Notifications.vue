<script lang="ts" setup>
import { storeToRefs } from "pinia";
import NotificationItem from "@/components/ui/Notifications/NotificationItem.vue";
import { useNotificationStore } from "@/stores/notification";

const store = useNotificationStore();
const { notifications } = storeToRefs(store);
</script>

<template>
    <div class="fixed top-24 right-8 z-9999 flex flex-col gap-2 w-full max-w-sm pointer-events-none p-4 sm:p-0 overflow-x-hidden">
        <TransitionGroup 
            enter-active-class="transform ease-out duration-300 transition"
            enter-from-class="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-4"
            enter-to-class="translate-y-0 opacity-100 sm:translate-x-0"
            leave-active-class="transition ease-in duration-300 absolute w-full"
            leave-from-class="opacity-100 translate-x-0"
            leave-to-class="opacity-0 translate-x-full"
            move-class="transition-all duration-300"
        >
            <NotificationItem
                v-for="notification in notifications"
                :key="notification.id"
                :notification="notification"
                @remove="store.remove"
            />
        </TransitionGroup>
    </div>
</template>