<!--
   Manage Invitations
-->

<template>
    <div>
        <v-toolbar   class="my-toolbar">
            <v-tooltip bottom>
                <span>Generate an Invitation for a Node</span>
                <template v-slot:activator="{ on }">
                    <v-btn v-on="on" v-on:click="invite()" >
                        <v-icon>cloud_upload</v-icon>
                    </v-btn>   
                </template>
            </v-tooltip>

            <v-tooltip bottom>
                <span>Join a VPN using an Invitation</span>
                <template v-slot:activator="{ on }">
                    <v-btn v-on="on" v-on:click="joinInvite()" >
                       <v-icon>cloud_download</v-icon>
                    </v-btn>   
                </template>
            </v-tooltip>

            <v-spacer/>
        </v-toolbar>
        <v-alert :value="error!=null" type="error">{{error}}</v-alert>
        <v-container fluid>
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
            <MyInputDialog :inputData="inputDialogData"/>
        </v-container>
    </div>
</template>

<script>
    import WebServices from "@/services/WebServices"
    import MyInputDialog from "@/components/MyInputDialog"

    export default {
        name: "Invitations",

        components: { MyInputDialog },

        data: () => ({
            headers: [
                { text: 'Invitation',value: 'Invitation' },
                { text: 'Name', value: 'Name' },
            ],
            items: [],
            inputDialogData: {
                 title: 'Title',
                 text: 'Text',
                 inputName: 'inputName', 
                 inputText: '', 
                 visible: false, 
                 processOK: null 
            }
        }),

        methods: {
            invite() {
                var self = this
                var ok = (nodeName) => {
                    WebServices.invite(nodeName).then( result => {
                        console.log(result); 
                        self.updateInvitations()
                    }, err => {
                        self.$store.dispatch('setError', err)
                    });
                }
                this.showInputDialog('Create Invitation','Enter the Node Name for the Invitation:','Node Name', ok )
            }, 

            joinInvite() {
                var self = this
                var ok = (invitation) => {
                    WebServices.joinInvitation(invitation).then( result => {
                        console.log(result); 
                        self.updateInvitations()
                    }, err => {
                        self.$store.dispatch('setError', err)
                    });
                }
                this.showInputDialog('Join Invitation','Enter the generated Invitation:','Invitation', ok )
            },

            showInputDialog(title, text, inputName, ok){
                this.$store.dispatch("setError", null)
                this.inputDialogData.title = title
                this.inputDialogData.text = text
                this.inputDialogData.inputName = inputName
                this.inputDialogData.inputText = ''
                this.inputDialogData.processOK = ok
                this.inputDialogData.visible = true
            },

            updateInvitations() {
                WebServices.getInvitations().then(result => {
                    this.items = result.data
                },error => {
                    this.$store.dispatch('setError', error)
                })
            },
        },

        computed: {
            error() {
                return this.$store.state.error
            },
        },

        mounted() {
            this.updateInvitations()
        }
    }
</script>

