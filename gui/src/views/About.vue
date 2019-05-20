<!--
 View: About - Provide Application Help. The content is loaded from Wordpress
-->

<template>
  <div>
    <v-alert :value="error.msg!=null" :type="error.type">{{error.msg}}</v-alert>
    <v-container fluid>
        <v-card>
            <v-container fluid>
              <h2>Tinc Web GUI</h2>
              <h3>Introduction</h3>
              <p>The Tinc Web GUI provides a user interface for the Tinc VPN.</p> <p>The primary goal of this
                functionality was to setup new tinc instances with just a few clicks.</p>
                In the actions menu you can find the most important functions. You can also create new invitations and display the actually existing open invitations.
                We can also display the current parameter.
                settings, an overview of the actually defined network and the performance information of the current network traffic.
                Finally, the whole Tinc functionaliy is also made available via  <a href="/swagger">web-services</a> 
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

<style>
  h1, h2, h3 {
      padding-top: 5px;
      padding-bottom: 5px
  }

</style>