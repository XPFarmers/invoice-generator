<script setup lang="ts">
import ApplicationBar from '@/components/ApplicationBar.vue'
import { RouterView, useRouter } from 'vue-router'

const router = useRouter();


const tabs = [
  { icon: 'mdi-list-box-outline', title: 'Dashboard', value: 'dashboard', colour: '#dc143c' },
  { icon: 'mdi-invoice-text-plus-outline', title: 'Generate Invoice', value: 'invoice', colour: '#32cd32' },
  { icon: 'mdi-list-box-outline', title: 'Archived', value: 'archived', colour: '#1e90ff' },
];

const navigateToView = (view: string) => {
  router.push({ path: view });
};

</script>

<style>
.scrollable-container {
  max-height: calc(100vh - 80px);
  overflow-y: auto;
}
</style>

<template>

  <ApplicationBar />

  <v-navigation-drawer expand-on-hover rail fixed permanent :width="250">

    <v-list>
      <v-list-item v-for="tab in tabs" :key="tab.value" :prepend-icon="tab.icon" :title="tab.title" :value="tab.value"
        :color="tab.colour" style="" class="rounded-pill non-selectable mt-2"
        @click="navigateToView(tab.value)"></v-list-item>
    </v-list>

  </v-navigation-drawer>
  <v-container fluid class="scrollable-container">
    <RouterView />
  </v-container>
</template>
