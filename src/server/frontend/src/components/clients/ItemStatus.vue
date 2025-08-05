<script setup>
import { ref, onMounted } from 'vue'
const props = defineProps(['url'])

import axios from 'axios'

async function updateStatus() {
  try {
    const response = await axios.get(props.url)

    uiStatus.value = 0
    return
  } catch (error) {
    try {
      if (!(200 <= error.response.status && error.response.status < 400)) {
        uiStatus.value = 2
        return
      }
    } catch {
      uiStatus.value = 1
      return
    }
  }
}

const uiStatus = ref(1)

setInterval(() => {
  updateStatus()
}, 20000)

onMounted(() => {
  updateStatus()
})
</script>

<template>
  <var-icon
    name="checkbox-marked-circle"
    title="Alive"
    color="var(--color-success)"
    v-if="uiStatus == 0"
  />
  <var-icon
    name="close-circle"
    title="Dead"
    color="var(--color-danger)"
    v-else-if="uiStatus == 1"
  />
  <var-icon
    name="error"
    title="Alive but error"
    color="var(--color-warning)"
    v-else-if="uiStatus == 2"
  />
</template>
