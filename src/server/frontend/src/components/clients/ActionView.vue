<script setup>
import { ref } from 'vue'
const props = defineProps(['url'])

import { Snackbar } from '@varlet/ui'

import axios from 'axios'
import { apiBaseURL } from '@/envs'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
dayjs.extend(relativeTime)

const sessions = ref([])

async function updateSessions() {
  try {
    const response = await axios.get(apiBaseURL + '/sessions/' + props.url)

    if (response.data.data) {
      response.data.data.sort((a, b) => b.time - a.time)
    }

    sessions.value = response.data.data
  } catch (error) {
    try {
      if (error.response.data.status != 'success') {
        Snackbar.error(error.response.data.message)
        return
      }
    } catch {
      Snackbar.error("Can't view sessions")
      return
    }
  }
}

async function actionView() {
  await updateSessions()

  await uiOpen()
}

const uiShow = ref(false)
const uiCurrent = ref(1)
const uiSize = ref(10)

async function uiOpen() {
  uiShow.value = true
}
async function uiClose() {
  uiShow.value = false
}
async function uiUpdate(current, size) {
  uiCurrent.value = current
  uiSize.value = size
}

setInterval(() => {
  if (uiShow.value) {
    updateSessions()
  }
}, 3000)
</script>

<template>
  <var-button title="View sessions" @click="actionView()" round text>
    <var-icon name="view" />
  </var-button>

  <var-popup v-model:show="uiShow" @close="uiClose()">
    <var-table>
      <thead>
        <tr>
          <th>IP</th>
          <th>Time</th>
          <th>Inputs</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="session in sessions.slice((uiCurrent - 1) * uiSize, uiCurrent * uiSize)">
          <th>{{ session.ip }}</th>
          <th>{{ dayjs(session.time * 1000).fromNow() }}</th>
          <th>{{ session.inputs.join(', ') }}</th>
        </tr>
      </tbody>

      <template #footer>
        <div class="pagination">
          <var-pagination
            :simple="false"
            :total="sessions.length"
            :show-size-changer="false"
            @change="uiUpdate"
          />
        </div>
      </template>
    </var-table>
  </var-popup>
</template>
