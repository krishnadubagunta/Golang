package twitchapi

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/krishnadubagunta/golang/gawkbox-assignment/dal"
	"github.com/krishnadubagunta/golang/gawkbox-assignment/models"
	"net/http"
)

var userURL = baseURL + "users/"

var users models.UserData

//GetUserByID from Twitch API
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	w, user := requests(w, r, &users, userURL+"?id="+id, "GET")
	json.NewEncoder(w).Encode(user)
}

//GetUsers gets the current user
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w, temp := requests(w, r, &users, userURL, "GET")
	users := (temp.(*models.UserData))
	user := users.Data[0]
	if dal.InsertOrUpdateUser(user) {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(user)
}

//GetUsersByName get the users from database by display name sarching as pattern
func GetUsersByName(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	users := dal.FindUsersByDisplayName(q)
	json.NewEncoder(w).Encode(users)
}
