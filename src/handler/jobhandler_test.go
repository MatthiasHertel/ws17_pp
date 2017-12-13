package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestAllJobsEndPoint(t *testing.T) {

	// a.Router = mux.NewRouter()

	req, err := http.NewRequest("GET", "/jobssss/23", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	log.Print(m)

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	log.Print(bodyString)
	// if m["error"] != "not found" {
	// 	t.Errorf("Expected the 'error' key of the response to be set to 'Product not found'. Got '%s'", m["error"])
	// }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(AllJobsEndPoint)
	a := mux.NewRouter()
	a.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
