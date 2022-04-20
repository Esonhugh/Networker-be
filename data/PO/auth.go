package PO

import (
	"Network-be/utils"
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	ID       int64    `gorm:"primary_key"`
	Username string   `gorm:"unqiue;not null"`
	Password [16]byte // password with md5
	Email    string
	Verify   bool
}

func (auth *Auth) TableName() string {
	return "auth"
}

func (auth *Auth) SetPassword(password string) {
	auth.Password = utils.MD5(password)
}

// CheckPassword 检查确认密码
func (auth *Auth) CheckPassword(password string) bool {
	return utils.MD5(password) == auth.Password
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
