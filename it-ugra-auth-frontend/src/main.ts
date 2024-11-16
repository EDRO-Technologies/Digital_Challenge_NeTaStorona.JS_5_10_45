import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import ruRU from 'element-plus/es/locale/lang/ru'

import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';
import router from "./router";

const app = createApp(App)

app.use(ElementPlus, {
  locale: ruRU
})
app.use(router);

app.component('QuillEditor', QuillEditor);

app.mount('#app')