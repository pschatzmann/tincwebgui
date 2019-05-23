import Oidc from 'oidc-client';
import 'babel-polyfill';

export var Mgr = new Oidc.UserManager({
  userStore: new Oidc.WebStorageStateStore({store: localStorage}),  
  authority: 'https://accounts.google.com',
  client_id: '284017214914-apb6s5blmmcc49e6prbqhqj497bvghfh.apps.googleusercontent.com',
  redirect_uri: window.location.origin + '/oidc-callback',
  response_type: 'id_token token',
  scope: 'openid profile',
  post_logout_redirect_uri: window.location.origin + '/',
  silent_redirect_uri: window.location.origin + '/silent-renew.html',
  accessTokenExpiringNotificationTime: 10,
  automaticSilentRenew: true,
  filterProtocolClaims: true,
  loadUserInfo: true,
  revokeAccessTokenOnSignout: true
})

Oidc.Log.logger = console;
Oidc.Log.level = Oidc.Log.INFO;

Mgr.events.addUserLoaded(function (user) {  
  console.log('New User Loaded：', user);
  console.log('Acess_token: ', user.access_token)
});

Mgr.events.addAccessTokenExpiring(function () {
  console.log('AccessToken Expiring：', arguments);
});

Mgr.events.addAccessTokenExpired(function () {
  console.log('AccessToken Expired：', arguments);  
  Mgr.signoutRedirect().then(function (resp) {
    console.log('signed out', resp);
  }).catch(function (err) {
    console.log(err)
  })
});

Mgr.events.addSilentRenewError(function () {
  console.error('Silent Renew Error：', arguments);
});

Mgr.events.addUserSignedOut(function () {
  console.log('UserSignedOut：', arguments);  
  Mgr.signoutRedirect().then(function (resp) {
    console.log('signed out', resp);
  }).catch(function (err) {
    console.log(err)
  })
});
