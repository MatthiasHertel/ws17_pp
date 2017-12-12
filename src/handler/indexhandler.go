package handler

import (
	"fmt"
	"net/http"
)

// IndexHandler Handler for /
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!!!")
}
