/
  Loads all tinc conmfiguration parameters from the webservice and displays them in a table
**/

<template>
    <div>
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
            <v-btn fab dark small color="red">
                <v-icon dark>add</v-icon>
            </v-btn>   
        </v-container>
    </div>
</template>

<script>
    import WebServices from "@/services/WebServices"

    export default {
        name: "currentConfiguration",
        data: () => ({
            headers: [
                { text: 'Invitation',value: 'Invitation' },
                { text: 'Name', value: 'Name' },
            ],
            items: []
        }),

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

