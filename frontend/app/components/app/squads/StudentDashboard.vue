<script setup lang="ts">
import TabsLayout from '@/components/ui/TabsLayout/TabsLayout.vue'
import StudentSquadList from './StudentSquadList.vue'
import MySquadMemberships from './MySquadMemberships.vue'

const tabs = [
  {
    label: 'Наборы отрядов',
    value: 'all',
    icon: 'ph:users-three',
    component: StudentSquadList,
    hasPermission: true
  },
  {
    label: 'Мои отряды',
    value: 'mine',
    icon: 'ph:user-circle',
    component: MySquadMemberships,
    hasPermission: true
  }
]
</script>

<template>
  <TabsLayout :tabs="tabs" default-tab="all" />
</template>
