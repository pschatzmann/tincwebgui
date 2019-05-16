package auth

/**
* Simple Password management for the tinc webservices.
 */
import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// Auth - authentication management information
type Auth struct {
	password       string
	passwordActive bool
	serviceActive  bool
}

var authValue = Auth{"", false, false}

// Generate - Generates a new Password
func Generate() string {
	f, _ := os.Open("/dev/urandom")
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	pwd := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return pwd
}

// Check - checks the password which is passed as bearer Authorization token
func Check(r *http.Request) error {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if (len(splitToken)) > 1 {
		reqToken = splitToken[1]
	}
	return checkAuthorization(reqToken)
}

// checkPassword - checks the password if password management is active
func checkAuthorization(pwd string) error {
	var result error
	if !authValue.serviceActive {
		result = errors.New("The service has not been activated. Please call export SERVICE_ACTIVE=true")
	} else {
		if authValue.passwordActive && pwd != authValue.password {
			result = errors.New("The password is not valid")
		}
	}
	return result
}

// Setup - sets up the password handling from the Env variables PASSWORD-ACTIVE and PASSWORD
func Setup(host string, r *mux.Router) Auth {
	serviceActiveParameter := os.Getenv("SERVICE_ACTIVE")
	if serviceActiveParameter != "false" {
		authValue.serviceActive = true
		passwordActiveParameter := os.Getenv("PASSWORD_ACTIVE")
		if passwordActiveParameter != "false" {
			authValue.passwordActive = true

			// we set the passoword if it has been defined, otherwise we generate one
			passwordParameter := os.Getenv("PASSWORD")
			if passwordParameter != "" {
				authValue.password = passwordParameter
				log.Println("The password management is active - please use your defined password to authenticate the webservices")
			} else {
				authValue.password = Generate()
				log.Println("Please use the following password to authenticate the webservices: ", authValue.password)
			}
		} else {
			// no password necessary
			authValue.passwordActive = false
		}
	} else {
		authValue.serviceActive = false
	}
	return authValue
}
