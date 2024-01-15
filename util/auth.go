package util

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	// Email string `json:"email"`
	jwt.RegisteredClaims
}

func ValidateToken(ctx *gin.Context, tokenString string) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 	return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		// }
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, "Token is invalid")
		return
	}
	ctx.Next()
}
