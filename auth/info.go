package auth

import (
	"encoding/json"
	"log"
	"net/http"
)

// AuthInfo - Provide ProviderUrl and ClientId
func AuthInfo(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"ProviderUrl": *getProviderUrl(),
		"ClientId":    *getClientId(),
	}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	log.Println("AuthInfo")
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
