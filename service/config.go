package service

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// ConfigDelete -
func ConfigDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName := "/etc/tinc/" + vars["name"] + ".conf"
	_, err := os.Stat(fileName)
	if err == nil {
		log.Printf("File %s exists and will be deleted", fileName)
		err = os.Remove(fileName)
	} else if os.IsNotExist(err) {
		log.Printf("File %s not exists - so no deletion necessary", fileName)
	} else {
		log.Printf("file %s stat error: %v", fileName, err)
		http.Error(w, err.Error(), 400)
	}
	w.Write([]byte("OK"))
}
