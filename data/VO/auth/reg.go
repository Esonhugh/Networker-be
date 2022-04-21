package auth

import "Network-be/data/VO"

type RegRequest struct {
	Username string `json:"username" binding:"required,min=4,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
	Email    string `json:"email" binding:"required,email"`
}

type RegResponse VO.CommonResp
