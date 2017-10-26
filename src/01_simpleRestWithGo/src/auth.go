package ws17_pp

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func BasicAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			http.Error(w, "Not authorized", http.StatusForbidden)
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Not authorized", http.StatusForbidden)
			return
		}

		if pair[0] != "username" || pair[1] != "password" {
			http.Error(w, "Not authorized", http.StatusForbidden)
			return
		}

		h.ServeHTTP(w, r)
	})
}
