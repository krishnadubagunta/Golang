package twitch

import (
	"net/http"
)

var userURL = baseURL + "users?id="

//User type for rendering the user Model
type User struct {
	ID           string `json:"id"`
	Login        string `json:"login"`
	DisplayName  string `json:"display_name"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	ProfileImage string `json:"profile_image_url"`
	ViewCount    int64  `json:"view_count"`
	Email        string `json:"email,omitempty"`
}

//UserData to User Struct
type UserData struct {
	Data [1]User `json:"data"`
}

//GetUser function to ger a user by ID
func GetUser(w http.ResponseWriter, r *http.Request, id string) (http.ResponseWriter, interface{}) {
	var user UserData
	return requests(w, r, &user, userURL+id, "GET")
}
