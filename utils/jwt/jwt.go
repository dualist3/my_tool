package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

func GetJwtToken(secretKey string, iat, seconds, UserType int64, UserId uint, openId, sessionKey string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["openId"] = openId
	claims["userId"] = UserId
	claims["userType"] = UserType
	claims["sessionKey"] = sessionKey
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
