package middlewares

import (
	"net/http"
	"strings"
)

//CORSMiddleware is a middleware we added to allow CORS against Gontact
func CORSMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	origins := strings.Join([]string{"*"}, ", ")
	methods := strings.Join([]string{"GET", "POST"}, ", ")
	headers := strings.Join([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"}, ", ")

	rw.Header().Set("Access-Control-Allow-Origin", origins)
	rw.Header().Set("Access-Control-Allow-Methods", methods)
	rw.Header().Set("Access-Control-Allow-Headers", headers)

	if r.Method == "OPTIONS" {
		rw.WriteHeader(200)
		return
	}

	next(rw, r)
}
