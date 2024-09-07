package service

import (
	"practise-sample-store/config/cache"
	"practise-sample-store/config/db"
	"practise-sample-store/entity"
)

var UserTokenCache cache.Cache

func UserLogin(username string, password string) (bool, error) {
	userByUsername, err := entity.SelectUserByUsername(db.MySql, username)
	if err != nil {
		return false, err
	}
	if userByUsername.Password == password {
		return true, nil
	}
	return false, nil
}

func GetUserInfo(username string) (entity.User, error) {
	userByUsername, err := entity.SelectUserByUsername(db.MySql, username)
	return userByUsername, err
}
