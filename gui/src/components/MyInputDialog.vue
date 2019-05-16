/**
* Simple Input Dialog where we can enter a text
**/
   
<template>
    <v-dialog v-model="this.inputData.visible"  max-width="290">
      <v-card>
        <v-card-title class="headline">{{title}}</v-card-title>

        <v-card-text>
          <v-text-field v-model="inputData.inputText" :label="this.inputData.title"></v-text-field>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" flat="flat" @click="inputData.dialogVisible = false">Cancel</v-btn>
          <v-btn color="green darken-1" flat="flat" @click="ok()">OK</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
</template>

<script>
  export default {
    name: "my-input-dialog",
    
    props: {
      inputData: {
        type: Object,
        default: function () {
          return { title: 'title', inputText:'', dialogVisible:true, processOK:null }
        }
      }
    },

    data () {
      return {
      }
    },

    methods: {
      ok() {
          if (this.inputData.processOK) {
            this.inputData.processOK(this.inputData.inputText)
          }
          this.inputData.dialogVisible = false
      }
    }
  }
</script>

<style>
</style>
