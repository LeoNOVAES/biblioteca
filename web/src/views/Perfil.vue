<template>
<div>
  <div>
  

  <b-modal id="modal-scrollable" scrollable title="Meus Dados">
    <p class="my-4">
        
    </p>
  </b-modal>
</div>
</div>
</template>

<script>

export default {

}
</script>

<style>

</style>
<template>
  <div>
      <a style="width:100%" class="btn btn-default"  @click="getDados" >Perfil</a>
   

    <b-modal
      v-model="show"
      title="Editar Dados"
      :header-bg-variant="green"
      :header-text-variant="white"
      :body-bg-variant="white"
      :body-text-variant="black"
      :footer-bg-variant="black"
      :footer-text-variant="white"
    >
      <b-container fluid>
         
           <form >  
               
               <div class="row">
                  <div class="form-group col-lg-6">
                      <div class="avatar"><img id="img" :src="avatar"/></div>
                 </div>
                  <div class="form-group col-lg-6">
                      <label>Mudar sua foto</label>
                      <input type="file" id="foto" @change="mudarFoto($event)"/>
                      
                 </div>

                 <div style="font-weight:bold; margin-top:10px; border:1px solid #000000"  class="form-group col-lg-12"> Dados </div>
                  <div class="form-group col-lg-6">
                    <label for="nome">Nome</label>
                    <input type="text" class="form-control" id="nome" value="h"  v-model="user.nome" aria-describedby="nome" />
                 </div>
                
                  <div style="font-weight:bold; margin-top:10px; border:1px solid #000000"  class="form-group col-lg-12"> Endereço </div>
                  <div class="form-group col-lg-6">
                    <label for="cidade">Cidade</label>
                    <input type="cidade" class="form-control" v-model="endereco.cidade" id="cidade" aria-describedby="cidade" placeholder="Cidade">
                 </div>
                  <div class="form-group col-lg-6">
                    <label for="rua">Rua</label>
                    <input type="text" class="form-control" v-model="endereco.rua" id="rua" aria-describedby="rua" placeholder="Rua">
                 </div> 
                  <div class="form-group col-lg-12">
                    <label for="bairro">Bairro</label>
                    <input type="text" class="form-control" v-model="endereco.bairro" id="bairro" aria-describedby="bairro" placeholder="bairro">
                 </div> 

                 
                 
                 <div style="font-weight:bold; margin-top:10px; border:1px solid #000000"  class="form-group col-lg-12"> Senhas </div>
                  <div class="form-group col-lg-6">
                    <label for="senha">Senha</label>
                    <input type="password" class="form-control" v-model="user.senha" id="senha" aria-describedby="senha" placeholder="Senha">
                 </div> 
                  <div class="form-group col-lg-6">
                    <label for="senhaConf">Confirmar Senha</label>
                    <input type="password" class="form-control" v-model="user.senhaConf" id="senhaConf" aria-describedby="senhaConf" placeholder="Senha Confirmacao">
                 </div>   
               </div>
            </form>
     
      </b-container>

      <div slot="modal-footer" class="w-100">
        <p class="float-left">DESEJA SALVAR AS ALTERAÇÕES?</p>
        <b-button size="sm" class="float-right" variant="danger" @click="show=false" style="margin-left:10px">Cancelar</b-button>
        <b-button size="sm" class="float-right" variant="success" @click="setDados">Salvar</b-button>
      </div>
    </b-modal>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        show: false,
        green: 'success',
        white: 'light',
        black:'dark',
        user:{
            nome:"teste",
            senha:'',
            senhaConf:''
        },
        endereco:{
            cidade:'',
            rua:'',
            bairro:'',
        },
        avatar:"",
        foto:""
      }
    },

    methods:{
        async getDados(){
            this.show = true
            const token = localStorage.getItem("token")
            const req = await fetch(`http://localhost:9000/dados/${this.$id}`,{
              headers:{
                Authorization:token
              }
            })
            const res = await req.json()
            this.$data.user.nome = res.name
            this.$data.endereco.cidade = res.cidade
            this.$data.endereco.rua = res.rua
            this.$data.endereco.bairro = res.bairro
            
            const reqF = await fetch(`http://localhost:9000/imagemE/${this.$id}`,{
              headers:{
                Authorization:token
              }
            })

            const resF = await reqF.json()
            if(resF == false){
                 this.$data.avatar = "http://localhost:9000/avatar/default.png"
            }else{
                  this.$data.avatar = "http://localhost:9000/getAvatar/"+this.$id
            }
            
        },

          
        async setDados(){

              let form = new FormData()

              if(this.$data.user.senha == this.$data.user.senhaConf){
                  form.append("password", this.$data.user.senha)
              }else{
                this.$swal.fire({
                  title: "Senhas Não conferem",
                  type: 'warning',
                  confirmButtonText: 'ok'
                })            
               
                return 
                 
              }

              form.append("nome", this.$data.user.nome)
              form.append("cidade", this.$data.endereco.cidade)
              form.append("rua", this.$data.endereco.rua)
              form.append("bairro", this.$data.endereco.bairro)


              const token = localStorage.getItem("token")
              const req = await fetch(`http://localhost:9000/alterar/${this.$id}`,{
                method:'POST',
                body:form,
                headers:{
                  Authorization:token
                }
              })
              let res = await req.json()
              
            this.$swal.fire({
                  title: "Sucesso ao Alterar Dados",
                  type: 'success',
            })
            
            this.show=false
        },

        async mudarFoto(e){
          this.$data.foto = e.target.files[0]
          document.getElementById('img').src = window.URL.createObjectURL(this.foto)
          let form = new FormData();

          form.append("avatar", this.$data.foto)
          const req = await fetch(`http://localhost:9000/upload/${this.$id}`,{
            method:"POST",
            body:form
          })

          const res = await req.json()
          console.log(res)

        }

      
        
    }
  }
</script>

<style >

  .avatar img{
    max-width: 7rem;
  }

  
  



</style>