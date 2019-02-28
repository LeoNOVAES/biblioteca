<template>
<div>
  <b-navbar toggleable="md" class="menu" type="light" variant="dark">
    <b-navbar-brand ><h3  style="color:#ffffff; font-weight:bold">HOME</h3></b-navbar-brand>

       <b-navbar-nav nav class="ml-4">
        <b-nav-item href="#" class="bg-success" style="margin-right:10px"><p style="color:#ffffff; font-weight:bold">Livros</p></b-nav-item>
        <b-nav-item href="#" class="bg-success" ><p style="color:#ffffff; font-weight:bold">Locações</p></b-nav-item>
      </b-navbar-nav>
     <b-navbar-nav class="ml-auto">
       <b-nav-item-dropdown right>
          <!-- Using button-content slot -->
          <template  slot="button-content" ><em style="color:#ffffff"><img :src="avatar" width="50" height="50" /></em></template>
            <Perfil/>
          <a style="width:100%" class="btn btn-default" @click="sair">Sair</a>
        </b-nav-item-dropdown>
    </b-navbar-nav> 
  </b-navbar>

  
</div>
</template>

<script>
import Perfil from "@/views/Perfil"
export default {

  components:{
    Perfil
  },
  props:["avatar"],
  data(){
    return {
      

    }
  },
  methods:{
        sair(){
            localStorage.removeItem("token")
             location.reload();     
        },
         async getAvatar(){
          alert(this.avatar)
        }
  },
  mounted(){
    this.getAvatar()
  }
 

}
</script>

<style>
  .menu{
    height: 70px
  }
</style>
