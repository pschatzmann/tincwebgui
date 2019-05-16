/
  Loads all tinc conmfiguration parameters from the webservice and displays them in a table
**/

<template>
    <div>
        <v-toolbar   class="my-toolbar">
            <v-btn v-on:click="invite()" >
                <v-icon>call_made</v-icon>
            </v-btn>   
            <v-btn v-on:click="joinInvite()"  >
                <v-icon>call_received</v-icon>
            </v-btn>   
            <v-spacer/>
        </v-toolbar>
        <v-alert :value="error!=null" type="error">{{error}}</v-alert>
        <v-container fluid>
            <MyInputDialog :inputData="inputDialogData"/>
            <v-card>
                    <v-data-table
                        :headers="headers"
                        :items="items"
                        hide-actions=true

                        class="elevation-1">
                        <template v-slot:items="props">
                        <td>{{ props.item.Invitation }}</td>
                        <td>{{ props.item.Name }}</td>
                        </template>
                    </v-data-table>
            </v-card>
        </v-container>
    </div>
</template>

<script>
    import WebServices from "@/services/WebServices"
    import "@/components/MyInputDialog"

    export default {
        name: "currentConfiguration",
        data: () => ({
            headers: [
                { text: 'Invitation',value: 'Invitation' },
                { text: 'Name', value: 'Name' },
            ],
            items: [],
            inputDialogData: {
                visible: false
            }
        }),

        methods: {
            invite() {
                this.$store.dispatch("setError", null)
                var self = this
                this.inputDialogData.title = 'Node Name'
                this.inputDialogData.ok = (nodeName) => {
                    WebServices.invite(nodeName).then( result => {
                        console.log(result); 
                    }, err => {
                        self.$store.dispatch('setError', err)
                    });
                }
                this.inputDialogData.visible = true
            }, 

            joinInvite() {
                this.$store.dispatch("setError", null)
                var self = this
                this.inputDialogData.title = 'Node Name'
                this.inputDialogData.ok = (invitation) => {
                    WebServices.join(invitation).then( result => {
                        console.log(result); 
                    }, err => {
                        self.$store.dispatch('setError', err)
                    });
                }
                this.inputDialogData.visible = true
            }
        },

        computed: {
            error() {
                return this.$store.state.error
            },
        },

        mounted() {
            WebServices.getInvitations().then(result => {
                this.items = result.data
            },error => {
                this.$store.dispatch('setError', error)
            })
        }
    }
</script>

