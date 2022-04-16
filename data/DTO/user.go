package DTO

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Verify   bool   `json:"verify"`
}

func (u *User) ToVO() {
}
