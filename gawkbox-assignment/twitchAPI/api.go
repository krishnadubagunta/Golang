package twitchapi

import "net/http"

//Client type
var client = &http.Client{}

//Base URL
var baseURL = "https://api.twitch.tv/helix/"

//SetClient sets the client to a new client with oauth info
func SetClient(c *http.Client) {
	client = c
}
