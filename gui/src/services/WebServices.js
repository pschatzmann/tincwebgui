import axios from 'axios'
/**
 * Access Tinc functionality via Webservices
 */
const WebServices = {
    get url() {
        var result = process.env.ServiceURL;
        if (!result) {
             result = "http://localhost:8000"
        }
        return result
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


    // start, stop, restart, purge, generate-keys

    async action(command) {
        return await axios.post(this.url + '/api/'+command)
    },

    async export() {
        return await axios.post(this.url + '/api/export')
    },

    async exportAll() {
        return await axios.post(this.url + '/api/exportAll')
    },

    async doImport() {
        return await axios.post(this.url + '/api/import')
    },

    async invite() {
        return await axios.post(this.url + '/api/invite')
    },

    async joinInvite() {
        return await axios.post(this.url + '/api/joinInvite')
    }
}

export default WebServices