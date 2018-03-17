package dal

import (
	"gawkbox-assignment/environment"
	"gawkbox-assignment/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //For mysql driver
	"log"
)

var (
	db  *gorm.DB
	err error
)

var insertString = "INSERT INTO Users VALUES"
var connectionString = environment.GetMySQLURL() + "?charset=utf8&parseTime=True&loc=Local"

func init() {
	db, err = gorm.Open("mysql", connectionString)
	log.Println("DB  :  ", db)
	if err != nil {
		log.Println("Error", err)
	}
	println(!db.HasTable("Users"))
	defer db.Close()
	if !db.HasTable("Users") {
		db.Table("Users").CreateTable(&models.UserGorm{})
		if err != nil {
			log.Println(err)
		}
	}
}

//InsertOrUpdateUser to update a user or insert new user
func InsertOrUpdateUser(v interface{}) bool {
	db, _ := gorm.Open("mysql", connectionString)
	V := v.(models.User)
	user := models.UserGorm{
		ID:           V.ID,
		Login:        V.Login,
		DisplayName:  V.DisplayName,
		Type:         V.Type,
		Description:  V.Description,
		ProfileImage: V.ProfileImage,
		OfflineImage: V.OfflineImage,
		ViewCount:    V.ViewCount,
		Email:        V.Email,
	}
	if db.Table("Users").Where("id = ?", V.ID).Find(&user).RecordNotFound() {
		db.Table("Users").NewRecord(user)
		db.Table("Users").Create(&user)
		if !db.Table("Users").NewRecord(user) {
			return true
		}
		db.Close()
	}
	return false
}

//FindUsersByDisplayName in the database with display name
func FindUsersByDisplayName(q string) interface{} {
	db, _ := gorm.Open("mysql", connectionString)
	var users []models.UserGorm
	db.Table("Users").Where("login LIKE ?", "%"+q+"%").Find(&users)
	return users
}
