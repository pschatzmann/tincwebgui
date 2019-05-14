<template>
    <v-container fluid>
        <v-alert :value="error!=null" type="error">{{error}}</v-alert>

        <v-card>
                <v-data-table
                    :headers="headers"
                    :items="items"
                    hide-actions=true

                    class="elevation-1">
                    <template v-slot:items="props">
                    <td>{{ props.item.Name }}</td>
                    <td>{{ props.item.Value }}</td>
                    </template>
                </v-data-table>
        </v-card>
    </v-container>
</template>


<script>
import WebServices from "@/services/WebServices"

export default {
    name: "currentConfiguration",

    data: () => ({
        headers: [
          { text: 'Name',value: 'Name' },
          { text: 'Value', value: 'Value' },
        ],
        items: [],
    }),

    computed: {
        error() {
            return this.$store.state.error
        },
    },

    mounted() {
        WebServices.getParameters().then(result => {
            this.items = result.data
        },error => {
           this.$store.dispatch('setError', error)
        })
    }
}

</script>

