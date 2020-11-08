package service

import (
	"context"
	"go-micro-service/models"
	"strconv"
)

type UserService struct {
}

func NewUser(id int32, name string) *Models.UserModel {
	return &Models.UserModel{UserID: id, Name: name}
}

func NewUserList(size int32) []*Models.UserModel {
	ret := make([]*Models.UserModel, 0)
	for i := int32(0); i < size; i++ {
		ret = append(ret, NewUser(i, "user"+strconv.FormatInt(int64(i), 10)))
	}
	return ret
}

func (*UserService) GetUserList(ctx context.Context, UserListReq *Models.UsersRequest, UserListResp *Models.UserListResponse) error {
	UserListResp.Data = NewUserList(UserListReq.Size)
	return nil
}
