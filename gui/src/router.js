import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'about',
      component: () => import( './views/About.vue')
    },

    {
      path: '/parameters',
      name: 'parameters',
      component: () => import('./views/Parameters.vue')
    },

    {
      path: '/invitations',
      name: 'invitations',
      component: () => import('./views/Invitations.vue')
    },
    

    {
      path: '/performance',
      name: 'performance',
      component: () => import('./views/Performance.vue')
    },
    {
      path: '/network',
      name: 'network',
      component: () => import('./views/Network.vue')
    },

    {
      path: '/swagger',
      name: 'swagger',
      component: () => import('./views/Swagger.vue')
    },

    {
      path: '/setup',
      name: 'setup',
      component: () => import('./views/Setup.vue')
    },

    {
      path: '/oidc-callback',
      name: 'callback',
      component: () => import('./views/Callback.vue')
    },

  ]

})
