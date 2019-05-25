<template>
    <div id="sticky">
        <v-toolbar   class="my-toolbar">

            <v-btn icon v-on:click="$emit('back')">
                <v-icon>navigate_before</v-icon>
            </v-btn>

            <v-btn icon :disabled="endOfPage" v-on:click="$emit('next')">
                <v-icon>navigate_next</v-icon>
            </v-btn>

            <v-spacer/>


        </v-toolbar>
        <v-alert :value="error.msg!=null" :type="error.type">
            <span v-html="error.msg"></span>
        </v-alert>
    </div>
</template>


<script>
    export default {
        name: "my-document-toolbar",

        props: ['endOfPage'],

        data: () => ({
            toolbarElement: null,
            top: 0,
        }),

        computed: {
            pageCount: {
                get() {
                    return this.$store.state.actualDocumentNumberOfPages
                }
            },
            
            error: {
                get() {
                    return this.$store.state.error
                }
            },
        },
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

    .myCheckBoxLabel {
        padding: 8px;
    }

    .myCheckBoxLabel > input {
        margin: 5px;
    }

</style>