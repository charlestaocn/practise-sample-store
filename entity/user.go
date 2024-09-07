package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id       int64     `db:"id" json:"id"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"username" json:"password"`
	Deleted  bool      `db:"deleted" json:"deleted"`
	CreateAt time.Time `db:"create_at" json:"createAt"`
}

func (User) TableName() string {
	return "user"
}

func SelectUserByUsername(mysqlDb *gorm.DB, username string) (User, error) {
	var user User
	mysqlDb.Where("username = ? and deleted != 1", username).Find(&user)
	return user, nil
}
