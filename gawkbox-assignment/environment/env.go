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
	redirectURI   string
	privateIP     string
	mode          string
)

func init() {
	err := godotenv.Load("./variables.env")
	if err != nil {
		log.Fatalln("Error At Environment Variables, ", err)
	}
	clientID = os.Getenv("TWITCH_CLIENT_ID")
	clientSectret = os.Getenv("TWITCH_CLIENT_SECRET")
	mysqlURL = os.Getenv("MYSQL_DATA_URL")
	redirectURI = os.Getenv("REDIRECT_URI")
	privateIP = os.Getenv("PRIVATE_IP")
	mode = os.Getenv("MODE")
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
	return redirectURI
}

//GetPrivateIP returns the url for callback
func GetPrivateIP() string {
	return privateIP
}

//GetMode return the mode of the application
func GetMode() string {
	return mode
}
