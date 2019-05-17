<template>
    <div>
        <v-toolbar  class="my-toolbar">
            <div>
                <label class="myCheckBoxLabel"> <input type="checkbox" class="myStatus"
                                                       v-model="documentStatusProcessed">Bytes</label>
                <label class="myCheckBoxLabel"> <input type="checkbox" class="myStatus"
                                                       v-model="documentStatusMarked">Packets</label>
            </div>
            <v-spacer/>
        </v-toolbar>
        <v-alert :value="error!=null" type="error">{{error}}</v-alert>
        <v-container id='performanceContainer' fluid>
            <v-card>
                <v-container fluid>
                    <area-chart :stacked="true" :refresh="5" :data="dataCalculated" legend="right"  title="Traffic In (bytes)"/>
                </v-container>
                <v-container fluid>
                    <area-chart :stacked="true" :refresh="5" :data="dataCalculated" legend="right" title="Traffic Out (bytes)"/>
                </v-container>
            </v-card>
        </v-container>
    </div>
</template>

<script>
import WebServices from "@/services/WebServices"
import Vue from 'vue'
import VueChartkick from 'vue-chartkick'
import Chart from 'chart.js'
Vue.use(VueChartkick, {adapter: Chart})

export default {
    name: "performance",
    data: () => ({
        actualData : []
       // [ 
       //     {name: 'node1', data: {'2017-01-01 00:00:00 -0800': 3, '2017-01-02 00:00:00 -0800': 4}},
       //     {name: 'node2', data: {'2017-01-01 00:00:00 -0800': 5, '2017-01-02 00:00:00 -0800': 3}}
       // ]

    }),
    
    computed: {
        error() {
            return this.$store.state.error
        },

        dataCalculated() {
            this.updateData()
            return this.actualData
        }
    }, 

    methods: {
        updateData() {
            // trigger reloading of data which will be diplayed with next call
            WebServices.getNetworkTraffic().then(result => {
                this.actualData = result.data
            },error => {
                this.$store.dispatch('setError', error)
            })
        }
    }
}

</script>

<style>
    #performanceContainer {
        height:80vh;
    }

    .myCheckBoxLabel {
        padding: 8px;
    }

    .myCheckBoxLabel > input {
        margin: 5px;
    }

</style>