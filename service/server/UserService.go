package service

import (
	"context"
	"fmt"
	"go-micro-service/models"
	"strconv"
	"time"
)

type UserService struct {
}

func NewUser(id int32, name string) *Models.UserModel {
	return &Models.UserModel{UserID: id, Name: name}
}

func NewUserDetail(id int32) *Models.UserModel {
	// Sex 1 男  0 女
	// Role user 普通用户  admin 管理端用户
	return &Models.UserModel{
		UserID: id, Name: "Lee" + strconv.FormatInt(int64(id), 10),
		Sex: "1", Address: "地址" + strconv.FormatInt(int64(id), 10), Role: "user",
	}
}

func NewUserList(size int32) []*Models.UserModel {
	ret := make([]*Models.UserModel, 0)
	for i := int32(0); i < size; i++ {
		ret = append(ret, NewUser(i, "user"+strconv.FormatInt(int64(i), 10)))
	}
	return ret
}

func (*UserService) GetUserList(ctx context.Context, UserListReq *Models.UsersRequest, UserListResp *Models.UserListResponse) error {
	fmt.Println("GetUserList")
	// 模拟触发降级
	time.Sleep(time.Second * 3)
	UserListResp.Data = NewUserList(UserListReq.Size)
	return nil
}

func (*UserService) GetUserDetail(ctx context.Context, UserDetailReq *Models.UsersRequest, UserDetailResp *Models.UserDetailResponse) error {
	fmt.Println("GetUserDetail")
	// 模拟触发降级
	time.Sleep(time.Second * 3)
	UserDetailResp.Data = NewUserDetail(UserDetailReq.UserID)
	return nil
}
