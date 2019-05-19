/<!--
  Display and manage Network Information
 -->
 
 <template>
    <div id="sticky">
        <v-toolbar  class="my-toolbar">
            <v-tooltip bottom>
                <span>Export host configuration of local node </span>
                <template v-slot:activator="{ on }">
                    <v-btn v-on="on" v-on:click="saveAs('export')" >
                        <v-icon>cloud</v-icon>
                    </v-btn>   
                </template>
            </v-tooltip>

            <v-tooltip bottom>
                <span>Export all host configuration files</span>
                <template v-slot:activator="{ on }">
                    <v-btn v-on="on" v-on:click="saveAs('export-all')" >
                        <v-icon>cloud_circle</v-icon>
                    </v-btn>   
                </template>
            </v-tooltip>

            <v-tooltip bottom>
                <span>Import host configuration file(s)</span>
                <template v-slot:activator="{ on }">
                    <v-btn v-on="on" v-on:click="doImport()"  >
                        <v-icon>cloud_download</v-icon>
                    </v-btn>   
                </template>
            </v-tooltip>

            <v-tooltip bottom>
                <span>Purge unreachable nodes</span>
                <template v-slot:activator="{ on }">
                    <v-btn v-on="on" v-on:click="purge()"  >
                        <v-icon>delete_outline</v-icon>
                    </v-btn>   
                </template>
            </v-tooltip>
            <v-spacer/>
        </v-toolbar>
        <v-alert :value="error.msg!=null" :type="error.type">{{error.msg}}</v-alert>
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
import '@/components/MyInputDialog'

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
                autoadjust: true,
                show: 'overNode',
                hide: 'outNode',
                renderer: null
            },
            edge: {
                autoadjust: true,
                show: 'overEdge',
                hide: 'outEdge',
                renderer: null
            },  
        },

        inputDialog: {
            visible: false
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
            var ref = obj['ref']
            if (ref) {
                return this.makeTableHTML(Object.entries(ref))
            } else {
                return JSON.stringify(obj)
            }
        },

        // convert 
        makeTableHTML(myArray) {
            var result = "<table>";
            for(var i=0; i<myArray.length; i++) {
                if (myArray[i][1]!='0') {
                    result += "<tr>";
                    for(var j=0; j<myArray[i].length; j++){
                        if (j==0){
                            result += "<td align='right'>"+this.firstCharAsCaps(myArray[i][j])+": </td>";
                        } else {
                            result += "<td>"+myArray[i][j]+"</td>";
                        }
                    }
                    result += "</tr>";
                }
            }
            result += "</table>";
            return result;
        },

        firstCharAsCaps(string) {
            return (string.charAt(0).toUpperCase() + string.slice(1)).trim();
        },   

        setup() {
            const thisCtx = this
            WebServices.getNodes().then(result => {
                thisCtx.resultData.nodes = result.data
                // convert to display format
                thisCtx.graph.nodes = thisCtx.resultData.nodes.map(n => {
                    return {id: n.id, label: n.name, color: '#6B0700', size: 1, ref: n }
                })
                // allocate on cirle
                thisCtx.graph.nodes.forEach(function(node, i, a) {
                    node.x = Math.cos(Math.PI * 2 * i / a.length);
                    node.y = Math.sin(Math.PI * 2 * i / a.length);
                });         
                
                WebServices.getEdges().then(result => {
                    thisCtx.resultData.edges = result.data
                    // convert to display format
                    thisCtx.graph.edges = thisCtx.resultData.edges.map(n => {
                        return {id: 'id'+n.from+'-'+n.to, source: n.from, target: n.to }
                    })

                    thisCtx.renderGraph()
                },error => {
                    this.$store.dispatch('setError', error)
                })

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
        },

        // execute post action
        action(action) {
            this.$store.dispatch("setError", null)
            var self = this
            WebServices.action(action).then( result => {
                console.log(result); 
            }, err => {
                self.$store.dispatch('setError', err)
            });
        },

        // export and display save as dialog
        saveAs(action){
            this.$store.dispatch("setError", null)
            var self = this
            WebServices.getDownloadLink(action).then(url => {
                var link = document.getElementById("documentContent")
                if (link == null) {
                    link = document.createElement('a');
                    link.setAttribute("id", "documentContent");
                }
                link.setAttribute('Download', action+'.txt' );
                link.href = url;
                document.body.appendChild(link);
                link.click();
             }).catch((error) => {
                self.$store.dispatch('setError', error)
            })
        },

        // import a selected file
        doImport() {
            var input = document.createElement('input');
            input.type = 'file';
            input.onchange = e => { 
                var file = e.target.files[0]; 
                WebServices.import(file).then(result => {
                    console.log(result)
                }).catch((error) => {
                    self.$store.dispatch('setError', error)
                })           
            }
            input.click();
        },

        // purge unreachable nodes
        purge() {
            this.action('purge')
            // reload data
            this.setup()
        }

    },

    mounted() {
        this.setup()
    }
}

</script>

<style>
    #sticky {
        top: 63px;
        height: 63px;
        z-index: 1;
        position: -webkit-sticky;
        position: sticky;
        visibility: visible;
    }

    .myToolbar > .v-toolbar__content {
        height: 63px
    }

    #sigmaContainer {
        height:70vh;
    }

    .sigma-tooltip {
        max-width: 240px;
        max-height: 280px;
        background-color: lightyellow;
        border: 1px solid ;
        box-shadow: 0 2px 6px rgba(0, 0, 0, 0.3);
        border-radius: 12px;
        cursor: auto;
        font-family: Arial;
        font-size: 9px;
        color: black;
        font-weight: normal;
        padding-left: 20px;
        padding-right: 20px;
    }

</style>
