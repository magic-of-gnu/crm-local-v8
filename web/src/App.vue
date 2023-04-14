<template>
  <router-view />
</template>

<script setup>
import store from '@/store/index.js'

store.dispatch('tryLogin')

</script>

<style lang="scss">
// Import Main styles for this application
@import 'styles/style';
</style>
