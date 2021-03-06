package middleware

import (
	"github.com/krishnadubagunta/golang/gawkbox-assignment/environment"
	"net/http"
)

var clientID string

//APISetup sets the authorization data
func APISetup(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Connection") != "" {
			r.Header.Del("Connection")
		}
		r.Header.Set("Client-ID", environment.GetTwitchClientID())
		cookie, err := r.Cookie("TwitchAccessToken")
		if err == nil && cookie != nil {
			r.Header.Set("Authorization", "Bearer "+cookie.Value)
		}
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
