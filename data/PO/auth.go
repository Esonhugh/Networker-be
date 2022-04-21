package PO

import (
	"Network-be/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type Auth struct {
	gorm.Model
	ID       int64  `gorm:"primary_key,AUTO_INCREMENT"`
	Username string `gorm:"unique;not null"`
	Password []byte // password with Hashed
	Email    string
	Verify   bool
}

func (auth *Auth) TableName() string {
	return "auth"
}

func (auth *Auth) SetPassword(password string) {
	auth.Password = utils.HASH(password)
}

// CheckPassword 检查确认密码
func (auth *Auth) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(auth.Password, []byte(password))
	if err != nil {
		log.Println("Password Compare failed:", err)
		return false
	}
	return true
}

/*
func (auth *Auth) ToDto() *DTO.User {
	return &DTO.User{
		ID:       auth.ID,
		Username: auth.Username,
		Email:    auth.Email,
		Verify:   auth.Verify,
	}
}
*/
