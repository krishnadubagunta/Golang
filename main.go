package main

import (
	"fmt"
	"gawkbox-assignment/dal"
	"gawkbox-assignment/middleware"
	"gawkbox-assignment/twitchapi"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Booting the server...")
	r := mux.NewRouter()

	// Configure a sample route
	r.HandleFunc("/api/streams", twitchapi.GetStreamsFromTwitch).Methods("GET")
	r.HandleFunc("/api/user/{id}", twitchapi.GetUserByID).Methods("GET")
	dal.DoSomething()
	// Run your server
	log.Fatal(http.ListenAndServe(":8080", middleware.APISetup(r)))
}
