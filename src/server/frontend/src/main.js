import './assets/main.css'
import '@varlet/ui/es/style'

import { createApp } from 'vue'
import App from './App.vue'

import router from './router'

import Varlet from '@varlet/ui'
import { Themes, StyleProvider } from '@varlet/ui'
import { Locale } from '@varlet/ui'
import '@varlet/touch-emulator'

const app = createApp(App)

app.use(router).use(Varlet)

StyleProvider({
  ...Themes.md3Dark,
  '--hsl-primary': '208, 100%, 79%',
  '--color-primary': 'hsla(var(--hsl-primary), 1)',
  '--hsl-on-primary': '202, 100%, 16%',
  '--color-on-primary': 'hsla(var(--hsl-on-primary), 1)',
  '--hsl-primary-container': '201, 100%, 23%',
  '--color-primary-container': 'hsla(var(--hsl-primary-container), 1)',
  '--hsl-on-primary-container': '211, 100%, 90%',
  '--color-on-primary-container': 'hsla(var(--hsl-on-primary-container), 1)',
  '--hsl-info': '212, 31%, 79%',
  '--color-info': 'hsla(var(--hsl-info), 1)',
  '--hsl-on-info': '208, 29%, 19%',
  '--color-on-info': 'hsla(var(--hsl-on-info), 1)',
  '--hsl-info-container': '210, 21%, 28%',
  '--color-info-container': 'hsla(var(--hsl-info-container), 1)',
  '--hsl-on-info-container': '212, 65%, 90%',
  '--color-on-info-container': 'hsla(var(--hsl-on-info-container), 1)',
  '--hsl-warning': '267, 45%, 83%',
  '--color-warning': 'hsla(var(--hsl-warning), 1)',
  '--hsl-on-warning': '264, 28%, 23%',
  '--color-on-warning': 'hsla(var(--hsl-on-warning), 1)',
  '--hsl-warning-container': '264, 20%, 32%',
  '--color-warning-container': 'hsla(var(--hsl-warning-container), 1)',
  '--hsl-on-warning-container': '269, 100%, 93%',
  '--color-on-warning-container': 'hsla(var(--hsl-on-warning-container), 1)',
  '--hsl-danger': '6, 100%, 84%',
  '--color-danger': 'hsla(var(--hsl-danger), 1)',
  '--hsl-on-danger': '357, 100%, 21%',
  '--color-on-danger': 'hsla(var(--hsl-on-danger), 1)',
  '--hsl-danger-container': '356, 100%, 29%',
  '--color-danger-container': 'hsla(var(--hsl-danger-container), 1)',
  '--hsl-on-danger-container': '6, 100%, 84%',
  '--color-on-danger-container': 'hsla(var(--hsl-on-danger-container), 1)',
  '--hsl-body': '210, 7%, 11%',
  '--color-body': 'hsla(var(--hsl-body), 1)',
  '--hsl-text': '240, 5%, 89%',
  '--color-text': 'hsla(var(--hsl-text), 1)',
  '--hsl-on-surface-variant': '215, 11%, 78%',
  '--color-on-surface-variant': 'hsla(var(--hsl-on-surface-variant), 1)',
  '--hsl-inverse-surface': '240, 5%, 89%',
  '--color-inverse-surface': 'hsla(var(--hsl-inverse-surface), 1)',

  '--table-thead-th-text-align': 'center',
  '--table-tbody-td-text-align': 'center',
})
Locale.add('en-US', Locale.enUS)
Locale.use('en-US')

app.mount('#app')
