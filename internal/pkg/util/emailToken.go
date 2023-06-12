package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

// EmailClaims token中携带的信息
type EmailClaims struct {
	ID                   uint   `json:"id"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	OperationType        uint   `json:"operationType"`
	jwt.RegisteredClaims        // jwt预定义结构体
}

// GenerateEmailToken 签发email token
func GenerateEmailToken(ID uint, Email string, Password string, OperationType uint) (string, error) {
	// 签发token时间
	nowTime := time.Now()
	// token失效时间
	expireTime := nowTime.Add(24 * time.Hour)
	// token中携带的信息
	claims := &EmailClaims{
		ID:            ID,
		Email:         Email,
		Password:      Password,
		OperationType: OperationType,
		RegisteredClaims: jwt.RegisteredClaims{
			// token签发者
			Issuer: viper.GetString("server.domain"),
			// token主题
			Subject: "email token",
			// token有效期
			ExpiresAt: jwt.NewNumericDate(expireTime),
			// token签发时间
			IssuedAt: jwt.NewNumericDate(nowTime),
		},
	}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 设置签名密钥
	return token.SignedString(jwtSignKey)
}

// ParseEmailToken 解析email token
func ParseEmailToken(tokenStr string) (*EmailClaims, error) {
	// ParseWithClaims接收第一个值是token
	// 第二个值是解析后的数据存放的claims
	// 第三个值是keyFunc将被Parse方法作为回调函数，提供用于验证的密钥，函数接收已解析但未验证的token
	token, err := jwt.ParseWithClaims(tokenStr, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})
	// 若err不为空
	if err != nil {
		return nil, err
	}
	// 校验token
	if claims, ok := token.Claims.(*EmailClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
