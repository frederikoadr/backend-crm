package route

import (
	"BackendCRM/modules/account"
	"BackendCRM/modules/customers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(engine *gin.Engine, db *gorm.DB) {
	customerHandler := customers.DefaultRequestHandler(db)
	engine.POST("/customers", customerHandler.Create)
	engine.GET("/customers", customerHandler.Read)
	engine.GET("/customers/:column", customerHandler.ReadBy)
	engine.DELETE("/customers/:id", customerHandler.Delete)
	engine.PUT("customers/:id", customerHandler.Update)

	accountHandler := account.DefaultRequestHandler(db)
	engine.POST("/login", accountHandler.Login)
	engine.POST("/actors", accountHandler.Create)
	engine.POST("/register/:id", accountHandler.CreateReg)
	engine.PUT("/register/:id", accountHandler.UpdateReg)
	engine.GET("/actors", accountHandler.Read)
	engine.GET("/registers", accountHandler.ReadRegis)
	engine.DELETE("/actors/:id", accountHandler.Delete)
	engine.PUT("/actors/:id", accountHandler.Update)
}
