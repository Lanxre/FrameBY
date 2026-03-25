<script setup lang="ts">
import TabsLayout from '@/components/ui/TabsLayout/TabsLayout.vue'
import CustomerEmploymentCreate from './CustomerEmploymentCreate.vue'
import CustomerEmploymentList from './CustomerEmploymentList.vue'

const tabs = [
  {
    label: 'Создать заявку',
    value: 'create',
    icon: 'ph:plus-circle',
    component: CustomerEmploymentCreate,
    hasPermission: true
  },
  {
    label: 'Мои заявки',
    value: 'list',
    icon: 'ph:briefcase',
    component: CustomerEmploymentList,
    hasPermission: true
  }
]
</script>

<template>
  <TabsLayout :tabs="tabs" default-tab="create" active-tab-class="bg-amber-50 text-amber-600" />
</template>
