package service

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// ConfigDelete -
func ConfigDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var err = os.Remove("/etc/tinc/" + name + ".conf")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
