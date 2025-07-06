<script setup lang="ts">
import { Colours } from '@/utils/colours'
import { useRouter } from 'vue-router'
import DashboardComponent from '@/components/DashboardComponent.vue'
const router = useRouter()
const goToInvoice = () => {
  router.push('/invoice')
}
</script>

<template>
  <DashboardComponent />
  <v-fab
    :color="Colours.ForestGreen"
    icon="mdi-plus"
    location="right bottom"
    @click="goToInvoice"
    app
  ></v-fab>
</template>
