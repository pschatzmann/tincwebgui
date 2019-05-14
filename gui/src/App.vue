<template>
    <v-app>
        <my-navigation-drawer/>

        <v-toolbar color="blue" dark fixed app clippedLeft>
            <v-toolbar-side-icon @click.stop="toggleDrawer()"/>
            <v-toolbar-title>Tinc VPN Web GUI</v-toolbar-title>
            <v-spacer/>
            
            <v-flex xs12 sm2 d-flex v-if="networks.length>1">
                <v-select :value="network" :items="networks" prepend-icon="wifi"/>
            </v-flex>
            
            <v-btn fab dark small color="green">
                <v-icon dark>toggle_on</v-icon>
            </v-btn>

            <v-menu bottom left>
                <v-btn icon slot="activator">
                    <v-icon>more_vert</v-icon>
                </v-btn>
                <v-list>
                    <v-list-tile  @click="action('restart')">
                        <v-list-tile-title>Restart</v-list-tile-title>
                        <v-list-tile-action/>
                    </v-list-tile>
                    
                    <v-divider inset></v-divider>

                    <v-list-tile  @click="action('export')">
                        <v-list-tile-title>Export</v-list-tile-title>
                        <v-list-tile-action/>
                    </v-list-tile>
                    <v-list-tile  @click="action('export-all')()">
                        <v-list-tile-title>Export All</v-list-tile-title>
                        <v-list-tile-action/>
                    </v-list-tile>
                    <v-list-tile  @click="doImport()">
                        <v-list-tile-title>Import</v-list-tile-title>
                        <v-list-tile-action/>
                    </v-list-tile>

                    <v-divider inset></v-divider>

                    <v-list-tile  @click="invite()">
                        <v-list-tile-title>Invite</v-list-tile-title>
                        <v-list-tile-action/>
                    </v-list-tile>
                    <v-list-tile  @click="joinInvite()">
                        <v-list-tile-title>Join Invitation</v-list-tile-title>
                        <v-list-tile-action/>
                    </v-list-tile>
                    
                    <v-divider inset></v-divider>
                    <v-list-tile  @click="action('purge')">
                        <v-list-tile-title>Purge unreachable Nodes</v-list-tile-title>
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
            buttons: { 
                buttonOn: {
                    on: true,
                    color: 'green',
                    icon: 'toggle_on'
                },
                buttonOff: {
                    on: false,
                    color: 'red',
                    icon: 'toggle_off'
                }
            },
            tincIsActive: null
        }),

        methods: {
            toggleDrawer() {
                this.drawer = !this.drawer
            },

            onChangeEventHandler(){
            },

            action(action) {
                WebServices.action(action).then(function(result) {
                    console.log(result); 
                }, function(err) {
                    this.$store.dispatch('setError', err)
                });
            }

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
        },

        mounted() {
            WebServices.getNetworks().then(result => {
                this.networks = result
                // select first network
                if (this.networks.length>=1){
                    this.network = this.networks[0]
                    // start tinc
                    WebServices.action('start').then( result => {
                        tincIsActive = true
                    }, err => {
                        this.$store.dispatch('setError', err)
                        tincIsActive = false
                    })
                }
            }, err => {
                this.$store.dispatch('setError', err)
                tincIsActive = false
            });
        }
    }

</script>

