package middlewares

import (
	"github.com/gin-gonic/gin"
	Models "go-micro-service/models"
)

func UserMiddleware(userListService Models.UserListService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Keys = make(map[string]interface{})
		ctx.Keys["userService"] = userListService
		ctx.Next()
	}

}
