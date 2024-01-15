package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(ctx *gin.Context, tokenString string) (c *gin.Context) {

	//if you are using cookie
	// tokenString, err := ctx.Cookie("Authorization")
	// if err != nil {
	// 	ctx.AbortWithStatus(http.StatusUnauthorized)
	// 	return
	// }

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
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
	return
}

func GetTokenHeaderAndValidate(ctx *gin.Context) {
	authorizationHeaderKey := ctx.GetHeader("authorization")
	fields := strings.Fields(authorizationHeaderKey)
	tokenToValidate := fields[1]
	errOnValidateToken := ValidateToken(ctx, tokenToValidate)
	if errOnValidateToken != nil {
		ctx.JSON(http.StatusUnauthorized, "Token is invalid")
		return
	}
	ctx.Next()
}
