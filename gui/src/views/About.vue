<!--
 View: About - Provide Application Help. The content is loaded from Wordpress
-->

<template>
  <div>
    <v-alert :value="error.msg!=null" :type="error.type">
          <span v-html="error.msg"></span>
    </v-alert>
    <v-container fluid>
        <v-card>
            <v-container fluid>
              <h2>Tinc Web GUI</h2>
              <h3>Introduction</h3>
              <p>The Tinc Web GUI provides a user interface for the <a href="https://www.tinc-vpn.org/">Tinc VPN</a> .</p> <p>The primary goal of this
                functionality is to help to setup a new tinc instance with just a few clicks.</p>
                In the actions menu you can find the most important general functions. You can also create new invitations and display the actual existing open invitations.
                We can also display the current parameter.
                settings, an overview of the actually defined network and the performance information of the current network traffic.
                Finally, the whole Tinc functionaliy is also made available via  <a href="/swagger">web-services</a> 
              <h3>Tinc Version</h3>
              This application is using the tinc command line untility which is available since tinc 1.1. Therfore it will not work with any Tinc 1.0.X Releases
              <h3>Installation</h3>
              The application is provided as Docker image and can be easily started with a docker-compose file but you can 
              build the project in your environment. It relies on the following dependencies
               <ul>
                <li>Go</li>
                <li>Vue with npm</li>
                <li>make</li>
              </ul>             
              The make command will build the tincwebgui executable in the dist directory. 
              <h3>Environement Variables</h3>
              <p>All tinc parameters can also be defined with the help of environment variables. The Tinc Web GUI supports the following
              additional environment variables to manage authrizations: </p>
              <p>
              <ul>
                <li>SERVICE_ACTIVE: true or false</li>
                <li>PASSWORD_ACTIVE: true or false</li>
                <li>PROVIDER_URL</li>
                <li>CLIENT_ID</li>
                <li>CLIENT_SECRET</li>
                <li>CLAIMS</li>
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
    },
  }
</script>

<style>
  h1, h2, h3 {
      padding-top: 5px;
      padding-bottom: 5px
  }

</style>