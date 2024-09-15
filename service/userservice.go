package service

import (
	"github.com/gookit/validate"
	"practise-sample-store/config/cache"
	"practise-sample-store/entity"
	"practise-sample-store/util"
)

var UserTokenCache cache.Cache

func UserLoginService(username string, password string) (bool, error) {
	userByUsername, err := entity.SelectUserByUsername(username)
	if err != nil {
		return false, err
	}
	if userByUsername.Password == password {
		return true, nil
	}
	return false, nil
}

func GetUserInfoService(username string) (entity.User, error) {
	userByUsername, err := entity.SelectUserByUsername(username)
	return userByUsername, err
}

func RegisterUserService(user *entity.User) (bool, string) {
	if validate.IsEmpty(user.Username) || validate.IsEmpty(user.Password) {
		return false, "参数不正确"
	}

	userInfo, err := GetUserInfoService(user.Username)
	if err != nil || !util.CheckTypeByReflectZero(userInfo) {
		return false, "用户以存在!"
	}

	insertRes, err := entity.InsertUser(user)
	if err != nil || !insertRes {
		return true, "注册失败! "
	}
	return true, "注册成功! "

}
