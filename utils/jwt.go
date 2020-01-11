package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"island/model"
)

// 一些常量
var (
	ErrTokenExpired     error = errors.New("Token已失效")
	ErrTokenNotValidYet error = errors.New("Token未激活")
	ErrTokenMalformed   error = errors.New("Token不合法")
	ErrTokenInvalid     error = errors.New("Token解析失败")
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID       uint   `json:"UserId"`
	UserName string `json:"UserName"`
	Status   int    `json:"Status"`
	Avatar   string `json:"Avatar"`
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(os.Getenv("SIGN_KEY")),
	}
}

// 新建一个customClaims实例
func NewCustomClaims(user *model.User) *CustomClaims {
	return &CustomClaims{
		user.ID,
		user.UserName,
		user.Status,
		user.Avatar,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "island",                      // 签名的发行者
		},
	}
}

// 生成Token
func GenerateToken(user *model.User) (string, error) {
	j := NewJWT()
	claims := NewCustomClaims(user)
	return j.CreateToken(claims)
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims *CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		log.Print(err.Error())
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			}
			return nil, ErrTokenInvalid
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(claims)
	}
	return "", ErrTokenInvalid
}
