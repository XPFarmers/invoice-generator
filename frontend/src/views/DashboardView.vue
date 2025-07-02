<script setup lang="ts">
import { Colours } from '@/utils/colours'
import { useRouter } from 'vue-router'

const router = useRouter()
const goToInvoice = () => {
  router.push('/invoice')
}
</script>

<template>

  Dashboard
  <v-fab :color="Colours.ForestGreen" icon="mdi-plus" location="right bottom" @click="goToInvoice" app></v-fab>

</template>
