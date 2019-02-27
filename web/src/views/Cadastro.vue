<template>
 <div>
    <Header/>
     
     <b-container class="cardCad">
         <div class="row">
            <div v-if="erros.length" class="col-sm-12">
                <b-card
                    header="ERROS"
                    header-text-variant="light"
                    header-tag="header"
                    header-bg-variant="danger"
                >
                    <b-card-text>
                        <div >
                                <p v-for="(error,key) in erros" :key="key">{{error}}</p>
                        </div>
                    </b-card-text>
                </b-card>
            </div> 
            
             <div  class="col-sm-12" style="margin-top:20px">
                <b-card
                    header="Dados Pessoais"
                    header-text-variant="light"
                    header-tag="header"
                    header-bg-variant="success"
                >
                    <b-card-text>
                        <form >  
                        <div class="row">
                            <div class="form-group col-lg-6">
                                <label for="nome">Nome</label>
                                <input type="text" v-model="user.nome" class="form-control" id="nome" aria-describedby="nome" placeholder="Nome">
                            </div>  
                             <div class="form-group col-lg-6">
                                <label for="email">Email</label>
                                <input type="email"  v-model="user.email" class="form-control" id="email" aria-describedby="email" placeholder="Enter email">
                            </div>
                        </div>
                        </form>
                    </b-card-text>
                </b-card>
            </div>
            <div  class="col-sm-6" style="margin-top:20px">
                <b-card
                    header="Endereço"
                     header-text-variant="light"
                    header-tag="header"
                    header-bg-variant="success"
                >
                    <b-card-text>
                        <form >  
                        <div class="row">
                            <div class="form-group col-lg-6">
                                <label for="cidade">Cidade</label>
                                <input type="text" v-model="endereco.cidade" class="form-control" id="cidade" aria-describedby="cidade" placeholder="Cidade">
                            </div>  
                             <div class="form-group col-lg-6">
                                <label for="rua">Rua</label>
                                <input type="text" v-model="endereco.rua" class="form-control" id="rua" aria-describedby="rua" placeholder="Rua">
                            </div>
                             <div class="form-group col-lg-12">
                                <label for="bairro">Bairro</label>
                                <input type="bairro" v-model="endereco.bairro" class="form-control" id="bairro" aria-describedby="bairro" placeholder="Bairro">
                            </div>
                             
                        </div>
                        </form>
                    </b-card-text>
                </b-card>
            </div>   
            <div  class="col-sm-6"  style="margin-top:20px">
                <b-card
                    header="Senha"
                     header-text-variant="light"
                    header-tag="header"
                    header-bg-variant="success"
                >
                    <b-card-text>
                        <form >  
                        <div class="row">
                            <div class="form-group col-lg-12">
                                <label for="senha">Senha</label>
                                <input type="password" v-model="user.senha" class="form-control" id="senha" aria-describedby="senha" placeholder="senha">
                            </div>  
                             <div class="form-group col-lg-12">
                                <label for="senhaConfirm">Confirme a Senha</label>
                                <input type="password" v-model="user.senhaConf"  class="form-control" id="senhaConfirm" aria-describedby="senhaConfirm" >
                            </div>
                        </div>
                        </form>
                    </b-card-text>
                </b-card>
            </div>        
         </div>
         <button id="btn" @click="isCadastro" class="btn btn-dark">CADASTRAR</button>
     </b-container>    
 </div>
 
</template>

<script>
import Header from '@/views/HeaderCadastro'
export default {
    components:{
        Header
    },
     data(){
    return {
      user:{
        nome:'',
        email:'',
        senha:'',
        senhaConf:''
      },
      endereco:{
         cidade:'',
         rua:'',
         bairro:''
      },
      erros:[]
    }
  },
  methods:{
      async isCadastro(e){

if(this.$data.user.nome && this.$data.user.email && this.$data.endereco.cidade && this.$data.endereco.rua && this.$data.endereco.bairro  && this.$data.user.senha && this.$data.user.senhaConf){
                var form = new FormData()

                form.append("nome",this.$data.user.nome)
                form.append("email",this.$data.user.email)
                form.append("cidade",this.$data.endereco.cidade)
                form.append("rua",this.$data.endereco.rua)
                form.append("bairro",this.$data.endereco.bairro)
             
                if(this.$data.user.senha == this.$data.user.senhaConf){
                    form.append("password",this.$data.user.senha)
                }else{
                    this.$data.erros.push("senhas nao sao iguais")
                }

                const response = await fetch("http://localhost:9000/cadastro/",{
                    method:"POST",
                    body: form
                })

                const res = await response.json()

                if(res == 'Ja existe este Usuario'){
                    this.$swal.fire({
                        title: res,
                        type: 'warning',
                        confirmButtonText: 'Faça o Login'
                    })
                }else{
                    this.$swal.fire({
                        title: 'Parabéns',
                        text: res,
                        type: 'success',
                        confirmButtonText:"<a style='color:#ffffff' href='#/login'>Faça o login</a>"
                    })

                    setTimeout(function() {
                        window.location.href = "#/login"
                    }, 1000);
                }
                  
          }

           
          if(!this.$data.user.nome){
              this.$data.erros.push("O nome é Obrigatorio")
          }
          
          if(!this.$data.user.email){
               this.$data.erros.push("O Email é Obrigatorio")
          }
          
          if(!this.$data.endereco.cidade){
               this.$data.erros.push("A Cidade é Obrigatorio")
          }
          
          if(!this.$data.endereco.rua){
               this.$data.erros.push("A Rua é Obrigatorio")
          } 
          
          if(!this.$data.endereco.bairro){
               this.$data.erros.push("O Bairro é Obrigatorio")
          } 
          
          if(!this.$data.endereco.bairro){
               this.$data.erros.push("O Email é Obrigatorio")
          } 
          
        
          
          if(!this.$data.user.senha){
               this.$data.erros.push("A Senha é Obrigatorio")
          } 
          
          if(!this.$data.user.senhaConf){
               this.$data.erros.push("A Senha de Confirmação é Obrigatorio")
          } 

        
      }
  }
}
</script>

<style scoped>
    .cardCad{
        margin-top:20px;
        background-color:#DCDCDC;
        border-radius: 5px;
        padding:20px
    }
    #btn{
        
        margin-top:20px;
    }
</style>
