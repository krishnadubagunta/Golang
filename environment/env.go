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
)

func init() {
	err := godotenv.Load("./variables.env")
	if err != nil {
		log.Fatalln("Error At Environment Variables, ", err)
	}
	clientID = os.Getenv("TWITCH_CLIENT_ID")
	clientSectret = os.Getenv("TWITCH_CLIENT_SECRET")
	mysqlURL = os.Getenv("MYSQL_DATA_URL")
}

//GetTwitchClientCreds sets passes the Environment Variables to the req handlers
func GetTwitchClientCreds() string {
	return clientID
}

//GetMySQLURL returns MySQL URL
func GetMySQLURL() string {
	return mysqlURL
}
