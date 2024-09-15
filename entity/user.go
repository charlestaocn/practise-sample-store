package entity

import (
	"practise-sample-store/config/db"
	"time"
)

type User struct {
	Id       int64     `db:"id" json:"id"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"password" json:"password"`
	Deleted  bool      `db:"deleted" json:"deleted"`
	CreateAt time.Time `gorm:"column:create_at;default:null" json:"createAt"`
}

func (User) TableName() string {
	return "user"
}

func SelectUserByUsername(username string) (User, error) {
	var user User
	db.MySql.Where("username = ? and deleted != 1", username).Find(&user)
	return user, nil
}

func InsertUser(user *User) (bool, error) {
	saveRes := db.MySql.Save(user)
	if saveRes.Error != nil || saveRes.RowsAffected == 0 {
		return false, saveRes.Error
	}
	return true, nil
}
