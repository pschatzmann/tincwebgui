<!--
  Display Parameters
 -->

<template>
    <div id='#sticky' >
        <v-toolbar  class="my-toolbar">
            <v-btn  v-on:click="toggleExpanded()">
                <v-icon dark>{{iconName}}</v-icon>
            </v-btn>  
            <v-spacer/>
        </v-toolbar>
        <v-alert :value="error.msg!=null" :type="error.type">
            <span v-html="error.msg"></span>
        </v-alert>
        <v-container fluid>
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
    </div>
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
        allItems: [],
        expanded: false
    }),

    methods: {
        toggleExpanded(){
            this.expanded = !this.expanded
            this.setItems()
        }, 

        setItems() {
            if (this.expanded) {
                this.items = this.allItems
            } else {
                this.items = this.allItems.filter(r => r.Value!="" )
            }
        }
    },

    computed: {
        error() {
            return this.$store.state.error
        },

        iconName() {
            return this.expanded ? "unfold_less" : "unfold_more"
        }
    },

    mounted() {
        WebServices.getParameters().then(result => {
            this.allItems = result.data
            this.setItems()
        },error => {
           this.$store.dispatch('setError', error)
        })
    }
}

</script>

<style scoped>
    #sticky {
        top: 63px;
        height: 63px;
        z-index: 1;
        position: -webkit-sticky;
        position: sticky;
        visibility: visible;
    }

   .myToolbar > .v-toolbar__content {
        height: 63px;
    }
</style>

