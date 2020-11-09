package web

import (
	"github.com/gin-gonic/gin"
	Models "go-micro-service/models"
	"go-micro-service/web/handlers"
	"go-micro-service/web/middlewares"
)

func NewRouter(userListService Models.UserListService) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middlewares.UserMiddleware(userListService))
	ginRouter.Handle("GET", "/users/:size", handlers.GetUsersHandler)
	return ginRouter
}
