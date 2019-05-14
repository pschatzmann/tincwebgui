<template>
    <v-container  >
        <v-card id='swaggerCard'/>
    </v-container>
</template>

<script>
import SwaggerUI from 'swagger-ui'
import 'swagger-ui/dist/swagger-ui.css'

export default {
    name: "swagger",
    data: () => ({
    }),

    mounted() {
        SwaggerUI({
            dom_id: '#swaggerCard',
            url: '/swagger.yml',
            deepLinking: false,
        })

    }
}
</script>

<style>
    .info {
        background-color:white!important; 
     }

    #swaggerCard {
        top:-50px;
    }
</style>
