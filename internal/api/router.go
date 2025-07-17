package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-rest/handler"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userGroup := router.Group("/users")
	handler.RegisterRoutes(userGroup, db)

	return route
}
