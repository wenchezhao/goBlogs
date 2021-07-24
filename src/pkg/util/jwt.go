package util

import (
	"github.com/Wenchuan-Zhao/goBlogs/pkg/setting"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

// 用于避免API的随意调用
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	// method对应着SigningMethodHMAC struct{}，其包含SigningMethodHS256、SigningMethodHS384、SigningMethodHS512三种crypto.Hash方案
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	// 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// Valid()验证基于时间的声明exp, iat, nbf，注意如果没有任何声明在令牌中，仍然会被认为是有效的
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
