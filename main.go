package main

import (
	"BackendCRM/modules/route"
	"BackendCRM/utility/db"
	"github.com/gin-gonic/gin"
	"log"
)

var err error

//	func initDB() (*gorm.DB, error) {
//		dsn := "root:1230@tcp(localhost:3306)/crm_service?parseTime=true"
//		return gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	}
func main() {
	dsn := "root:1230@tcp(localhost:3306)/crm_service?parseTime=true"
	mysqlDb, err := db.InitDB(dsn)
	if err != nil {
		log.Fatalln("initDB:", err)
	}

	r := gin.Default()
	route.Router(r, mysqlDb)
	//customerHandler := customers.DefaultRequestHandler(db)
	//r.POST("/customers", customerHandler.Create)
	//r.GET("/customers", customerHandler.Read)
	//r.GET("/customers/:column", customerHandler.ReadBy)
	//r.DELETE("/customers/:id", customerHandler.Delete)
	//r.PUT("customers/:id", customerHandler.Update)
	//
	//accountHandler := account.DefaultRequestHandler(db)
	//r.POST("/login", accountHandler.Login)
	//r.POST("/actors", accountHandler.Create)
	//r.POST("/register/:id", accountHandler.CreateReg)
	//r.PUT("/register/:id", accountHandler.UpdateReg)
	//r.GET("/actors", accountHandler.Read)
	//r.GET("/registers", accountHandler.ReadRegis)
	//r.DELETE("/actors/:id", accountHandler.Delete)
	//r.PUT("/actors/:id", accountHandler.Update)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	// request handler: menerima request, mengirim response
	// controller: validasi dan transformasi data
	// use case: pemrosesan data
	// repository: persistensi data
}
