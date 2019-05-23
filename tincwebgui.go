package main

/**
*  Server for Javascript Vue GUI implemented in GO
**/
import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/pschatzmann/tincwebgui/auth"
	"github.com/pschatzmann/tincwebgui/gui"
	"github.com/pschatzmann/tincwebgui/service"
	"github.com/pschatzmann/tincwebgui/utils"
)

func main() {
	address := "localhost:8000"
	if len(os.Args) >= 2 {
		address = os.Args[1]
	}

	// start tinc
	service.StartTinc()

	r := mux.NewRouter()
	// setup authrizations
	auth.Setup(address, r)

	// Register API services
	service.RegisterAPIServices("/api/", r)

	// Serve static assets directly.
	r.PathPrefix("/").Handler(htmlHandler())
	// listen on localhost:8000 or as specified in the arg 1

	srv := &http.Server{
		Handler: r,
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	utils.PrintIP(&address)
	log.Fatal(srv.ListenAndServe())
}

func htmlHandler() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		setupResponse(&w, req)
		if (*req).Method == "OPTIONS" {
			return
		}

		path := req.URL.Path
		file, err := gui.Assets.Open(path)

		if err == nil && path != "/" {
			defer file.Close()
			fileInfo, err := file.Stat()
			if err == nil {
				http.ServeContent(w, req, path, fileInfo.ModTime(), file)
				return
			}
		}
		// the all 404 gonna be served as root
		file, err = gui.Assets.Open("/index.html")
		defer file.Close()
		fileInfo, err := file.Stat()
		http.ServeContent(w, req, path, fileInfo.ModTime(), file)
	})
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
