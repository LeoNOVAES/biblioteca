
// import cssVars from 'css-vars-ponyfill'
import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'
import App from './App'
import router from './router'
import sweet from 'vue-sweetalert2'
import resource from 'vue-resource'

Vue.use(BootstrapVue)
Vue.use(sweet)
Vue.use(VueSweetalert2)
Vue.use(resource)

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import VueSweetalert2 from 'vue-sweetalert2';
/* eslint-disable no-new */








new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {
    App
  }
})
