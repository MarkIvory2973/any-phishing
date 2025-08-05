<script setup>
import { ref, reactive } from 'vue'
const props = defineProps(['name', 'description', 'url'])
const emit = defineEmits(['onRefresh'])

import z from 'zod'

import axios from 'axios'
import { apiBaseURL } from '@/envs'

import { Snackbar } from '@varlet/ui'

const dbClient = reactive({
  name: props.name,
  description: props.description,
  url: props.url,
})

async function actionCancel() {
  await uiClose()
}
async function actionEdit() {
  await uiValidate()
  Snackbar.loading('Editing client')

  try {
    const response = await axios.patch(apiBaseURL + '/clients/' + props.url, dbClient)
  } catch (error) {
    try {
      if (error.response.data.status != 'success') {
        Snackbar.error(error.response.data.message)
        return
      }
    } catch {
      Snackbar.error("Can't edit client")
      return
    }
  }

  Snackbar.success('Edited client')
  emit('onRefresh')
  await uiClose()
  return
}

const uiEdit = ref(null)
const uiShow = ref(false)

async function uiOpen() {
  uiShow.value = true
}
async function uiClose() {
  uiShow.value = false
}
async function uiValidate() {
  uiEdit.value.validate()
}
</script>

<template>
  <var-button title="Edit client" @click="uiOpen()" round text>
    <var-icon name="cog" />
  </var-button>

  <var-popup v-model:show="uiShow" @close="uiClose()">
    <var-form class="popup" ref="uiEdit">
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

        <var-space size="8px" justify="flex-end">
          <var-button type="primary" @click="actionCancel()" text>Cancel</var-button>
          <var-button type="primary" @click="actionEdit()" text>Edit</var-button>
        </var-space>
      </var-space>
    </var-form>
  </var-popup>
</template>

<style scoped>
#edit {
  margin: 18px;
}
</style>
