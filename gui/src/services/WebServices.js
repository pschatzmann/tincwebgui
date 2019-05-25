import axios from 'axios'
import { SecurityService } from '@/services/SecurityService'

/**
 * Access Tinc functionality via Webservices
 */

const WebServices = {

    async defineHeaderAxios () {
        var acessToken = await SecurityService.getAcessToken()
        axios.defaults.headers.common['Authorization'] = 'Bearer ' + acessToken
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
        await this.defineHeaderAxios();
        return await axios.get(this.url + '/api/networks')
    },

    async getInvitations() {
        await this.defineHeaderAxios();
        return await axios.get(this.url + '/api/invitations')
    },

    async getParameters() {
        await this.defineHeaderAxios();
        return await axios.get(this.url + '/api/parameters')
    },

    async getNodes() {
        await this.defineHeaderAxios();
        return await axios.get(this.url + '/api/nodes')
    },

    async getEdges() {
        await this.defineHeaderAxios();
        return await axios.get(this.url + '/api/edges')
    },

    async get(command) {
        await this.defineHeaderAxios();
        return await axios.get(this.url + '/api/'+command)
    },

    // E.g start, stop, restart, purge, generate-keys
    async action(command) {
        await this.defineHeaderAxios();
        return await axios.post(this.url + '/api/'+command)
    },

    // download blob
    async getDownload(command) {
        await this.defineHeaderAxios();
        return axios.post(this.url + "/api/"+command)
    },

    // create a URL to a Blob for the indicated mime object
    async getDownloadLink(command) {
        await this.defineHeaderAxios();
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
        await this.defineHeaderAxios();
        var formData = new FormData();
        formData.append('file', file);
        return await axios.post(this.url + '/api/import', formData)
    },

    // process invite
    async invite(nodeName) {
        await this.defineHeaderAxios();
        return await axios.post(this.url + '/api/invite?node='+nodeName)
    },

    // process join-invite
    async joinInvitation(invitation) {
        await this.defineHeaderAxios();
        return await axios.post(this.url + '/api/join?invitation='+invitation)
    },

    // get performance info
    async getNetworkTraffic(asBytes) {
        await this.defineHeaderAxios();
        return await axios.get(this.url + '/api/network-traffic?bytes='+asBytes)
    },

    async getNetworkTrafficStatus() {
        await this.defineHeaderAxios();
        return await axios.get(this.url + '/api/network-traffic-recording')
    },

    async networkTrafficOn() {
        await this.defineHeaderAxios();
        return await axios.post(this.url + '/api/network-traffic-recording')
    },

    async networkTrafficOff() {
        await this.defineHeaderAxios();
        return await axios.delete(this.url + '/api/network-traffic-recording')
    },

    async ping() {
        await this.defineHeaderAxios();
        return await axios.get(this.url + '/api/ping')
    },

    async getParameter(name) {
        await this.defineHeaderAxios();
        return await axios.get(this.url + '/api/parameter?name='+name)
    },

    async setParameter(name, value) {
        await this.defineHeaderAxios();
        return await axios.post(this.url + '/api/parameter',{name: name, value: value})
    },


    async init(name){
        await this.defineHeaderAxios();
        return await axios.post(this.url + '/api/init',{name:name})
    },

    async remoteExchange(url){
        await this.defineHeaderAxios();
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
    },

    async deleteConfig(name){
        await this.defineHeaderAxios();
        return await axios.delete(this.url + '/api/config/' + name )
    }

}

export default WebServices