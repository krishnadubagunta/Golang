package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/krishnadubagunta/golang/gawkbox-assignment/auth"
	"github.com/krishnadubagunta/golang/gawkbox-assignment/middleware"
	"github.com/krishnadubagunta/golang/gawkbox-assignment/twitchapi"
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
