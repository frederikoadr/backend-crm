package main

import (
	"BackendCRM/dto"
	"BackendCRM/modules/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func initDB() (*gorm.DB, error) {
	dsn := "root:1230@tcp(localhost:3306)/crud?parseTime=true"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func main() {
	var request = dto.Request{
		Body: map[string]string{
			"id": "1",
		},
		Method: "GET",
		Path:   "/get-user",
		Header: map[string]string{
			"Authorization": "token",
		},
	}

	router := user.NewRouter()
	router.Route(request)
}
