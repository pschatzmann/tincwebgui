import axios from 'axios'
/**
 * Access Tinc functionality via Webservices
 */

const WebServices = {

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
        return await axios.post(this.url + '/api/import', formData, { headers: { 'Content-Type': 'multipart/form-data'}})
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


}

export default WebServices