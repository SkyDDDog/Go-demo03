package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	//除了满足下面的Claims，还需要以下用户信息
	Username string `json:"username"`
	Password string `json:"password"`
	//jwt中标准的Claims
	jwt.StandardClaims
}

// 使用指定的 secret 签名声明一个 key ，便于后续获得完整的编码后的字符串token
var key = []byte("west2")

const TokenExpireDuration = time.Hour * 24 * 365
const Issuer = "MrHuang"

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return key, nil
}

// GenToken 生成access token 和 refresh token
func GenToken(username string, password string) (aToken, rToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, // 自定义字段
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    Issuer,                                     // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(key)

	// refresh token 不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(), // 过期时间
		Issuer:    Issuer,                                  // 签发人
	}).SignedString(key)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return
}

// ParseToken 解析token的方法
func ParseToken(tokenString string) (*MyClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return key, nil
		})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { //校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {

	// 从旧access token中解析出claims数据	解析出payload负载信息
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)

	// 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token

	return GenToken(claims.Username, claims.Password)
}

func TokenValid(tokenString string) bool {
	claim, _ := ParseToken(tokenString)
	return claim != nil
}
