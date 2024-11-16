import {createRouter, createWebHistory} from 'vue-router';
import SignUp from "../views/SignUp.vue";
import Admin from '../views/Admin.vue';

const routes = [
  {
    path: '/',
    name: 'SignUp',
    component: SignUp
  },
  {
    path: '/agentreqs',
    name: 'AgentRequests',
    component: Admin
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;