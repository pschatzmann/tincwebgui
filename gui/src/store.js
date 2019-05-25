import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        isProcessing: false,
        isLoggedIn: false,
        isActive: false,

        error: { msg: null, type:'error'},
        // navigation
        navigationDrawer: {
            permanent: true,
            drawer: true,
        },
    },

    mutations: {
        SET_PROCESSING(state, proc) {
            state.isProcessing = proc
        },
    
        SET_ERROR(state, error) {
            state.error = error
        },

        SET_NAVIGATION_DAWER(state, visible) {
            state.navigationDrawer.drawer = visible
        },

        SET_NAVIGATION_DAWER_PERMANENT(state, permanent) {
            state.navigationDrawer.permanent = permanent
        },

        SET_LOGGED_IN(state, active) {
            state.isLoggedIn = active
        },

        SET_ACTIVE(state, active) {
            state.isActive = active
        },

    },
    
    actions: {
        setProcessing(context, status) {
            context.commit("SET_PROCESSING", status)
        },

        setLoggedIn(context, status) {
            context.commit("SET_LOGGED_IN", status)
        },

        setActive(context, status) {
            context.commit("SET_ACTIVE", status)
        },

        setError(context, error) {
            if (error==null) {
                context.commit("SET_ERROR", {msg: null})
                return
            }  
            var errorMsg = error
            if (error && error.response && error.response.data) {
                errorMsg = error.response.data
            } else if (error && error.msg){
                errorMsg = error.msg
            }            
            var type = 'error'
            if (error && error.type){
                type = error.type
            }
            errorMsg = errorMsg.split("\n").join("<br>")
            context.commit("SET_ERROR", {msg: errorMsg, type: type})
        },

        setNavigationDrawer(context, visible) {
            context.commit("SET_NAVIGATION_DAWER", visible)
        },

        setNavigationDrawerPermanent(context, permanent) {
            context.commit("SET_NAVIGATION_DAWER_PERMANENT", permanent)
        },
    },

    getters: {
        isEnabled: state => state.isLoggedIn && state.isActive
      }
    
})
    