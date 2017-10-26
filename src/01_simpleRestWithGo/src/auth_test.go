package main

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

// basicAuth string helper
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// TODO
func TestServerSendsHeader(t *testing.T) {}

func TestServerDeniesNoCredentials(t *testing.T) {
	handler := BasicAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("SERVER ALLOWS WHAT IT SHOULD DENY %v", r)
	}))

	req, err := http.NewRequest("GET", "/no-credentials", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusForbidden)
	}
}

func TestServerDeniesWrongPassword(t *testing.T) {
	handler := BasicAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("SERVER ALLOWS WHAT IT SHOULD DENY")
	}))

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Basic "+basicAuth("username", "wrong password"))

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusForbidden)
	}
}

func TestServerDeniesWrongUsername(t *testing.T) {
	handler := BasicAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("SERVER ALLOWS WHAT IT SHOULD DENY")
	}))

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Basic "+basicAuth("wrong username", "password"))

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusForbidden)
	}
}

func TestServerAllowsCorrectCredentials(t *testing.T) {

	handler := BasicAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// everything ok
	}))

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Authorization", "Basic "+basicAuth("username", "password"))

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
