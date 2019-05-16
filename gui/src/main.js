import Vue from 'vue'
import './plugins/vuetify'
import store from '@/store'
import App from './App.vue'
import router from './router'

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')

// Reset errors on each navigation change
router.beforeEach((to, from, next) => {
  console.log("Navigating to ", to.fullPath)

  // clear errors when we change the screen
  store.dispatch("setError", null)
  store.dispatch("setProcessing", false)
  next();

})