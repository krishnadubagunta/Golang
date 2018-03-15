package auth

import (
	"gawkbox-assignment/environment"
	"github.com/dchest/uniuri"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
	"log"
	"net/http"
	"time"
)

//Cookie for storing the access token in sessions
type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string
	MaxAge     int
	Secure     bool
	HTTPOnly   bool
	Raw        string
	Unparsed   []string
}

var (
	twitchOAuth = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/auth/callback",
		ClientID:     environment.GetTwitchClientID(),
		ClientSecret: environment.GetTwitchClientSecret(),
		Scopes:       []string{"user:read:email"},
		Endpoint:     twitch.Endpoint,
	}
	oauthString = "random"
)

func init() {
	log.Println("There initializes the oauth")
	twitch.Endpoint.AuthURL = "https://id.twitch.tv/oauth2/authorize"
}

//OAuthHandler to handle login with Twitch
func OAuthHandler(w http.ResponseWriter, r *http.Request) {
	oauthString = uniuri.New()
	url := twitchOAuth.AuthCodeURL(oauthString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

//OAuthCallBack called when the callback is arrived from the OAuth
func OAuthCallBack(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthString {
		log.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	code := r.FormValue("code")
	token, err := twitchOAuth.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println("Code exchange failed with ?\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	cookie := http.Cookie{
		Name:     "TwitchAccessToken",
		Value:    token.AccessToken,
		Expires:  token.Expiry,
		Domain:   "localhost",
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/api/user", http.StatusPermanentRedirect)
}

//Logout is used to log out of the application
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("TwitchAccessToken")
	cookie.Expires = time.Now()
	http.SetCookie(w, cookie)
}
