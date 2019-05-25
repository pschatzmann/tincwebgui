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
	"github.com/pascaldekloe/jwt"
)

// Auth - authentication management information
type Auth struct {
	password       string
	passwordActive bool
	serviceActive  bool
}

var authValue = Auth{"", false, false}

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
func checkAuthorization(token string) error {
	log.Println("Authorization", token)
	var result error
	if !authValue.serviceActive {
		result = errors.New("The service has not been activated. Please call export SERVICE_ACTIVE=true")
	} else {
		result = checkToken(token)
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
		} else {
			// no password necessary
			authValue.passwordActive = false
		}
	} else {
		authValue.serviceActive = false
	}
	return authValue
}

func checkToken(tokenString string) error {
	// verify a JWT

	return nil
}

func getPublicKey() *string {
	env := os.Getenv("SERVICE_ACTIVE")
	return &env
}

func checkClaims(claims *jwt.Claims) error {
	claimsString := os.Getenv("PASSWORD_ACTIVE")
	claimsArray := strings.Split(claimsString, ";")

	for _, claim := range claimsArray {
		key, values := splitClaim(claim)
		claimValue, ok := claims.String(key)
		if !ok || claimValue == "" {
			return fmt.Errorf("The expected claim '%q' is empty ", key)
		}
		if !(contains(values, claimValue)) {
			return fmt.Errorf("The expected claim '%q' with value '%q' is not valid ", key, claimValue)
		}
	}
	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func splitClaim(claim string) (string, []string) {
	sa := strings.Split(claim, "=")
	return sa[0], strings.Split(sa[1], ",")
}
