package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	Models "go-micro-service/models"
	"strconv"
)

func GetUsersHandler(ctx *gin.Context) {
	userService := ctx.Keys["userService"].(Models.UserCommonService)
	var req Models.UsersRequest
	size, _ := strconv.ParseInt(ctx.Param("size"), 10, 32)
	req = Models.UsersRequest{Size: int32(size)}
	resp, err := userService.GetUserList(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"status": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"data": resp.Data})
	}
}

func GetUserDetailHandler(ctx *gin.Context) {
	userService := ctx.Keys["userService"].(Models.UserCommonService)
	var req Models.UsersRequest
	userId, _ := strconv.ParseInt(ctx.Param("userId"), 10, 32)
	req = Models.UsersRequest{UserID: int32(userId)}
	resp, err := userService.GetUserDetail(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"status": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"data": resp.Data})
	}
}
