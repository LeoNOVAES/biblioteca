import Vue from "vue"
import Router from "vue-router"

const Login = () => import("@/views/Login")
const Home = () => import("@/views/Home")
const Cadastro = () => import("@/views/Cadastro")

Vue.use(Router)

var router = new Router({
    routes :[
        {
            path:"/",
            redirect:'/login',
            name:"Login",
            component: Login,
            children:[
                {
                    path:'login',
                    name:'Login',
                    component:Login
                }
            ]
        },
        {
            path:"/home",
            name:"Home",
            beforeEnter:(to,from,next)=>{
                const token = localStorage.getItem("token")
                if (!token){
                    next('/login')
                }else{
                    next()
                }
            },
            component:Home
           
           
        },
        {
            path:"/cadastro",
            name:"Cadastro",
            component:Cadastro
        }
    ]
})

export default router
