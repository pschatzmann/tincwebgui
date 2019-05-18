<!--
  Display the tinc PDF Manual
 -->

<template>
    <div>
        <my-document-toolbar  :endOfPage="pageCount==currentPage" v-on:next="next()" v-on:back="back()"
                             v-on:print="print()"/>

        <v-container>
            <div v-if="pageCount">{{currentPage}} / {{pageCount}}</div>

            <pdf id="myPdfComponent" :src='pdfDocument'
                 @num-pages="pageCount = $event"
                 @loaded="loadingCompled()"
                 @error="setError"
                 :page="currentPage"/>

        </v-container>

    </div>
</template>

<script>
import MyDocumentToolbar from '@/components/MyDocumentToolbar'
import pdf from 'vue-pdf'

export default {
    name: 'manual',

    components: {
            pdf, MyDocumentToolbar
    },

    data: () => ({
        currentPage: 1,
        pageCount: null,
        pdfDocument: "/tinc.pdf" 
    }),

    methods: {

        next() {
            console.log("next")
            if (this.currentPage < this.pageCount) {
                this.currentPage++;
            }
        },

        back() {
            if (this.currentPage > 1) {
                this.currentPage--;
            } else {
                this.$router.go(-1)
            }
        },

        setError(error) {
            this.$store.dispatch("setProcessing", false)
            this.$store.dispatch("setError", error)
        },

        loadingCompled() {
            this.$store.dispatch("setProcessing", false)
        },

    },

    computed: {
        error() {
            return this.$store.state.error
        },
    },

    created() {
        this.$store.dispatch("setProcessing", true)
    }
}
    
</script>