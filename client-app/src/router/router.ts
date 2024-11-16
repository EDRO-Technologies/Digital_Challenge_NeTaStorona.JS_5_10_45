import { createRouter, createWebHistory } from "vue-router";
import App from "@/App.vue";
import SupportView from "@/views/SupportView.vue";
import EventsView from "@/views/EventsView.vue";

const routes = [
  {
    path: "/",
    component: App,
  },
  {
    path: "/support",
    component: SupportView,
  },
  {
    path: "/events",
    component: EventsView,
  },
];

const Router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default Router;