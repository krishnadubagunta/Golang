package middleware

import (
	"gawkbox-assignment/environment"
	"net/http"
)

var clientID string

//APISetup sets the authorization data
func APISetup(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Client-ID", environment.GetTwitchClientCreds())
		next.ServeHTTP(w, r)
	})
}
