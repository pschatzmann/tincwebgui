package auth

import (
	"os"
	"testing"
)

func TestNoService(t *testing.T) {
	os.Setenv("SERVICE_ACTIVE", "false")
	Setup("",nil)
	err := checkAuthorization("na")
	expected := "The service has not been activated. Please call export SERVICE_ACTIVE=true"
	if err.Error() != expected {
		t.Errorf("The error was %v instead of %v", err, expected)
	}
}

func TestNoPassword(t *testing.T) {
	os.Setenv("SERVICE_ACTIVE", "true")
	os.Setenv("PASSWORD_ACTIVE", "false")
	Setup("",nil)
	err := checkAuthorization("na")
	if err != nil {
		t.Errorf("The checkAuthorization has failed with PASSWORD_ACTIVE false")
	}
}

func TestPassword(t *testing.T) {
	os.Setenv("SERVICE_ACTIVE", "true")
	os.Setenv("PASSWORD_ACTIVE", "true")
	os.Setenv("PASSWORD", "test")
	Setup("", nil)
	err := checkAuthorization("test")
	if err != nil {
		t.Errorf("The checkAuthorization has failed with PASSWORD_ACTIVE false")
	}
}

func TestGeneratedPassword(t *testing.T) {
	os.Setenv("SERVICE_ACTIVE", "true")
	os.Setenv("PASSWORD_ACTIVE", "true")
	os.Setenv("PASSWORD", "")
	auth := Setup("", nil)
	err := checkAuthorization(auth.password)
	if err != nil {
		t.Errorf("The checkAuthorization has failed with PASSWORD_ACTIVE false")
	}
}
