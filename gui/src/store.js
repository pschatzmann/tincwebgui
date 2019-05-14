import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        isProcessing: false,

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
        }
    },
    
    actions: {
        setProcessing(context, status) {
            context.commit("SET_PROCESSING", status)
        },

        setError(context, error) {
            context.commit("SET_ERROR", error)
        },

        setNavigationDrawer(context, visible) {
            context.commit("SET_NAVIGATION_DAWER", visible)
        },

        setNavigationDrawerPermanent(context, permanent) {
            context.commit("SET_NAVIGATION_DAWER_PERMANENT", permanent)
        },
    },
})
    