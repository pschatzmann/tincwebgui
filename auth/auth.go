package auth

/**
* Simple Password management for the tinc webservices.
 */
import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/antonmedv/expr"
	oidc "github.com/coreos/go-oidc"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

// Auth - authentication management information
type Auth struct {
	passwordActive bool
	serviceActive  bool
}

var authValue = Auth{true, true}

// Authenticator - Authenticator
type Authenticator struct {
	provider     *oidc.Provider
	clientConfig oauth2.Config
	ctx          context.Context
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

// checkAuthorization - checks the authorization token if password management is active
func checkAuthorization(token string) error {
	log.Println("Authorization", token)
	var result error
	if !authValue.serviceActive {
		result = errors.New("The service has not been activated. Please call export SERVICE_ACTIVE=true")
	} else {
		result = checkToken(strings.TrimSpace(token))
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
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, *getProviderUrl())
	if err != nil {
		log.Fatalf("Failed to get provider: %v", err)
	}

	config := oauth2.Config{
		ClientID:     *getClientId(),
		ClientSecret: *getClientSecret(),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	oidcConfig := &oidc.Config{
		ClientID: config.ClientID,
	}

	a := &Authenticator{
		provider:     provider,
		clientConfig: config,
		ctx:          ctx,
	}

	idToken, err := a.provider.Verifier(oidcConfig).Verify(a.ctx, tokenString)
	if err != nil {
		log.Println("Error: Could not decode token", err)
		return err
	}

	claimsMap := make(map[string]interface{})
	idToken.Claims(&claimsMap)

	// get claims to map
	for k, v := range claimsMap {
		fmt.Println("key:", k, "value:", v)
	}

	evalStr := *getClaimsEval()
	if evalStr != "" {
		out, err := expr.Eval(evalStr, claimsMap)
		if err != nil {
			log.Println("Error: Could not evaluate claims", err)
			return err
		}

		evalResult := fmt.Sprint(out)
		if evalResult != "false" {
			log.Println("Claims Evaluation is "+evalResult, err)
			return errors.New("Evaluation result is " + evalResult)
		}
	}

	return nil
}

func getProviderUrl() *string {
	env := os.Getenv("PROVIDER_URL")
	return &env
}

func getClientId() *string {
	id := os.Getenv("CLIENT_ID")
	return &id
}

func getClientSecret() *string {
	secret := os.Getenv("CLIENT_SECRET")
	return &secret
}

func getClaimsEval() *string {
	claims := os.Getenv("CLAIMS")
	return &claims
}
