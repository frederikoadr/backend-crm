package main

import (
	"BackendCRM/modules/users"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Customers struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

var db *gorm.DB
var err error

func initDB() (*gorm.DB, error) {
	dsn := "root:1230@tcp(localhost:3306)/crm_service?parseTime=true"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func getUserById(c *gin.Context) {
	var user Customers
	userID := c.Param("id")

	// Dapatkan data user dari database berdasarkan ID
	if err := db.First(&user, userID).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customers not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tampilkan data user
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalln("initDB:", err)
	}

	r := gin.Default()
	usersHandler := users.DefaultRequestHandler(db)

	r.POST("/customers", usersHandler.Create)
	r.GET("/customers", usersHandler.Read)
	r.GET("/customers/:id", usersHandler.Read)
	r.DELETE("/customers/:id", usersHandler.Delete)
	r.PUT("customers/:id", usersHandler.Update)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	// request handler: menerima request, mengirim response
	// controller: validasi dan transformasi data
	// use case: pemrosesan data
	// repository: persistensi data
}
