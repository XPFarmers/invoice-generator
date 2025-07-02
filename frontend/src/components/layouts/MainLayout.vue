<script setup lang="ts">
import ApplicationBar from '@/components/ApplicationBar.vue'
import { RouterView, useRouter, useRoute } from 'vue-router'
import { watchEffect } from 'vue'
import { useNavigationStore } from '@/stores/navigation'
import { Colours } from '@/utils/colours'

const router = useRouter()
const route = useRoute()
const nav = useNavigationStore()

const tabs = [
  { icon: 'mdi-list-box-outline', title: 'Dashboard', value: 'dashboard', colour: Colours.Crimson },
  { icon: 'mdi-invoice-text-plus-outline', title: 'Generate Invoice', value: 'invoice', colour: Colours.ForestGreen },
  { icon: 'mdi-archive-outline', title: 'Archived', value: 'archived', colour: Colours.DodgerBlue },
]

// Sync current route with selected nav
watchEffect(() => {
  const current = route.path.replace('/', '')
  nav.navigate(current)
})

const navigateToView = (view: string) => {
  nav.navigate(view)
  router.push({ path: view })
}
</script>

<template>
  <ApplicationBar />

  <v-navigation-drawer expand-on-hover rail fixed permanent :width="250">
    <v-list>
      <v-list-item v-for="tab in tabs" :key="tab.value" :prepend-icon="tab.icon" :title="tab.title" :value="tab.value"
        :active="nav.section === tab.value" :color="tab.colour" class="rounded-pill non-selectable mt-2"
        @click="navigateToView(tab.value)" />
    </v-list>
  </v-navigation-drawer>

  <v-container fluid class="scrollable-container">
    <RouterView />
  </v-container>
</template>

<style>
.scrollable-container {
  max-height: calc(100vh - 80px);
  overflow-y: auto;
}
</style>
