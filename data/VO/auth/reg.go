package auth

import "Network-be/data/VO"

type RegRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegResponse VO.CommonResp
