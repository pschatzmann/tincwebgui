import axios from 'axios'
import { Mgr } from '@/services/SecurityMgr'

/**
 * Access Tinc functionality via Webservices
 */

const WebServices = {

    async defineHeaderAxios () {
        await Mgr.getAcessToken().then(
          acessToken => {
            axios.defaults.headers.common['Authorization'] = 'Bearer ' + acessToken
          }, err => {
            console.log(err)
          })  
      },
    

    get url() {
        if (!this._url){
            this._url = "http://localhost:8000"
        }
        return this._url
    },

    set url(url){
        this._url = url
    },

    async getNetworks() {
        return await axios.get(this.url + '/api/networks')
    },

    async getInvitations() {
        return await axios.get(this.url + '/api/invitations')
    },

    async getParameters() {
        return await axios.get(this.url + '/api/parameters')
    },

    async getNodes() {
        return await axios.get(this.url + '/api/nodes')
    },

    async getEdges() {
        return await axios.get(this.url + '/api/edges')
    },

    async get(command) {
        return await axios.get(this.url + '/api/'+command)
    },

    // E.g start, stop, restart, purge, generate-keys
    async action(command) {
        return await axios.post(this.url + '/api/'+command)
    },

    // download blob
    async getDownload(command) {
        return axios.post(this.url + "/api/"+command)
    },

    // create a URL to a Blob for the indicated mime object
    async getDownloadLink(command) {
        var response = await this.getDownload(command)
        if (response.request.status == 200) {
            const url = URL.createObjectURL(new Blob([response.data], {type: 'text/plain'}));
            return await url
        } else {
            throw new Error("DOWNLOAD-ERROR")
        }
    },

    // import a file
    async doImport(file) {
        var formData = new FormData();
        formData.append('file', file);
        return await axios.post(this.url + '/api/import', formData)
    },

    // process invite
    async invite(nodeName) {
        return await axios.post(this.url + '/api/invite?node='+nodeName)
    },

    // process join-invite
    async joinInvitation(invitation) {
        return await axios.post(this.url + '/api/join?invitation='+invitation)
    },

    // get performance info
    async getNetworkTraffic(asBytes) {
        return await axios.get(this.url + '/api/network-traffic?bytes='+asBytes)
    },

    async getNetworkTrafficStatus() {
        return await axios.get(this.url + '/api/network-traffic-recording')
    },

    async networkTrafficOn() {
        return await axios.post(this.url + '/api/network-traffic-recording')
    },

    async networkTrafficOff() {
        return await axios.delete(this.url + '/api/network-traffic-recording')
    },

    async ping() {
        return await axios.get(this.url + '/api/ping')
    },

    async getParameter(name) {
        return await axios.get(this.url + '/api/parameter?name='+name)
    },

    async setParameter(name, value) {
        return await axios.post(this.url + '/api/parameter',{value: value})
    },


    async init(name){
        return await axios.post(this.url + '/api/init',{name:name})
    },

    async remoteExchange(url){
        return await axios.post(this.url + '/api/remote-exchange', {url: url})
    },

    async setup(setupData){
        var initResult = await this.init(setupData.nodeName)        
        await this.setParameter('Subnet',setupData.subnet)
        await this.setParameter('VpnIP',setupData.localIP)
        await this.setParameter('AutoConnect','on')
        if(setupData.invitation){
            return await this.joinInvitation(setupData.invitation)
        }
        if(setupData.url){
           return await this.remoteExchange(setupData.url)
        }
        return initResult
    }

}

export default WebServices