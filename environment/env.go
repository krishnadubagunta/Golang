package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//Environtment variables
var (
	clientID      string
	clientSectret string
	mysqlURL      string
	redirect_uri  string
)

func init() {
	err := godotenv.Load("./variables.env")
	if err != nil {
		log.Fatalln("Error At Environment Variables, ", err)
	}
	clientID = os.Getenv("TWITCH_CLIENT_ID")
	clientSectret = os.Getenv("TWITCH_CLIENT_SECRET")
	mysqlURL = os.Getenv("MYSQL_DATA_URL")
	redirect_uri = os.Getenv("REDIRECT_URI")
}

//GetTwitchClientID sets passes the Environment Variables to the req handlers
func GetTwitchClientID() string {
	return clientID
}

//GetMySQLURL returns MySQL URL
func GetMySQLURL() string {
	return mysqlURL
}

//GetTwitchClientSecret sets
func GetTwitchClientSecret() string {
	return clientSectret
}

//GetRedirectURI returns the url for callback
func GetRedirectURI() string {
	return redirect_uri
}
