/**
* Simple Input Dialog where we can enter a text
**/
   
<template>
    <v-dialog v-model="inputData.visible"  max-width="350">
      <v-card>
        <v-card-title class="headline">{{inputData.title}}</v-card-title>

        <v-card-text>{{inputData.text}}
          <v-text-field v-model="inputData.inputText" :label="inputData.inputName"></v-text-field>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" flat="flat" @click="inputData.visible = false">Cancel</v-btn>
          <v-btn :disabled="isEmpty()" color="green darken-1" flat="flat" @click="ok()">OK</v-btn>
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
          return { title: 'title',text: 'text',inputName:'inputName', inputText:'', visible:true, processOK:null }
        }
      }
    },

    data () {
      return {
      }
    },

    methods: {
      ok() {
          if (typeof this.inputData.processOK === "function") {
            this.inputData.processOK(this.inputData.inputText)
          }
          this.inputData.visible = false
      },
      isEmpty() {
        return this.inputData.inputText == ''
      }
    }
  }
</script>

<style>
</style>
