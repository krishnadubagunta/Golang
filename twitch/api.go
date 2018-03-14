package twitch

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Client type
var client = &http.Client{}

//Base URL
var baseURL = "https://api.twitch.tv/helix/"

func init() {
	fmt.Println("Initializing Twitch API...")
}

//GetStreamsFromTwitch from Twitch API
func GetStreamsFromTwitch(w http.ResponseWriter, r *http.Request) {
	w, data := GetStreams(w, r)
	json.NewEncoder(w).Encode(data)
}

//GetUserByID from Twitch API
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	w, user := GetUser(w, r, id)
	json.NewEncoder(w).Encode(user)
}
