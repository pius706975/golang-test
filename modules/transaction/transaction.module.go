package transaction

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransactionModule(router *gin.Engine, db *gorm.DB, prefix string) {
	userRepo := NewRepo(db)
	userService := NewService(userRepo)
	userController := NewController(userService)

	TransactionRoutes(router, userController, prefix)
}