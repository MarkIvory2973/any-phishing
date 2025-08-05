<script setup>
import { ref, reactive } from 'vue'
const emit = defineEmits(['onRefresh'])

import z from 'zod'

import axios from 'axios'
import { apiBaseURL } from '@/envs'

import { Snackbar } from '@varlet/ui'

const dbClient = reactive({
  name: '',
  description: '',
  sessions: [],
  url: '',
})

async function actionCancel() {
  await uiClose()
}
async function actionAdd() {
  console.log(await uiValidate())
  Snackbar.loading('Adding client')

  try {
    const response = await axios.put(apiBaseURL + '/clients/' + dbClient.url, dbClient)
  } catch (error) {
    try {
      if (error.response.data.status != 'success') {
        Snackbar.error(error.response.data.message)
        return
      }
    } catch {
      Snackbar.error("Can't add client")
      return
    }
  }

  Snackbar.success('Added client')
  emit('onRefresh')
  await uiClose()
  return
}

const uiAdd = ref(null)
const uiShow = ref(false)

async function uiOpen() {
  uiShow.value = true
}
async function uiClose() {
  uiAdd.value.reset()

  uiShow.value = false
}
async function uiValidate() {
  uiAdd.value.validate()
}
</script>

<template>
  <var-fab type="primary" @click="uiOpen()" />

  <var-popup v-model:show="uiShow" @close="uiClose()">
    <var-form class="popup" ref="uiAdd">
      <var-space direction="column" size="24px">
        <var-space direction="column" size="24px">
          <var-input
            variant="outlined"
            placeholder="Name"
            :rules="z.string().min(1, '* Name is required')"
            v-model="dbClient.name"
          />
          <var-input variant="outlined" placeholder="Description" v-model="dbClient.description" />
          <var-input
            variant="outlined"
            placeholder="URL"
            :rules="z.url('* Invalid URL')"
            v-model="dbClient.url"
          />
        </var-space>

        <var-space justify="flex-end" size="8px">
          <var-button type="primary" @click="actionCancel()" text>Cancel</var-button>
          <var-button type="primary" @click="actionAdd()" text>Add</var-button>
        </var-space>
      </var-space>
    </var-form>
  </var-popup>
</template>
