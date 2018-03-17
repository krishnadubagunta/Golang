package twitchapi

import (
	"encoding/json"
	"gawkbox-assignment/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var streamURL = baseURL + "streams?language=en"

//GetStreams from twitch
func GetStreams(w http.ResponseWriter, r *http.Request) {
	var data models.Streams
	w, streams := requests(w, r, &data, streamURL, "GET")
	log.Println("Response : ", streams)
	json.NewEncoder(w).Encode(streams)
}

//GetStreamsWithPagination gets after a certain page ID
func GetStreamsWithPagination(w http.ResponseWriter, r *http.Request) {
	var data models.Streams
	params := mux.Vars(r)
	page := params["page"]
	w, streams := requests(w, r, &data, streamURL+"?after="+page, "GET")
	json.NewEncoder(w).Encode(streams)
}
