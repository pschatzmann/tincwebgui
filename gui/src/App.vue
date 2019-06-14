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
            
            <v-btn v-on:click="toggleOnOff()" :disabled="!isLoggedIn" fab dark small :color="onOffInfo.color">
                <v-icon dark>{{onOffInfo.icon}}</v-icon>
            </v-btn>

            <v-menu bottom left>
                <v-btn icon slot="activator">
                    <v-icon>more_vert</v-icon>
                </v-btn>
                <v-list>
                    <v-list-tile  @click="login()" :disabled="isLoggedIn">
                        <v-list-tile-title>Login</v-list-tile-title>
                        <v-list-tile-action />
                    </v-list-tile>

                    <v-list-tile to="/setup" :disabled="!isLoggedIn">
                        <v-list-tile-title>Setup</v-list-tile-title>
                        <v-list-tile-action/>
                    </v-list-tile>

                    <v-list-tile :disabled="!isLoggedIn" @click="action('restart')">
                        <v-list-tile-title>Restart</v-list-tile-title>
                        <v-list-tile-action/>
                    </v-list-tile>
                    
                    <v-list-tile :disabled="!isLoggedIn" @click="action('generate-keys')">
                        <v-list-tile-title>(Re)generate Keys</v-list-tile-title>
                        <v-list-tile-action/>
                    </v-list-tile>

                     <v-list-tile  @click="logoff()" :disabled="!isLoggedIn">
                        <v-list-tile-title>Logoff</v-list-tile-title>
                        <v-list-tile-action />
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
    import { SecurityService } from '@/services/SecurityService'

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
                this.doOnOff(this.tincIsActive)
            },

            // update the status in tinc with tincIsActive
            doOnOff(value) {
                this.$store.dispatch("setError", null)
                var self = this
                WebServices.ping().then( result => {
                    console.log(result)
                    if (value){
                        WebServices.action('start').then( result => {
                            console.log(result); 
                            self.tincIsActive = true
                        }, err => {
                            self.$store.dispatch('setError', err)
                            self.tincIsActive = false
                            // test:  self.tincIsActive = true
                        })
                    } else {
                        self.tincIsActive = false
                        WebServices.action('stop').then( result => {
                            console.log(result); 
                        }, err => {
                            console.log(err)
                            //self.$store.dispatch('setError', err)
                            self.checkOn()
                        })
                    } 
                }, err => {
                    console.log(err)
                    self.tincIsActive = false
                    this.$store.dispatch('setError', "The Tinc Webservice is not available")
                })
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
                    self.$store.dispatch('setError', {msg: result.data, type:'success'})
                }, err => {
                    self.$store.dispatch('setError', err)
                });
            },

            login() {
                var self = this
                // show login screen
                SecurityService.signIn().then(() => {
                    console.log("login")
                }, err => {
                    self.$store.dispatch('setError', err)
                })
            },

            logoff() {
                var self = this
                SecurityService.signOut().then(()=>{
                    self.isLoggedIn = false
                    self.$router.push("/")
                }, err => {
                    self.isLoggedIn = false
                    console.log("logoff", err)
                    // we do not really care if the endpoint does not support a logoff
                    if (err && err.message != "no end session endpoint"){
                        self.$store.dispatch('setError', err)
                    }
                    self.$router.push("/")
                })
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

            isLoggedIn: {
                get() {
                    return this.$store.state.isLoggedIn
                },
                set(value) {
                    this.$store.dispatch('setLoggedIn', value)
                }  
            },

            isActive() {
                return this.$store.getters.isEnabled
            },


            onOffInfo() {
                return this.tincIsActive ?  {color:"green", icon:"toggle_on"} : {color:"red", icon:"toggle_off"}
            }
        },

        created() {
            // setup service URL
            var self = this
            WebServices.url = window.location.origin

            // manage this.isLoggedIn
            SecurityService.getMgr().then(mgr => {
                // we mignt sign on later
                mgr.events.addUserLoaded(user => { 
                    console.log("addUserLoaded", user) 
                    self.checkOn()
                }, error => console.log(error))

                // we might be signed in already
                mgr.isSignedIn().then(signedIn => {
                    self.isLoggedIn = signedIn
                    if (signedIn){
                        self.checkOn()
                    }
                })
            
                mgr.events.addUserSignedOut(() => {
                    console.log("addUserSignedOut")
                    self.isLoggedIn = false
                },error => console.log(error))
            })


            // update login flag
            SecurityService.isSignedIn().then(result => {
                self.isLoggedIn = result
                // update tinc status - we need to be logged in for this to work
                if (result){
                    this.checkOn()
                }
            })

        },

    }

</script>

