package service

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPing(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Ping)
	handler.ServeHTTP(rr, req)

	// check status code
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Wrong status code: got %v instead of %v", status, http.StatusOK)
	}

	// check body
	body := rr.Body.String()
	log.Println(body)
	if !strings.Contains(body, "pong") {
		t.Errorf("Wrong body")
	}

}
