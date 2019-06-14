<!--
---
--- Navigation Drawer for the Application
---
-->
<template>
    <v-navigation-drawer
            :clipped="drawerClipped"
            :fixed="drawerFixed"
            :permanent="false"
            enable-resize-watcher
            :mini-variant="drawerMini"
            v-model="drawerModel"
            app>

        <v-list>
            <!-- Minimize Navigator -->
            <v-list-tile @click="toggleMiniDrawer">
                <v-list-tile-action>
                   <v-icon v-if="drawerMini">mdi-chevron-double-right</v-icon>  
                   <v-icon v-if="!drawerMini">mdi-chevron-double-left</v-icon>  
                </v-list-tile-action>
            </v-list-tile>

            <v-divider/>

            <!-- About -->
            <v-list-tile to="/">
                <v-list-tile-action>
                    <v-icon>info</v-icon>
                </v-list-tile-action>
                <v-list-tile-content>
                    <v-list-tile-title>About</v-list-tile-title>
                </v-list-tile-content>
            </v-list-tile>

            <!-- Properties -->
            <v-list-tile v-if="isActive" to="/parameters">
                <v-list-tile-action>
                    <v-icon>settings</v-icon>
                </v-list-tile-action>
                <v-list-tile-content>
                    <v-list-tile-title>Parameters</v-list-tile-title>
                </v-list-tile-content>
            </v-list-tile>

            <!-- Invitations -->
            <v-list-tile v-if="isActive" to="/invitations">
                <v-list-tile-action>
                    <v-icon>list</v-icon>
                </v-list-tile-action>
                <v-list-tile-content>
                    <v-list-tile-title>Invitations</v-list-tile-title>
                </v-list-tile-content>
            </v-list-tile>

            <!-- Network -->
            <v-list-tile v-if="isActive"  to="/network">
                <v-list-tile-action>
                    <v-icon>network_check</v-icon>
                </v-list-tile-action>
                <v-list-tile-content>
                    <v-list-tile-title>Network</v-list-tile-title>
                </v-list-tile-content>
            </v-list-tile>


            <!-- Perfrmance -->
            <v-list-tile v-if="isActive"  to="/performance">
                <v-list-tile-action>
                    <v-icon>bar_chart</v-icon>
                </v-list-tile-action>
                <v-list-tile-content>
                    <v-list-tile-title>Performance</v-list-tile-title>
                </v-list-tile-content>
            </v-list-tile>
            
             <!-- Swagger -->
            <v-list-tile to="/swagger">
                <v-list-tile-action>
                    <v-icon>view_list</v-icon>
                </v-list-tile-action>
                <v-list-tile-content>
                    <v-list-tile-title>Tinc Service API</v-list-tile-title>
                </v-list-tile-content>
            </v-list-tile>

        </v-list>
    </v-navigation-drawer>
</template>

<script>

    export default {
        name: "my-navigation-drawer",

        data: () => ({
            drawerClipped: true,
            drawerFixed: false,
            drawerMini: false,
        }),

        computed: {
            drawerModel: {
                get() {
                    return this.$store.state.navigationDrawer.drawer
                },

                set(value) {
                    this.$store.dispatch('setNavigationDrawer', value)
                }
            },
            isActive() {
                return this.$store.getters.isEnabled
            }
        },

        methods: {
            // toggles the drawer variant (mini/full)
            toggleMiniDrawer() {
                this.drawerMini = !this.drawerMini
                if (this.drawerMini) {
                    this.filterActive = false
                    this.masterDataActive = false
                    this.historyActive = false
                }
            }
        }
    }
</script>

<style scoped>
    .my-navigator >>> label {
        font-size: 12px;
        left: 100px;
        color: black;
        margin-top: 0px;
        padding-top: 0px;
        padding-bottom: 0px;
    }

    .my-navigator >>> .v-input--radio-group__input {
        margin-top: 0px;
        margin-bottom: 0px;
        padding-left: 100px;
        padding-top: 0px;
        padding-bottom: 0px;
    }

    .my-navigator >>> .v-list__group__items {
        margin-top: 0px;
        margin-bottom: 0px;
        padding-left: 100px;
        padding-top: 0px;
        padding-bottom: 0px;
    }

    .v-input.my-navigator {
        margin-top: 0px;
        padding-top: 0px;
        padding-bottom: 0px;
    }

    .my-navigator-lg1 >>> .v-icon {
        position: relative;
        left: -20px;
    }

    .my-navigator-cb {
        margin-top: 0px;
        padding-top: 0px;
        padding-bottom: 0px;
        margin-bottom: 0px;
        height: 30px;
    }

    .my-spacer {
        min-width:10%
    }
    .my-badge {
        font-size: 12px;
        font-family: "Roboto", "Helvetica", "Arial", sans-serif;
        font-weight: 600;
        line-height: 23px;
        background: darkblue;
        color: white;
        height: 30px;
        border-radius: 100%;
        padding-left: 4px;
        padding-right: 4px;
    }

    .my-icon {
        padding-left: 40px;
    }

</style>