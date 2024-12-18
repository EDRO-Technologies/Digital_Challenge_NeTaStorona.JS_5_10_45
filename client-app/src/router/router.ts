import { createRouter, createWebHistory } from "vue-router";
import App from "@/App.vue";
import SupportView from "@/views/SupportView.vue";
import EventsView from "@/views/EventsView.vue";
import SurveyView from "@/views/SurveyView.vue";
import ProfileView from "@/views/ProfileView.vue";
import Auther from "@/views/Auther.vue";

const routes = [
  {
    path: "/",
    redirect: '/events',
  },
  {
    path: "/auth",
    component: Auther,
  },
  {
    path: "/support",
    component: SupportView,
  },
  {
    path: "/events",
    component: EventsView,
  },
  {
    path: "/survey",
    component: SurveyView,
  },
  {
    path: "/profile",
    component: ProfileView,
  }
];

const Router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default Router;
