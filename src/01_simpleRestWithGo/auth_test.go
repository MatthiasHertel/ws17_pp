package main

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func TestServerDeniesWrongPassword(t *testing.T) {
	handler := BasicAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("SERVER ALLOWS WHAT IT SHOULD DENY")
	}))

	req, err := http.NewRequest("GET", "/wrong-password", nil)
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

// func TestServerDeniesWrongPassword(t *testing.T) {
// 	handler := BasicAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		t.Fatalf("SERVER ALLOWS WHAT IT SHOULD DENY")
// 	}))
//
// 	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
// 	// pass 'nil' as the third parameter.
// 	req, err := http.NewRequest("GET", "/health-check", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
// 	rr := httptest.NewRecorder()
// 	// handler := http.HandlerFunc(HealthCheckHandler)
//
// 	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
// 	// directly and pass in our Request and ResponseRecorder.
// 	handler.ServeHTTP(rr, req)
//
// 	// Check the status code is what we expect.
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
//
// 	// Check the response body is what we expect.
// 	expected := `{"alive": true}`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
//
// }
func TestServerDeniesWrongUsername(t *testing.T) {
	// auth := BasicAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	t.Fatalf("SERVER ALLOWS WHAT IT SHOULD DENY")
	// }))
}
func TestServerAllowsCorrectCredentials(t *testing.T) {

	handler := BasicAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// everything ok
	}))

	req, err := http.NewRequest("GET", "/wrong-password", nil)
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
