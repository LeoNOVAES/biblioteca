<template>
    <div>
        <Header/>
        <div class="container-fluid bg-success banner" >
            
        </div>  
    </div>
</template>

<script>
import Header from "@/views/Header"

export default {
    components:{Header}
   
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
