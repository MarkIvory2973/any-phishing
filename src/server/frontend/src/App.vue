<script setup>
import { RouterView } from 'vue-router'

import NavItem from './components/NavItem.vue'
</script>

<template>
  <NavItem />

  <main>
    <RouterView />
  </main>
</template>

<style scoped>
main {
  margin: var(--space-normal);
}
</style>
