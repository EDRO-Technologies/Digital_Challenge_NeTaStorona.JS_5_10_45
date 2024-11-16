import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import ruRU from 'element-plus/es/locale/lang/ru'

const app = createApp(App)

app.use(ElementPlus, {
  locale: ruRU
})
app.mount('#app')