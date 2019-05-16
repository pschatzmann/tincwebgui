/
  Loads all tinc conmfiguration parameters from the webservice and displays them in a table
**/

<template>
    <div>
        <v-toolbar   class="my-toolbar">
            <v-btn v-on:click="invite()" >
                <v-icon>cload_upload</v-icon>
            </v-btn>   
            <v-btn v-on:click="joinInvite()"  >
                <v-icon>cload_download</v-icon>
            </v-btn>   
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
                    WebServices.join(invitation).then( result => {
                        console.log(result); 
                        self.updateInvitations()
                    }, err => {
                        self.$store.dispatch('setError', err)
                    });
                }
                this.showInputDialog('Join Invitation','Enter the generated Invitation:','Invitation', ok )
            }
        },

        showInputDialog(title, text, inputName, ok){
            this.$store.dispatch("setError", null)
            var self = this
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

