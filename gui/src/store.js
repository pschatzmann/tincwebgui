import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        isProcessing: false,
        isActive: true,

        error: null,
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

        SET_ACTIVE(state, active) {
            state.isActive = active
        },

    },
    
    actions: {
        setProcessing(context, status) {
            context.commit("SET_PROCESSING", status)
        },

        setActive(context, status) {
            context.commit("SET_ACTIVE", status)
        },

        setError(context, error) {
            var errorMsg = error
            if (error && error.response && error.response.data) {
                errorMsg = error.response.data
            }
            context.commit("SET_ERROR", errorMsg)
        },

        setNavigationDrawer(context, visible) {
            context.commit("SET_NAVIGATION_DAWER", visible)
        },

        setNavigationDrawerPermanent(context, permanent) {
            context.commit("SET_NAVIGATION_DAWER_PERMANENT", permanent)
        },


    },
})
    