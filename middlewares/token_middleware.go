package middlewares

import (
	"net/http"
	"os"
)

//TokenMiddleware is used to secure our Gontact endpoint by comparing `X-Gontact-Token` header
//with GONTACT_TOKEN environment variable.
func TokenMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	header := r.Header["X-Gontact-Token"]
	token := os.Getenv("GONTACT_TOKEN")
	cont := true

	if len(header) == 0 || token == "" || header[0] != token {
		rw.WriteHeader(401)
		cont = false
	}

	if !cont {
		if token == "" {
			rw.Write([]byte("You need to setup a GONTACT_TOKEN to secure your server."))
		}

		return
	}

	next(rw, r)
}
