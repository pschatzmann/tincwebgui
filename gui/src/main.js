import Vue from 'vue'
import './plugins/vuetify'
import store from '@/store'
import App from './App.vue'
import router from './router'
import { vuexOidcCreateRouterMiddleware } from 'vuex-oidc'

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')




// Reset errors on each navigation change
router.beforeEach(vuexOidcCreateRouterMiddleware(store))
router.beforeEach((to, from, next) => {
  console.log("Navigating to ", to.fullPath)

  // clear errors when we change the screen
  store.dispatch("setError", null)
  store.dispatch("setProcessing", false)
  next();

})