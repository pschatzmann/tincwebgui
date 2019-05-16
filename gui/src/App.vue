<template>
    <v-app>
        <my-navigation-drawer/>

        <v-toolbar color="blue" dark fixed app clippedLeft>
            <v-toolbar-side-icon @click.stop="toggleDrawer()"/>
            <v-toolbar-title>Tinc VPN - Web GUI</v-toolbar-title>
            <v-spacer/>
            
            <v-flex xs12 sm2 d-flex v-if="networks.length>1">
                <v-select :value="network" :items="networks" prepend-icon="wifi"/>
            </v-flex>
            
            <v-btn v-on:click="toggleOnOff()"  fab dark small :color="onOffInfo.color">
                <v-icon dark>{{onOffInfo.icon}}</v-icon>
            </v-btn>

            <v-menu bottom left>
                <v-btn icon slot="activator">
                    <v-icon>more_vert</v-icon>
                </v-btn>
                <v-list>
                    <v-list-tile  @click="doRestart()">
                        <v-list-tile-title>Restart</v-list-tile-title>
                        <v-list-tile-action/>
                    </v-list-tile>
                    
                    <v-list-tile  @click="action('generate-keys')">
                        <v-list-tile-title>(Re)generate Keys</v-list-tile-title>
                        <v-list-tile-action/>
                    </v-list-tile>
                </v-list>
            </v-menu>

        </v-toolbar>

        <v-content>
            <router-view/>
        </v-content>

    </v-app>

</template>

<script>
    import MyNavigationDrawer from "@/components/MyNavigationDrawer";
    import WebServices from '@/services/WebServices'

    export default {
        name: "my-app",
        components: {MyNavigationDrawer},
        data: () => ({
            networks : [],
            network: null,
        }),

        methods: {
            // show / hide drawer
            toggleDrawer() {
                this.drawer = !this.drawer
            },

            // activate / deactivate tinc
            toggleOnOff() {
                this.tincIsActive = !this.tincIsActive
                this.doOnOff()
            },

            // update the status in tinc with tincIsActive
            doOnOff(){
                this.$store.dispatch("setError", null)
                var self = this
                if (this.tincIsActive){
                    WebServices.action('start').then( result => {
                         console.log(result); 
                        self.tincIsActive = true
                    }, err => {
                        self.$store.dispatch('setError', err)
                        self.tincIsActive = false
                        // test: self.tincIsActive = true
                    })
                } else {
                    WebServices.action('stop').then( result => {
                        console.log(result); 
                    }, err => {
                        console.log(err)
                        //self.$store.dispatch('setError', err)
                        self.checkOn()
                    })
                }
            },

            // restart tinc and then check if it is on
            doRestart() {
                this.$store.dispatch("setError", null)
                var self = this
                WebServices.action("restart").then( result => {
                    console.log(result); 
                    self.checkOn();
                }, err => {
                    self.$store.dispatch('setError', err)
                });
            },

            // check if tinc is active and update tincIsActive
            checkOn() {
                var self = this
                WebServices.get('pid').then( result => {
                    console.log('pid',result); 
                    var isnum = /^\d+$/.test(result.data);
                    self.tincIsActive = isnum
                }, err => {
                    console.log(err); 
                    self.tincIsActive = false
                })
            },

            // execute post action
            action(action) {
                this.$store.dispatch("setError", null)
                var self = this
                WebServices.action(action).then( result => {
                    console.log(result); 
                }, err => {
                    self.$store.dispatch('setError', err)
                });
            },

        }, 

        computed: {
            drawer: {
                get() {
                    return this.$store.state.navigationDrawer.drawer
                },
                set(value) {
                    this.$store.dispatch('setNavigationDrawer', value)
                }
            },

            tincIsActive:{
                get() {
                    return this.$store.state.isActive
                },
                set(value) {
                    this.$store.dispatch('setActive', value)
                }  
            },

            onOffInfo() {
                return this.tincIsActive ?  {color:"green", icon:"toggle_on"} : {color:"red", icon:"toggle_off"}
            }

        },

        mounted() {
            this.doOnOff()
        }
    }

</script>

