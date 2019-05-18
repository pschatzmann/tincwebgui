<!--
  Display Chart with download and upload statistics
-->

<template>
    <div>
        <v-toolbar  class="my-toolbar">
            <v-switch class="my-switch" v-model="active" :label="activeLabel"></v-switch>

            <div>
                <label class="myCheckBoxLabel"> <input type="radio" value="Bytes" class="myStatus"
                                                       v-model="asBytes" v-on:change="updateData()">Bytes</label>
                <label class="myCheckBoxLabel"> <input type="radio" value="Packets" class="myStatus"
                                                       v-model="asBytes" v-on:change="updateData()">Packets</label>
            </div>
            <v-spacer/>
        </v-toolbar>
        <v-alert :value="error!=null" type="error">{{error}}</v-alert>
        <v-container id='performanceContainer' fluid>
            <v-card>
                <v-container fluid>
                    <area-chart :stacked="true" :data="chartData.rx" legend="right"  :title="'Traffic RX '+unit"/>
                </v-container>
                <v-container fluid>
                    <area-chart :stacked="true" :data="chartData.tx" legend="right" :title="'Traffic TX '+unit"/>
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
        chartData : {rx:[], tx: []},
        unit: "",
        activeData: false,

       // [ 
       //     {name: 'node1', data: {'2017-01-01 00:00:00 -0800': 3, '2017-01-02 00:00:00 -0800': 4}},
       //     {name: 'node2', data: {'2017-01-01 00:00:00 -0800': 5, '2017-01-02 00:00:00 -0800': 3}}
       // ]

    }),
    
    computed: {
        error() {
            return this.$store.state.error
        },

        activeLabel() {
            return this.active ? 'On' : 'Off'
        },

        active: {
            set(value){
                const self = this
                if (value){
                    WebServices.networkTrafficOn().then(result => {
                        console.log(result)
                        this.activeData = value
                    },error => {
                        self.$store.dispatch('setError', error)
                    })
                } else {
                    WebServices.networkTrafficOff().then(result => {
                        console.log(result)
                        this.activeData = value
                    },error => {
                        self.$store.dispatch('setError', error)
                    })
                }
            },
            get() {
                return this.activeData
            }
        }
    }, 

    methods: {
        // trigger reloading of data which will be diplayed with next call
        updateData() {
            WebServices.getNetworkTraffic(this.asBytes == 'Bytes').then(result => {
                this.chartData = result.data
                this.unit = this.asBytes == 'Bytes' ? '(Bytes)' : '(Packets)'
            },error => {
                this.$store.dispatch('setError', error)
            })
        },
        // update status with ws
        updateStatus() {
            WebServices.getNetworkTrafficStatus().then(result => {
                this.activeData = result.data.Status
            },error => {
                console.log(error)
                this.activeData = false
            })
        }
    },

    created() {
        this.updateData();
        this.updateStatus();
        this.timer = setInterval(() => this.updateData(), 5000)
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

    .my-switch {
        max-width: 100px;
        max-height: 30px;
    }

</style>