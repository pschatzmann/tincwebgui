package service

/**
*  Hander for tinc import command.
**/
import (
	"log"
	"net/http"
	"os/exec"
)

// ImportHandler - Process Import of file via upload
func ImportHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("tinc", "import")
	cmd.Stdin = r.Body
	out, err := cmd.CombinedOutput()
	if err != nil {
		resultString := string(out)
		log.Println(resultString)
		http.Error(w, resultString, 400)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write(out)
}
