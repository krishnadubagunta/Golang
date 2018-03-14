package main

import (
	"fmt"
	"gawkbox-assignment/dal"
	"gawkbox-assignment/middleware"
	"gawkbox-assignment/twitch"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Booting the server...")
	r := mux.NewRouter()

	// Configure a sample route
	r.HandleFunc("/api/streams", twitch.GetStreamsFromTwitch).Methods("GET")
	r.HandleFunc("/api/user/{id}", twitch.GetUserByID).Methods("GET")
	dal.DoSomething()
	// Run your server
	log.Fatal(http.ListenAndServe(":8080", middleware.APISetup(r)))
}
