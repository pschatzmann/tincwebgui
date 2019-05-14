package service

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Ping - Execute ping command
func Ping(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"response":  "pong",
		"timeStamp": time.Now().Format(time.RFC3339),
	}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	log.Println("ping")
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
