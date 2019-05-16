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

    async get(command) {
        return await axios.get(this.url + '/api/'+command)
    },

    // start, stop, restart, purge, generate-keys

    async action(command) {
        return await axios.post(this.url + '/api/'+command)
    },


    async getDownload(command) {
        await this.defineHeaderAxios()
        return axios.get(this.url + "/api/"+command, {
            responseType: 'arraybuffer'
        })
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