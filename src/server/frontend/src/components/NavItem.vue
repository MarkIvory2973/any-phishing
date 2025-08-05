<script setup>
import router from '@/router'
import { RouterLink } from 'vue-router'

async function actionBack() {
  router.back()
}
</script>

<template>
  <var-sticky :offset-top="0">
    <var-app-bar title="AnyPhishing" round>
      <template #left>
        <var-button
          class="left"
          color="transparent"
          text-color="#fff"
          @click="actionBack()"
          round
          text
        >
          <var-icon name="chevron-left" :size="24" />
        </var-button>
      </template>

      <template #right>
        <var-menu>
          <var-button color="transparent" text-color="#fff" round text>
            <var-icon name="menu" :size="24" />
          </var-button>

          <template #menu>
            <RouterLink to="/">
              <var-cell ripple>Clients</var-cell>
            </RouterLink>
          </template>
        </var-menu>
      </template>
    </var-app-bar>
  </var-sticky>
</template>

<style>
button.left {
  margin-right: 6px;
}
</style>
