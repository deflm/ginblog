package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" binding:"required"`
	Password string `gorm:"type:varchar(20);not null " json:"password" binding:"required"`
	Role     int    `gorm:"type:int" json:"role" binding:"required"`
}

// 查询用户
func CheckUser(name string) int {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 创建用户
func CreateUser(user *User) int {
	user.Password = ScryptPw(user.Password)
	err := db.Create(user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := []byte{10, 2, 1, 99, 45, 24, 78, 65}
	dk, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatalln(err)
	}
	return base64.StdEncoding.EncodeToString(dk)
}
