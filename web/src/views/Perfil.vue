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
      <a style="width:100%" class="btn btn-default"  @click="show=true">Perfil</a>
   

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
                    <div style="width:50%;height:100px; background-color:#dddddd"><img src="img.png"/></div>
                 </div>

                 <div style="font-weight:bold; margin-top:10px; border:1px solid #000000"  class="form-group col-lg-12"> Dados </div>
                  <div class="form-group col-lg-6">
                    <label for="nome">Nome</label>
                    <input type="text" class="form-control" id="nome" v-model="user.senha" aria-describedby="nome" placeholder="Nome">
                 </div>
                
                  <div class="form-group col-lg-6">
                    <label for="email">Email</label>
                    <input type="email" class="form-control" v-model="user.email" id="email" aria-describedby="email" placeholder="Enter email">
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
                  <div class="form-group col-lg-6">
                    <label for="bairro">Bairro</label>
                    <input type="text" class="form-control" v-model="endereco.bairro" id="bairro" aria-describedby="bairro" placeholder="bairro">
                 </div> 

                  <div class="form-group col-lg-6">
                    <label for="email">Numero</label>
                    <input type="number" class="form-control" v-model="endereco.numero" id="numero" aria-describedby="numero" placeholder="numero">
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
        <b-button size="sm" class="float-right" variant="success" @click="show=false">Salvar</b-button>
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
            nome:"",
            email:'',
            senha:'',
            senhaConf:''
        },
        endereco:{
            cidade:'',
            rua:'',
            bairro:'',
            numero:''
        }
      
      }
    },

    methods:{
        getDados(){
            alert("==="+this.$id)
        }
    },
    mounted(){
        this.getDados()
    }
  }
</script>