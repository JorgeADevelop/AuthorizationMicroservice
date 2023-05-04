package utils

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SignJWT(userID uint) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iss": os.Getenv("JWT_ISSUER"),
		"exp": time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
		"iat": time.Now().Format("2006-01-02 15:04:05"),
	})

	tokenString, err = token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return
}

func SignRefreshJWT(userID uint) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24 * 3).Format("2006-01-02 15:04:05"),
		"iat": time.Now().Format("2006-01-02 15:04:05"),
	})

	tokenString, err = token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return
}

func IsValidJWT(ctx *gin.Context) (bool, error) {
	tokenString, err := splitToken(ctx.GetHeader("Authoriztion"))
	if err != nil {
		return false, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if token.Valid {
		return true, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return false, jwt.ErrTokenMalformed
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return false, jwt.ErrTokenSignatureInvalid
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return false, jwt.ErrTokenNotValidYet
	}
	return false, err
}

func GetSubClaim(ctx *gin.Context) (uint, error) {
	tokenString, _ := splitToken(ctx.GetHeader("Authoriztion"))
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["sub"].(uint), nil
	}
	return 0, jwt.ErrTokenMalformed
}

func splitToken(tokenStringComplete string) (string, error) {
	tokenStringSplits := strings.Split(tokenStringComplete, " ")

	if len(tokenStringSplits) != 2 {
		return "", jwt.ErrTokenMalformed
	}
	return tokenStringSplits[1], nil
}
