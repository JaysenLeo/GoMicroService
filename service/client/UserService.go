package service

import (
	"go-micro-service/models"
	"log"
	"strconv"
)

type UserService struct {
}

func NewUser(id int32, name string) *Models.UserModel {
	return &Models.UserModel{UserID: id, Name: name}
}

func GetDataDefault(Resp interface{}) {
	switch t := Resp.(type) {
	case *Models.UserListResponse:
		if err := GetUserListDefault(Resp); err != nil {
			log.Fatal(err.Error())
		}
	case *Models.UserDetailResponse:
		t.Data = NewUser(999, "Leee")
		//GetUserDetailDefault(Resp)
	}
}

// 获取用户列表 降级方法
func GetUserListDefault(UserListResp interface{}) error {
	userListResp := UserListResp.(*Models.UserListResponse)
	ret := make([]*Models.UserModel, 0)
	for i := int32(0); i < 10; i++ {
		ret = append(ret, NewUser(i, "user_default"+strconv.FormatInt(int64(i), 10)))
	}
	userListResp.Data = ret
	return nil
}

// 获取用户详情信息 降级方法
func GetUserDetailDefault(UserDetailResp interface{}) error {
	userDetailResp := UserDetailResp.(*Models.UserDetailResponse)
	userDetailResp.Data = NewUser(999, "Leee")
	return nil
}
