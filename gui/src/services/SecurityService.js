import 'babel-polyfill';
import Oidc from 'oidc-client';
import 'babel-polyfill';import WebServices from '@/services/WebServices'

var _mgr=null;
Oidc.Log.logger = console;
Oidc.Log.level = Oidc.Log.INFO;

export const SecurityService = {
    // Returns a Oidc.UserManager
    async getMgr() {
        if (_mgr==null) {
            const response = await WebServices.getAuth()
            if (response.status!=200) {
                throw new Error("The Auth Service did not return a valid status");
            }
            // check the result
            const auth =  response.data
            if (!auth){
                throw new Error("The Auth Service did not return a valid result");
            }
            if (!auth.ProviderUrl){
                throw new Error("The ProviderUrl is not defined");
            }
            if (!auth.ClientId){
                throw new Error("The ClientId is not defined");
            }
            _mgr = new Oidc.UserManager({
                // userStore: new Oidc.WebStorageStateStore({store: localStorage}),  
                authority: auth.ProviderUrl,
                client_id: auth.ClientId,
                redirect_uri: window.location.origin + '/oidc-callback',
                response_type: 'id_token',
                scope: 'openid profile email',
                post_logout_redirect_uri: window.location.origin + '/',
                silent_redirect_uri: window.location.origin + '/silent-renew.html',
                accessTokenExpiringNotificationTime: 10,
                automaticSilentRenew: true,
                filterProtocolClaims: false,
                loadUserInfo: true,
                revokeAccessTokenOnSignout: true
            })
        }        
        return _mgr
    },

    // Renew the token manually
    async renewToken () {       
      var mgr = await this.getMgr()
      var user = await mgr.signinSilent()
      if (user == null) {
            return await this.signIn(null)
      }
      return user
    },

    async signinRedirectCallback() {
        var mgr = await this.getMgr()
        return await mgr.signinRedirectCallback()
    },
  
    // Get the user who is logged in
    async getUser () {
        var mgr = await this.getMgr()
        var user = await mgr.getUser()
      if (user==null){
          return await this.signIn()
      }
      return user
    },
  
    // Check if there is any user logged in
    async isSignedIn () {
        var mgr = await this.getMgr()
        var user = await mgr.getUser()
        if (user == null) {
            user = await mgr.signinSilent()      
            return user!=null
        }
        return true
    },
  
    // Redirect of the current window to the authorization endpoint.
    async signIn () {
        var mgr = await this.getMgr()
        return await mgr.signinRedirect()
    },
    
    // Redirect of the current window to the end session endpoint
    async signOut () {    
        var mgr = await this.getMgr()
        const result = await mgr.signoutRedirect()
        // remove the user from the cookies/database
        mgr.removeUser()

        return result  
    },
  
    // Get the token id
    async getIdToken() {
        var mgr = await this.getMgr()
        var user = await mgr.getUser()
        return user!=null ? user.id_token : null
    },
  
    // Get the session state
    async getSessionState() {
        var mgr = await this.getMgr()
        var user = await mgr.getUser()
        return user!=null ? user.session_state : null
    },
  
    // Get the access token of the logged in user
    async getAccessToken(){
        var mgr = await this.getMgr()
        var user = await mgr.getUser()
        return user!=null ? user.access_token : null
    },
  
    // Takes the scopes of the logged in user
    async getScopes() {
        var mgr = await this.getMgr()
        var user = await mgr.getUser()
        return user!=null ? user.scopes : null

    },

    // Get the user roles logged in
    async getRole () {
        var mgr = await this.getMgr()
        var user = await mgr.getUser()
        return user!=null ? user.profile.role : null
    },

    async checkLogin(store) {
        if (!store.state.isLoggedIn) {
            var mgr = await this.getMgr()
            var signedIn = await this.isSignedIn()
            if (!signedIn) {
                var user = await mgr.signinSilent()
                store.dispatch('setLoggedIn', user!=null)
            } else {
                store.dispatch('setLoggedIn', true)
            }
        }
    }
    
  }
  