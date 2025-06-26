<template>
  <v-app-bar>
    <v-app-bar-nav-icon variant="text" icon="mdi-home" @click="navigateToDashboard"></v-app-bar-nav-icon>
    <v-app-bar-title class="non-selectable">
      Invoice Generator
    </v-app-bar-title>

    <div class="text-center">
      <span class="non-selectable mr-8">v0.0.1</span>
    </div>

  </v-app-bar>
</template>

<script setup lang="ts">

import { useRouter } from 'vue-router';


const router = useRouter();


const navigateToDashboard = () => {
  router.push('/dashboard');
};



</script>

