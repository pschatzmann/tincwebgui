package main

/**
*  Server for Javascript Vue GUI implemented in GO
**/
import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/pschatzmann/tincwebgui/auth"
	"github.com/pschatzmann/tincwebgui/service"
	"github.com/pschatzmann/tincwebgui/utils"
)

func main() {
	address := "localhost:8000"
	if len(os.Args) >= 2 {
		address = os.Args[1]
	}
	r := mux.NewRouter()
	// setup authrizations
	auth.Setup(address, r)
	//auth.SetupGoth(address, r)

	// Register API services
	service.RegisterAPIServices("/api/", r)

	// Serve static assets directly.
	path := getPath()
	log.Println("Path is ", path)
	r.PathPrefix("/").Handler(htmlHandler(path))
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

func htmlHandler(publicDir string) http.Handler {
	handler := http.FileServer(http.Dir(publicDir))

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		_path := req.URL.Path

		// static files
		if strings.Contains(_path, ".") || _path == "/" {
			handler.ServeHTTP(w, req)
			return
		}
		// the all 404 gonna be served as root
		http.ServeFile(w, req, path.Join(publicDir, "/index.html"))
	})
}

// check if the file exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// determines the path for the javascript resources. It might be in the actual directory or in the dist subdirectory
func getPath() string {
	path := ""
	if len(os.Args) >= 3 {
		// get path from arguments
		path = os.Args[2]
	} else if fileExists("./index.html") {
		path = "."
	} else if fileExists("dist/index.html") {
		path = "dist"
	}
	return path
}
