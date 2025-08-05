<script setup>
const props = defineProps(['url'])

import { Snackbar } from '@varlet/ui'

async function actionCopy() {
  Snackbar.loading('Copying URL to clipboard')

  try {
    await navigator.clipboard.writeText(props.url)
  } catch {
    Snackbar.error("Can't copy URL to clipboard")
  }

  Snackbar.success('Copied URL to clipboard')
}
</script>

<template>
  <var-button title="Copy URL" @click="actionCopy()" round text>
    <var-icon name="content-copy" />
  </var-button>
</template>
