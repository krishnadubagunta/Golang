package main

import (
	"fmt"
	"gawkbox-assignment/auth"
	"gawkbox-assignment/environment"
	"gawkbox-assignment/middleware"
	"gawkbox-assignment/twitchapi"
	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/twitch"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Booting the server...")
	r := mux.NewRouter()
	goth.UseProviders(
		twitch.New(environment.GetTwitchClientID(), environment.GetTwitchClientSecret(), "http://localhost:3000/api/auth/twitch/callback"),
	)
	// Configure a sample route
	r.HandleFunc("/api/streams", twitchapi.GetStreams).Methods("GET")
	r.HandleFunc("/api/streams/{page}", twitchapi.GetStreamsWithPagination).Methods("GET")
	r.HandleFunc("/api/user/{id}", twitchapi.GetUserByID).Methods("GET")
	r.HandleFunc("/api/user", twitchapi.GetUsers).Methods("GET")
	r.HandleFunc("/api/search/user", twitchapi.GetUsersByName).Methods("GET")
	r.HandleFunc("/api/auth", auth.OAuthHandler).Methods("GET")
	r.HandleFunc("/api/auth/callback", auth.OAuthCallBack).Methods("GET")

	// Run your server
	log.Fatal(http.ListenAndServe(":8080", middleware.APISetup(r)))
}
