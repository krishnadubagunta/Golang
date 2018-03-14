package twitch

import "net/http"

var streamURL = baseURL + "streams?language=en"

//Stream type of API_Authorization data
type Stream struct {
	ID           string   `json:"id"`
	Type         string   `json:"type"`
	Language     string   `json:"language"`
	GameID       string   `json:"game_id"`
	ThumbnailURL string   `json:"thumbnail_url"`
	Title        string   `json:"title"`
	UserID       string   `json:"user_id"`
	StartedAt    string   `json:"started_at"`
	CommunityIDs []string `json:"community_ids,omitempty"`
}

//Streams strcutre with array
type Streams struct {
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
	Data []Stream `json:"data,omitempty"`
}

//GetStreams from twitch
func GetStreams(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, interface{}) {
	var data Streams
	return requests(w, r, &data, streamURL, "GET")
}
