package main

import (
	"BackendCRM/modules/account"
	"BackendCRM/modules/customers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var err error

func initDB() (*gorm.DB, error) {
	dsn := "root:1230@tcp(localhost:3306)/crm_service?parseTime=true"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalln("initDB:", err)
	}

	r := gin.Default()
	customerHandler := customers.DefaultRequestHandler(db)

	r.POST("/customers", customerHandler.Create)
	r.GET("/customers", customerHandler.Read)
	r.GET("/customers/:column", customerHandler.ReadBy)
	r.DELETE("/customers/:id", customerHandler.Delete)
	r.PUT("customers/:id", customerHandler.Update)

	accountHandler := account.DefaultRequestHandler(db)
	r.GET("/login", accountHandler.Login)
	r.POST("/actors", accountHandler.Create)
	r.GET("/actors", accountHandler.Read)
	r.DELETE("/actors/:id", accountHandler.Delete)
	r.PUT("/actors/:id", accountHandler.Update)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	// request handler: menerima request, mengirim response
	// controller: validasi dan transformasi data
	// use case: pemrosesan data
	// repository: persistensi data
}
