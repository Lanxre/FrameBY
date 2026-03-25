<script setup lang="ts">
import TabsLayout from '@/components/ui/TabsLayout/TabsLayout.vue'
import UniversitySquadApproval from './UniversitySquadApproval.vue'
import UniversityEmploymentRequests from './UniversityEmploymentRequests.vue'

const tabs = [
  {
    label: 'Заявки отрядов',
    value: 'squads',
    icon: 'ph:users-three',
    component: UniversitySquadApproval,
    hasPermission: true
  },
  {
    label: 'Заявки на работу',
    value: 'employment',
    icon: 'ph:briefcase',
    component: UniversityEmploymentRequests,
    hasPermission: true
  }
]
</script>

<template>
  <TabsLayout :tabs="tabs" default-tab="squads" />
</template>
