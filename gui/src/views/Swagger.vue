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
        token: null
    }),

    created() {
        SecurityService.getAcessToken().then(accessToken => {
            this.token = accessToken
        })
    },

    mounted() {
        SwaggerUI({
            dom_id: '#swaggerCard',
            url: '/swagger.yml',
            deepLinking: false,
            requestInterceptor: (request) => {
                console.log('Swagger intercept request -> add Bearer token', this.token);
                if (this.token!=null) {
                    request.headers.Authorization = "Bearer "+this.token;
                }
                return request;
            },
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
