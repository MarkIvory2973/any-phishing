<script setup>
const props = defineProps(['url'])
const emit = defineEmits(['onRefresh'])

import axios from 'axios'
import { apiBaseURL } from '@/envs'

import { Dialog, Snackbar } from '@varlet/ui'

async function onBeforeClose(action, done) {
  if (action != 'confirm') {
    done()
    return
  }

  Snackbar.loading('Deleting client')

  try {
    const response = await axios.delete(apiBaseURL + '/clients/' + props.url)
  } catch (error) {
    try {
      if (error.response.data.status != 'success') {
        Snackbar.error(error.response.data.message)
        done()
        return
      }
    } catch {
      Snackbar.error("Can't delete client")
      done()
      return
    }
  }

  Snackbar.success('Deleted client')
  emit('onRefresh')
  done()
  return
}
async function actionDelete() {
  Dialog({
    title: 'Warning',
    message: 'Delete client?',
    onBeforeClose,
  })
}
</script>

<template>
  <var-button title="Delete client" @click="actionDelete()" round text>
    <var-icon name="delete" />
  </var-button>
</template>
