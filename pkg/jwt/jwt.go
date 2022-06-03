package jwt

import (
	"douyin/global"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(global.Secrete)

//token持续时间设置为1天
const TokenExpireDuration = time.Hour * 24

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"name"`
	jwt.StandardClaims
}

// GetToken 生成Token
func GetToken(userid int64, username string) (string, error) {
	//创建一个自己声明的数据
	c := MyClaims{
		UserID:   userid,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "douyin",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名获得完整的编码后字符串token
	return token.SignedString(jwtSecret)
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*MyClaims, error) {
	//解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { //校验Token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
