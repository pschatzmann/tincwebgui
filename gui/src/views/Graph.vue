<template>
    <div>
        <v-alert :value="error!=null" type="error">{{error}}</v-alert>

        <v-container fluid>
            <v-card>
                <v-container fluid>
                    <div ref="sigmaContainer" id='sigmaContainer'/>
                </v-container>
            </v-card>
        </v-container>
    </div>
</template>


<script>
import sigma from 'sigma'
import WebServices from "@/services/WebServices"
import '@/plugins/sigma.plugins.tooltips.js'

export default {
    name: "graph",
    data: () => ({
        graph : {
            nodes: [],
            edges: [],
            verbose: true
        },

        resultData : {
            nodes: [],
            edges: []
        },

        settings: {
            minEdgeSize: 0.5,
            maxEdgeSize: 2,
            minNodeSize: 5,
            maxNodeSize: 8,
            enableEdgeHovering: true,
            labelThreshold: 1,
            defaultEdgeColor: '#917373',
            defaultEdgeType:'line',
            enableHovering:true,
            edgeColor: "default",
            edgeHoverColor: '#ff6347',
            defaultEdgeHoverColor: '#ff6347',
            edgeHoverSizeRatio: 1,
            edgeHoverExtremities: true,
            fixedScaling: true,
            zoomMin: 1,
            zoomMax: 1
        },

        tooltipSetup : {
            node: {
                show: 'overNode',
                hide: 'outNode',
                renderer: null
            },
            edge: {
                show: 'overEdge',
                hide: 'outEdge',
                renderer: null
            },  
        },

        renderer : {
            container: null,
            type: 'canvas'
        },
    }),

    computed: {
        error() {
            return this.$store.state.error
        },
    },

    methods: {
        // Tooltip renderer
        tooltipRenderer(obj, v1, v2){
            console.log(obj,v1,v2)
            if (obj.ref) {
                return this.makeTableHTML(obj.ref.entries())
            } else {
                return JSON.stringify(obj)
            }
        },

        // convert 
        makeTableHTML(myArray) {
            var result = "<table>";
            for(var i=0; i<myArray.length; i++) {
                result += "<tr>";
                for(var j=0; j<myArray[i].length; j++){
                    result += "<td>"+myArray[i][j]+"</td>";
                }
                result += "</tr>";
            }
            result += "</table>";
            return result;
        },


        setup() {
            const thisCtx = this
            WebServices.getNodes().then(result => {
                thisCtx.resultData.nodes = result.data
                // convert to display format
                thisCtx.graph.nodes = thisCtx.resultData.nodes.map(n => {
                    return {id: n.id, label: n.name, color: '#666', size: 1, ref: n }
                })
                // allocate on cirle
                thisCtx.graph.nodes.forEach(function(node, i, a) {
                    node.x = Math.cos(Math.PI * 2 * i / a.length);
                    node.y = Math.sin(Math.PI * 2 * i / a.length);
                });                
                thisCtx.renderGraph()

            },error => {
                this.$store.dispatch('setError', error)
            })
        },

        renderGraph() {
            // define container
            this.renderer.container = this.$refs.sigmaContainer

            // setup sigma
            var s = new sigma({ renderer: this.renderer, settings: this.settings})

            // setup tooltips
            this.tooltipSetup.node.renderer = this.tooltipRenderer
            this.tooltipSetup.edge.renderer = this.tooltipRenderer
            sigma.plugins.tooltips(s, s.renderers[0],this.tooltipSetup);

            s.graph.read(this.graph);
            s.refresh();
        }
    },

    mounted() {
        this.setup()
    }
}

</script>

<style>
    #sigmaContainer {
        height:80vh;
    }

    .sigma-tooltip {
        max-width: 240px;
        max-height: 280px;
        background-color: black;
        border: 1px solid ;
        box-shadow: 0 2px 6px rgba(0, 0, 0, 0.3);
        border-radius: 6px;
        cursor: auto;
        font-family: Arial;
        font-size: 12px;
        color: white;
        font-weight: bold 
    }

</style>
