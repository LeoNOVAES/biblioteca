
// import cssVars from 'css-vars-ponyfill'
import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'
import App from './App'
import router from './router'
import sweet from 'vue-sweetalert2'

Vue.use(BootstrapVue)
Vue.use(sweet)
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import VueSweetalert2 from 'vue-sweetalert2';
/* eslint-disable no-new */

Vue.use(VueSweetalert2)

async function getPrototypes(){
  let token = localStorage.getItem("token")
  const req = await fetch(`http://localhost:9000/getUser/${localStorage.getItem("email")}`,{
    method:"GET",
    headers:{Authorization:token}
  })
  const res = await req.json()
  console.log(res)
  return res

}

getPrototypes()

let dados = getPrototypes()
Vue.prototype.$nome = dados.nome
Vue.prototype.$email = dados.email
Vue.prototype.$id = dados.id

new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {
    App
  }
})
