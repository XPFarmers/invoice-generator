<template>
  <v-app>
    <v-main>
      <RouterView />
    </v-main>
  </v-app>
</template>

<style>
.non-selectable {
  -webkit-user-select: none;
  /* Safari, Chrome, Opera */
  -khtml-user-select: none;
  /* Konqueror */
  -moz-user-select: none;
  /* Firefox */
  -ms-user-select: none;
  /* Internet Explorer, Edge */
  -o-user-select: none;
  /* Opera */
  user-select: none;
  /* Standard syntax */
}
</style>
