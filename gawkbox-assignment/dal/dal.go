package dal

import (
	"gawkbox-assignment/environment"
	"log"
	"os"
)

func init() {
	log.Println(os.Getenv("MYSQL_DATA_URL"))
}

//DoSomething does something
func DoSomething() {
	log.Println(environment.GetMySQLURL())
}
