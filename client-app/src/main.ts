import "./assets/main.css";

import { createApp } from "vue";
import App from "./App.vue";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import router from "@/router/router";
import ruRU from "element-plus/es/locale/lang/ru";
import { createPinia } from "pinia";
import { storePlugin } from "pinia-plugin-store";

const pinia = createPinia();
const plugin = storePlugin({
  stores: ["EventStore"],
  storage: localStorage,
});
pinia.use(plugin);

const app = createApp(App);

app.use(ElementPlus, { locale: ruRU });
app.use(router);
app.use(pinia);
app.mount("#app");
