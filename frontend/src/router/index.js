import { createRouter, createWebHistory } from 'vue-router'
import EventsView from '../views/EventsView.vue'
import DataView from '../views/DataView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'events',
      component: EventsView,
    },
    {
      path: '/registrations',
      name: 'registrations',
      component: () => import('../views/RegistrationsView.vue'),
    },
    {
      path: '/data',
      name: 'data',
      component: DataView,
    },
    {
      path: '/visits',
      name: 'visits',
      component: () => import('../views/VisitsView.vue'),
    },
    {
      path: '/admin/dashboard',
      name: 'admin',
      component: () => import('../views/AdminMainView.vue'),
    },
    {
      path: '/admin/registrations',
      name: 'admin-registrations',
      component: () => import('../views/AdminRegistrationsView.vue'),
    },
    {
      path: '/admin/registration/:id',
      name: 'admin-registration',
      component: () => import('../views/AdminRegistrationView.vue'),
    },
    {
      path: '/admin/visitors',
      name: 'visitors',
      component: () => import('../views/AdminVisitorsView.vue'),
    },
    {
      path: '/admin/event/:id',
      name: 'event',
      component: () => import('../views/AdminEventView.vue'),
    },
    {
      path: '/admin/login',
      name: 'login',
      component: () => import('../views/AdminLoginView.vue'),
    },
    {
      path: '/tools',
      name: 'tools',
      component: () => import('../views/ToolsView.vue'),
    },
    {
      path: '/qr/:id',
      name: 'qr',
      component: () => import('../views/QrView.vue'),
    },
    {
      path: '/v/:id',
      name: 'visit',
      component: () => import('../views/VisitView.vue'),
    },
    {
      path: '/r/:id',
      name: 'reg',
      component: () => import('../views/RegistrationView.vue'),
    },
  ],
})

export default router
