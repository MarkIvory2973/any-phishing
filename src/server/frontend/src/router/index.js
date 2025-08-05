import { createRouter, createWebHistory } from 'vue-router'

import ClientsView from '../views/ClientsView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'clients',
      component: ClientsView,
    },
  ],
})

export default router
