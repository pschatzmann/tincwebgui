<template>
    <v-container id='performanceContainer' fluid>
        <v-card>
            <v-container fluid>
                <area-chart :stacked="true" :refresh="5" :data="data" legend="right"  title="Traffic In (bytes)"/>
            </v-container>
            <v-container fluid>
                <area-chart :stacked="true" :refresh="5" :data="data" legend="right" title="Traffic Out (bytes)"/>
            </v-container>
        </v-card>
    </v-container>
</template>


<script>
import Vue from 'vue'
import VueChartkick from 'vue-chartkick'
import Chart from 'chart.js'
Vue.use(VueChartkick, {adapter: Chart})

export default {
    name: "performance",
    data: () => ({
        data : [
            {name: 'node1', data: {'2017-01-01 00:00:00 -0800': 3, '2017-01-02 00:00:00 -0800': 4}},
            {name: 'node2', data: {'2017-01-01 00:00:00 -0800': 5, '2017-01-02 00:00:00 -0800': 3}}
        ]

    }),
}

</script>

<style>
    #performanceContainer {
        height:80vh;
    }
</style>