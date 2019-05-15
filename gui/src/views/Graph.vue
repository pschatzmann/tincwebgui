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

        renderer : {
            container: null,
            type: 'canvas'
        }

    }),

    computed: {
        error() {
            return this.$store.state.error
        },
    },

    methods: {
        // setupTestData(){
        //     var i,
        //     N = 20,
        //     E = 50

        //     // Generate a random graph:
        //     for (i = 0; i < N; i++){
        //         this.graph.nodes.push({
        //             id: 'n' + i,
        //             label: 'Node ' + i,
        //             size: 1,
        //             color: '#666'
        //         });
        //     }

        //     // allocate on cirle
        //     this.graph.nodes.forEach(function(node, i, a) {
        //         node.x = Math.cos(Math.PI * 2 * i / a.length);
        //         node.y = Math.sin(Math.PI * 2 * i / a.length);
        //     });

        //     for (i = 0; i < E; i++) {
        //         var sourceId = 'n'+Math.floor(Math.random() * N )
        //         var targetId = 'n'+Math.floor(Math.random() * N)
        //         if (sourceId!=targetId){
        //             this.graph.edges.push({
        //                 id: 'e' + i,
        //                 label: 'Edge ' + i,
        //                 source: sourceId,
        //                 target: targetId,
        //                 size: 10,
        //                 color: '#ccc',
        //             });
        //         }
        //     }
        //     renderGraph()
        // },

        setup() {
            WebServices.getNodes().then(result => {
                this.resultData.nodes = result.data
                // convert to display format
                this.graph.nodes = this.resultData.nodes.map(n => {id: n.id; label: n.name; color: '#666'; size: 1 })
                // allocate on cirle
                this.graph.nodes.forEach(function(node, i, a) {
                    node.x = Math.cos(Math.PI * 2 * i / a.length);
                    node.y = Math.sin(Math.PI * 2 * i / a.length);
                });                
                renderGraph()

            },error => {
                this.$store.dispatch('setError', error)
            })
        },

        renderGraph() {
            // define container
            this.renderer.container = this.$refs.sigmaContainer

            // setup sigma
            var s = new sigma({ renderer: this.renderer, settings: this.settings})
            s.graph.read(this.graph);
            s.refresh();

            s.bind('overEdge',(e)=>{ 
            // e.data.edge.color = '#ff6347'
                console.log(e.data);
            }) 

            s.bind('clickEdge', function(e) {            
                console.log(e.data);
            });
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
</style>
