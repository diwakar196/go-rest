package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-rest/model"
	"go-rest/repository"
	"go-rest/service"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewRepository(db)
	service := service.NewService(repo)

	r.POST("/", func(c *gin.Context) { CreateUser(c, service) })

	r.GET("/", func(c *gin.Context) { GetAllUsers(c, service) })

}

func GetAllUsers(c *gin.Context, service *service.Service) {
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context, service *service.Service) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
