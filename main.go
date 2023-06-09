package main

import (
	"BackendCRM/modules/route"
	"BackendCRM/utility/db"
	"github.com/gin-gonic/gin"
	"log"
)

var err error

func main() {
	dsn := "root:1230@tcp(db:3306)/crm_service?parseTime=true"
	mysqlDb, err := db.InitDB(dsn)
	if err != nil {
		log.Fatalln("initDB:", err)
	}

	r := gin.Default()
	route.Router(r, mysqlDb)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	// request handler: menerima request, mengirim response
	// controller: validasi dan transformasi data
	// use case: pemrosesan data
	// repository: persistensi data
}
