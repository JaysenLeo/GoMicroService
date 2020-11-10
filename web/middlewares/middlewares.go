package middlewares

import (
	"github.com/gin-gonic/gin"
	Models "go-micro-service/models"
)

func UserMiddleware(userService Models.UserCommonService) gin.HandlerFunc {
	// 此处传入 rpc 业务服务
	return func(ctx *gin.Context) {
		ctx.Keys = make(map[string]interface{})
		ctx.Keys["userService"] = userService
		ctx.Next()
	}

}
