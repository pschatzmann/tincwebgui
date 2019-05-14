<!--
---
--- View: About - Provide Application Help. The content is loaded from Wordpress
---
-->

<template>
  <div>
    <v-alert :value="error!=null" type="error">{{error}}</v-alert>
    <v-container fluid>
        <v-card>
            <v-container fluid>
              <h2>Tinc Web GUI</h2>
              <h3>Introduction</h3>
              <p>The Tinc Web GUI provides a user interface for the Tinc VPN. We display the current parameter
                settings, an overview of the actually defined network and the performance information of the current network traffic.</p> 
                <p>In the actions menu you can find the most important functions. You can also create new invitations and display the actually existing open invitations.</p>
                <p>The Tinc functionaliy is also made available via  <a href="/swagger">web-services</a> </p>
              <h3>Environement Variables</h3>
              <p>All tinc parameters can also be defined with the help of environment variables. The Tinc Web GUI supports the following
              additional environment variables to manage authrizations: </p>
              <p>
              <ul>
                <li>SERVICE_ACTIVE: true or false</li>
                <li>PASSWORD_ACTIVE: true or false</li>
                <li>PASSWORD</li>
              </ul>
              </p>
            </v-container>
        </v-card>
    </v-container>
  </div>
</template>

<script>
export default {
    name: "about",
    data: () => ({
    }),
    computed: {
        error() {
            return this.$store.state.error
        },
    }
}
</script>
