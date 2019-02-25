<template>
    <div>
    <h1>Home</h1>
     <button @click="sair()">Sair</button>
    </div>
</template>

<script>
export default {
    methods:{
        sair(){
            localStorage.removeItem("token")
             location.reload();     
        }
    }
}
</script>

<style >

</style>
