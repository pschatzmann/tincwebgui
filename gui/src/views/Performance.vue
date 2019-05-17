<template>
    <div>
        <v-toolbar  class="my-toolbar">
            <div>
                <label class="myCheckBoxLabel"> <input type="radio" value="Bytes" class="myStatus"
                                                       v-model="asBytes">Bytes</label>
                <label class="myCheckBoxLabel"> <input type="radio" value="Packets" class="myStatus"
                                                       v-model="asBytes">Packets</label>
            </div>
            <v-spacer/>
        </v-toolbar>
        <v-alert :value="error!=null" type="error">{{error}}</v-alert>
        <v-container id='performanceContainer' fluid>
            <v-card>
                <v-container fluid>
                    <area-chart :stacked="true" :data="chartData.rx" legend="right"  :title="'Traffic In '+unit"/>
                </v-container>
                <v-container fluid>
                    <area-chart :stacked="true" :data="chartData.tx" legend="right" :title="'Traffic Out '+unit"/>
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
        timer: null,
        asBytes: 'Bytes',
        chartData : null,
        unit: null

       // [ 
       //     {name: 'node1', data: {'2017-01-01 00:00:00 -0800': 3, '2017-01-02 00:00:00 -0800': 4}},
       //     {name: 'node2', data: {'2017-01-01 00:00:00 -0800': 5, '2017-01-02 00:00:00 -0800': 3}}
       // ]

    }),
    
    computed: {
        error() {
            return this.$store.state.error
        }
    }, 

    methods: {
        updateData() {
            // trigger reloading of data which will be diplayed with next call
            WebServices.getNetworkTraffic(this.asBytes == 'Bytes').then(result => {
                this.chartData = result.data
                this.unit = asBytes == 'Bytes' ? 'Bytes' : 'Packets'
            },error => {
                this.$store.dispatch('setError', error)
            })
        },
    },

    created() {
        this.updateData();
        this.timer = setInterval(this.updateData(), 10000)
    },

    beforeDestroy() {
        clearInterval(this.timer)
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