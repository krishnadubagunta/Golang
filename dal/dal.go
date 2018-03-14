package dal

import (
	"database/sql"
	"gawkbox-assignment/environment"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", environment.GetMySQLURL())
	if err != nil {
		log.Println(err)
	}
	if db == nil {
		_, err := db.Exec("create database twitch")
		if err != nil {
			log.Println(err)
		}

	}
}

//DoSomething does something
func DoSomething() {
	log.Println(db)
}

//Unexported loadData module
func loadData() {

}
