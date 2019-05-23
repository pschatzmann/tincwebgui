import 'babel-polyfill';
import { Mgr } from '@/services/SecurityMgr'

export const SecurityService = {

    // Renew the token manually
    async renewToken () {        
      var user = await Mgr.signinSilent()
      if (user == null) {
            return await this.signIn(null)
      }
      return user
    },

    async signinRedirectCallback() {
        return await Mgr.signinRedirectCallback()
    },
  
    // Get the user who is logged in
    async getUser () {
      var user = await Mgr.getUser()
      if (user==null){
          return await this.signIn()
      }
      return user
    },
  
    // Check if there is any user logged in
    async isSignedIn () {
        var user = await Mgr.getUser()
        if (user == null) {
            user = await Mgr.signinSilent()      
            return user!=null
        }
        return true
    },
  
    // Redirect of the current window to the authorization endpoint.
    async signIn () {
        return await Mgr.signinRedirect()
    },
    
    // Redirect of the current window to the end session endpoint
    async signOut () {    
        return await Mgr.signoutRedirect()
    },
  
    // Get the token id
    async getIdToken() {
        var user = await Mgr.getUser()
        return user!=null ? user.id_token : null
    },
  
    // Get the session state
    async getSessionState() {
        var user = await Mgr.getUser()
        return user!=null ? user.session_state : null
    },
  
    // Get the access token of the logged in user
    async getAcessToken(){
        var user = await Mgr.getUser()
        return user!=null ? user.access_token : null
    },
  
    // Takes the scopes of the logged in user
    async getScopes() {
        var user = await Mgr.getUser()
        return user!=null ? user.scopes : null

    },

    // Get the user roles logged in
    async getRole () {
        var user = await Mgr.getUser()
        return user!=null ? user.profile.role : null
    },

    async checkLogin(store) {
        if (!store.state.isLoggedIn) {
            var signedIn = await this.isSignedIn()
            if (!signedIn) {
                var user = await this.Mgr.signinSilent()
                store.dispatch('setLoggedIn', user!=null)
            } else {
                store.dispatch('setLoggedIn', true)
            }
        }
    }
    
  }
  