package service

/**
 * Registration of all web service handlers for /api webservices
 */
import (
	"github.com/gorilla/mux"
)

//RegisterAPIServices - defines the handler functions for the /api/ services
func RegisterAPIServices(path string, r *mux.Router) {
	api := r.PathPrefix(path).Subrouter()

	api.HandleFunc("/ping", Ping).Methods("GET")
	api.HandleFunc("/import", ImportHandler).Methods("POST")

	// simple commands
	api.HandleFunc("/fsck", NewCommand("fsck").Handler()).Methods("GET")
	api.HandleFunc("/parameters", TincParameters).Methods("GET")
	api.HandleFunc("/stop", NewCommand("stop").Handler()).Methods("POST")
	api.HandleFunc("/start", NewCommand("start").Handler()).Methods("POST")
	api.HandleFunc("/retry", NewCommand("retry").Handler()).Methods("POST")
	api.HandleFunc("/exchange", NewCommand("exchange").Handler()).Methods("POST")
	api.HandleFunc("/exchange-all", NewCommand("exchange-all").Handler()).Methods("POST")
	api.HandleFunc("/restart", NewCommand("restart").Handler()).Methods("POST")
	api.HandleFunc("/reload", NewCommand("reload").Handler()).Methods("POST")
	api.HandleFunc("/retry", NewCommand("retry").Handler()).Methods("POST")
	api.HandleFunc("/purge", NewCommand("purge").Handler()).Methods("POST")
	api.HandleFunc("/generate-keys", NewCommand("generate-keys").Handler()).Methods("POST")
	api.HandleFunc("/generate-rsa-keys", NewCommand("generate-rsa-keys").Handler()).Methods("POST")
	api.HandleFunc("/generate-ed25519-keys", NewCommand("generate-ed25519-keys").Handler()).Methods("POST")
	api.HandleFunc("/pid", NewCommand("pid").Handler()).Methods("GET")
	api.HandleFunc("/export", NewCommand("export").Handler()).Methods("POST")
	api.HandleFunc("/export-all", NewCommand("export-all").Handler()).Methods("POST")

	// commands with additioinal formatting
	api.HandleFunc("/networks", NewCommand2(ArrayFormatter, "network").Handler()).Methods("GET")
	api.HandleFunc("/invitations", NewCommand2(InvitationsFormatter, "dump", "invitations").Handler()).Methods("GET")
	api.HandleFunc("/nodes", NewCommand2(NodesFormatter, "dump", "nodes").Handler()).Methods("GET")
	api.HandleFunc("/connections", NewCommand2(ConnectionsFormatter, "dump", "connections").Handler()).Methods("GET")
	api.HandleFunc("/subnets", NewCommand2(SubnetsFormatter, "dump", "subnets").Handler()).Methods("GET")
	api.HandleFunc("/edges", NewCommand2(ArrayFormatter, "dump", "edges").Handler()).Methods("GET")
	api.HandleFunc("/graph", NewCommand2(DottyFormatter, "dump", "graph").Handler()).Methods("GET")
	api.HandleFunc("/digraph", NewCommand2(DottyFormatter, "dump", "digraph").Handler()).Methods("GET")

	api.HandleFunc("/info", NewCommand3(InfoFormatter, "name", "node").Handler()).Methods("GET")

	// commands with 1 parameter
	api.HandleFunc("/verify", NewCommand3(CommandFormatter, "node", "verify").Handler()).Methods("GET")
	api.HandleFunc("/network", NewCommand3(CommandFormatter, "netname", "network").Handler()).Methods("POST")
	api.HandleFunc("/join", NewCommand3(CommandFormatter, "invitation", "join").Handler()).Methods("POST")
	api.HandleFunc("/invite", NewCommand3(CommandFormatter, "node", "invite").Handler()).Methods("POST")
	api.HandleFunc("/disconnect", NewCommand3(CommandFormatter, "node", "disconnect").Handler()).Methods("POST")
	api.HandleFunc("/init", NewCommand3(CommandFormatter, "name", "init").Handler()).Methods("POST")
	api.HandleFunc("/parameter", NewCommand3(CommandFormatter, "name", "get").Handler()).Methods("GET")
	api.HandleFunc("/parameter", NewCommand3(CommandFormatter, "name", "del").Handler()).Methods("DELETE")

	// commadns with 2 parameters
	api.HandleFunc("/parameter", NewCommand4(CommandFormatter, "name", "value", "add").Handler()).Methods("POST")
	api.HandleFunc("/parameter", NewCommand4(CommandFormatter, "name", "value", "set").Handler()).Methods("PUT")


}

