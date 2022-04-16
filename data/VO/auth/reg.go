package auth

import "Network-be/data/VO"

type RegRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegResponse VO.CommonResp
