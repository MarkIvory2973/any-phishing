<script setup>
import { ref, onMounted } from 'vue'

import { Snackbar } from '@varlet/ui'

import ItemStatus from '@/components/clients/ItemStatus.vue'
import ActionView from '@/components/clients/ActionView.vue'
import ActionCopy from '@/components/clients/ActionCopy.vue'
import ActionEdit from '@/components/clients/ActionEdit.vue'
import ActionDelete from '@/components/clients/ActionDelete.vue'
import ActionAdd from '@/components/clients/ActionAdd.vue'

import axios from 'axios'
import { apiBaseURL } from '@/envs'

const clients = ref([])

async function updateClients() {
  try {
    const response = await axios.get(apiBaseURL + '/clients')

    clients.value = response.data.data
  } catch (error) {
    try {
      if (error.response.data.status != 'success') {
        Snackbar.error(error.response.data.message)
        return
      }
    } catch {
      Snackbar.error("Can't update clients")
      return
    }
  }
}

const uiCurrent = ref(1)
const uiSize = ref(10)

async function uiUpdate(current, size) {
  uiCurrent.value = current
  uiSize.value = size
}

setInterval(() => {
  updateClients()
}, 3000)

onMounted(() => {
  updateClients()
})
</script>

<template>
  <span class="subtitle">Clients</span>

  <var-table>
    <thead>
      <tr>
        <th>Name</th>
        <th>Description</th>
        <th>Sessions</th>
        <th>URL</th>
        <th>Status</th>
        <th>Actions</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="client in clients.slice((uiCurrent - 1) * uiSize, uiCurrent * uiSize)">
        <td>{{ client.name }}</td>
        <td>{{ client.description }}</td>
        <td v-if="client.sessions == null">0</td>
        <td v-else>{{ client.sessions.length }}</td>
        <td>{{ client.url }}</td>
        <td>
          <ItemStatus :url="client.url" />
        </td>
        <td>
          <ActionView :url="client.url" />
          <ActionCopy :url="client.url" />
          <ActionEdit
            :name="client.name"
            :description="client.description"
            :url="client.url"
            v-on:on-refresh="updateClients"
          />
          <ActionDelete :url="client.url" v-on:on-refresh="updateClients" />
        </td>
      </tr>
    </tbody>

    <template #footer>
      <div class="pagination">
        <var-pagination
          :simple="false"
          :total="clients.length"
          :show-size-changer="false"
          @change="uiUpdate"
        />
      </div>
    </template>
  </var-table>

  <ActionAdd v-on:on-refresh="updateClients" />
</template>
