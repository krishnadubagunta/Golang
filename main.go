package main

import (
	"fmt"
	"gawkbox-assignment/auth"
	"gawkbox-assignment/middleware"
	"gawkbox-assignment/twitchapi"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// Configure a sample route
	r.HandleFunc("/api/streams", twitchapi.GetStreams).Methods("GET")
	r.HandleFunc("/api/streams/{page}", twitchapi.GetStreamsWithPagination).Methods("GET")
	r.HandleFunc("/api/user", twitchapi.GetUsers).Methods("GET")
	r.HandleFunc("/api/user/{id}", twitchapi.GetUserByID).Methods("GET")
	r.HandleFunc("/api/search/user", twitchapi.GetUsersByName).Methods("GET")
	r.HandleFunc("/api/auth", auth.OAuthHandler).Methods("GET")
	r.HandleFunc("/api/auth/callback", auth.OAuthCallBack).Methods("GET")
	r.HandleFunc("/api/logout", auth.Logout).Methods("GET")

	fmt.Println("Server Booted...")
	// Run your server
	log.Fatal(http.ListenAndServe(":8080", middleware.APISetup(r)))
}
