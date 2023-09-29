import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import AccountsView from '@/views/AccountsView.vue'
import DomainsView from '@/views/DomainsView.vue'
import DomainsAliasesView from '@/views/DomainsAliasesView.vue'
import AddressesAliasesView from '@/views/AddressesAliasesView.vue'
import { Authenticator } from '@/repositories/authenticator'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
      props: route => ({redirect: route.query.redirect}),
      meta: {
        requiresAuth: false
      }
    },
    {
      path: '/accounts',
      name: 'accounts',
      component: AccountsView,
      props: route => ({domain: route.query.domain}),
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/domains',
      name: 'domains',
      component: DomainsView,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/domains-aliases',
      name: 'domains-aliases',
      component: DomainsAliasesView,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/addresses-aliases',
      name: 'addresses-aliases',
      component: AddressesAliasesView,
      meta: {
        requiresAuth: true
      }
    }
  ]
})

// eslint-disable-next-line @typescript-eslint/no-unused-vars
router.beforeEach(async (to, _from) => {
  if (to.meta.requiresAuth && !Authenticator.isLoggedIn()) {
    return {
      name: 'login',
      query: { redirect: to.fullPath }
    }
  }

  return true
})

export default router
