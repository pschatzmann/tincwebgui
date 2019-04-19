package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// It's important that this is before your catch-all route ("/")
	api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/users", GetUsersHandler).Methods("GET")
	// Serve static assets directly.
	r.PathPrefix("/static").Handler(http.FileServer(http.Dir("dist/")))
	// Catch-all: Serve our JavaScript application's entry-point (index.html).
	r.PathPrefix("/").HandlerFunc(IndexHandler("dist/index.html"))

	address := "127.0.0.1:8000"
	srv := &http.Server{
		Handler: r,
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server started at ", address)
	log.Fatal(srv.ListenAndServe())

}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"id":        "123",
		"timeStamp": time.Now().Format(time.RFC3339),
	}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write(b)
}
