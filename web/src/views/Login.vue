<template >
   <b-container fluid class="body" >
      <div id="login">
        <b-card
          class="cardLog"
          header="LOGIN"
          header-text-variant="black"
          header-tag="header"
          header-bg-variant="light"
        >
          <b-card-text>
            <form v-on:submit="isLogin($event)">  
               <div class="row">
                 <div class="form-group col-lg-6">
                    <label for="email">Email</label>
                    <input type="email" v-model="user.email" class="form-control" id="email" aria-describedby="email" placeholder="Enter email">
                 </div>  
                  <div class="col-lg-6">
                    <label for="password">Senha</label>
                    <input type="password" v-model="user.password" class="form-control" id="password" aria-describedby="password" placeholder="Password">
                 </div> 
                 <div class="form-group col-lg-6" style="margin-top:10px">
                    <button  class="btn btn-success">Entrar</button>
                 </div>  
                 <div class="col-lg-6" style="margin-top:10px">
                    <a class="btn btn-dark" href="#/cadastro">Cadastre-se</a>
                 </div>   
               </div>
            </form>
          </b-card-text>
          
        </b-card>
      </div>
  </b-container>
</template>

<script>
export default {
  
  data(){
    return{
      user:{
        email:'',
        password:''
      }
    }
  },
  methods:{
  
    async isLogin(){
        var formDta = new FormData()
        formDta.append("email", this.user.email)
        formDta.append("password", this.user.password)
        const rest = await fetch('http://localhost:9000/isLogin/',{
          method:"POST",
          body: formDta
        })

        const res = await rest.json()
        if(res != "Email ou Senha Incorretos"){
          localStorage.setItem("token", res)
          localStorage.setItem("email", formDta.get("email"))
         
        }else{
          this.$swal.fire({
            title: 'Error',
            text: res,
            type: 'warning',
            confirmButtonColor: '#B22222',
            confirmButtonText: 'OK'
          })
        }
          window.location.href = "#/home"
    }
  }
}
</script>

<style scoped>
  #login{
      display:flex;
      flex-direction: row;
      justify-content: center;
      align-items: center;
      
  }
  .cardLog{
    top:270px;
    width:30%;
    left:20px

  }
  .body{
      background-image: url("../../public/img/fundo.jpg");
      background-repeat: no-repeat;
      background-size: 100% 100%;
      height:797px
    }
    
</style>
