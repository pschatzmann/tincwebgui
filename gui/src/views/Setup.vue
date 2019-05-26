<!--
   Setup new node
-->

<template>
    <div >
        <v-alert :value="error.msg!=null" :type="error.type">
            <span v-html="error.msg"></span>
        </v-alert>
        <v-container grid-list-md >
            <v-stepper v-model="stepperModel">
                <v-stepper-header>
                    <v-stepper-step :complete="stepperModel > 1" step="1">Select Type of Creation</v-stepper-step>
                    <v-divider></v-divider>
                    <v-stepper-step :complete="stepperModel > 2" step="2">Create Node</v-stepper-step>
                    <v-divider></v-divider>
                </v-stepper-header>

                <v-stepper-items>
                    <v-stepper-content class="myContent"  step="1">
                        <v-container
                        class="mb-5">
                            <v-flex xs12 >
                                <div class="myRadio"><input type="radio" value="standAlone" v-model="mode"> New (Initial) Node </div>
                                <div class="myRadio"><input type="radio" value="web" v-model="mode"> New Node with existing Web GUI</div>
                                <div class="myRadio"><input type="radio" value="invitation" v-model="mode"> New Node with Tinc Invitation</div>
                            </v-flex>
                        </v-container>
                        <div class="mySpace"/>

                        <v-btn flat to="/">Cancel</v-btn>
                        <v-btn color="primary" @click="stepperModel = 2">Continue</v-btn>
                    </v-stepper-content>

                    <v-stepper-content step="2">
                        <v-text-field v-model="setup.nodeName" label="Node Name"></v-text-field>
                        <v-text-field v-model="setup.subnet" label="Subnet"></v-text-field>
                        <v-text-field v-model="setup.localIP" label="Local IP"></v-text-field>

                        <v-text-field v-model="setup.invitation" v-if="mode=='invitation'" label="Invitation"></v-text-field>
                        <v-text-field v-model="setup.connectTo" v-if="mode=='web'" label="ConnectTo"></v-text-field>
                        <div/>
                        <v-btn flat @click="stepperModel = 1">Back</v-btn>
                        <v-btn flat to="/">Cancel</v-btn>
                        <v-btn color="primary" @click="setupNode()" >Create</v-btn>
                    </v-stepper-content>
                </v-stepper-items>
            </v-stepper>
        </v-container>
    </div>
</template>

<script>
import WebServices from "@/services/WebServices"


  export default {
    data () {
        return {
            stepperModel: 0,
            mode: 'web',
            setup : {
                nodeName: '',
                subnet: '',
                localIP: '',
                invitation: '',
                connectTo: ''
            }
        }
    },

    methods: {
        setupData() {
            var self = this
            WebServices.getParameter("NodeName").then(v => {
                self.setup.nodeName = v.data
            }, () => {
                WebServices.getParameter("Name").then(v => {
                    self.setup.nodeName = v.data
                })
            })

            WebServices.getParameter("Subnet").then(v => {
                self.setup.subnet = v.data
            })

            WebServices.getParameter("VpnIP").then(v => {
                self.setup.localIP = v.data
            })

            WebServices.getParameter("ConnectTo").then(v => {
                self.setup.connectTo = v.data
            })
        },

        setupNode() {
            var self = this
            WebServices.setup(this.setup).then(result => {
                self.$store.dispatch('setError', { mode:'success', msg: result.data })
                self.$router.push('/network') 
            }, err => {
                self.$store.dispatch('setError', err)
            })
        },

    },

    computed: {
        error() {
            return this.$store.state.error
        },
    },

    created() {
        this.setupData() 
    },
  }
</script>

<style>
    #myContainer {
        height:80vh;
    }
    .mySpace {
        height: 5px
    }

    .myRadio {
        font-size: 16px;
        font-family: 'Roboto', sans-serif;
        padding: 12px
    }

</style>