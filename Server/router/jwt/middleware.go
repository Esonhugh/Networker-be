package jwt

import (
	"Network-be/config"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWT struct {
	Username string
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 24 // 一天过期时间

var MySecret []byte

func Init() {
	MySecret = []byte(config.GlobalConfig.GetString("jwt.secret"))
}

// GenToken 生成token
func GenToken(username string) (string, error) {
	stdClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),  // 过期时间
		Issuer:    config.GlobalConfig.GetString("jwt.issuer"), // 签发人
	}

	c := JWT{
		username,
		stdClaims,
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*JWT, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &JWT{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWT); ok && token.Valid { // 校验token
		return claims, nil
	} else {
		return nil, err
	}
	return nil, errors.New("invalid token")
}
