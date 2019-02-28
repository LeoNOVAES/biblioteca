<template>
    <div>
        <Header :avatar="this.$id"/>
        <div class="container-fluid bg-success banner" >
            
        </div>  
    </div>
</template>

<script>
import Header from "@/views/Header"
import Vue from "vue"

export default {
    components:{Header},
    data(){
        image:''
    },
    methods:{
        async getPrototypes(){
            let token = localStorage.getItem("token")
            const req = await  fetch(`http://localhost:9000/getUser/${localStorage.getItem("email")}`,{
                headers:{
                Authorization:token
                },
                method:"GET"
            })
            const res = await req.json()
                console.log(res)
                Vue.prototype.$nome = res.nome  
                Vue.prototype.$email = res.email
                Vue.prototype.$id = res.id
                this.$data.image = res.id
            }
    },
    created(){
        this.getPrototypes()
    }
    
}
</script>

<style >
    .banner{
       background-image: url("../../public/img/banner.jpg");
       background-repeat: no-repeat;
       background-size: 100% 100%;
       height:450px
    }
</style>
