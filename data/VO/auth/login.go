package auth

import "Network-be/data/VO"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse VO.CommonResp
