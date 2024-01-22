package api

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	UserID int32  `json:"user_id"`
	Token  string `json:"token"`
}

func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//find the user by Email
	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	//if password is valid, generate the token
	expirationTime := time.Now().Add(time.Minute * 15)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": req.Email,
		"exp":   expirationTime.Unix(),
	})

	//Sign and get the complete encoded token as string
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//send it back and store on cookie
	// ctx.SetSameSite(http.SameSiteLaxMode)
	// ctx.SetCookie("Authorization", tokenString, 3600*30, "", "", false, true)

	response := &loginResponse{
		UserID: user.ID,
		Token:  tokenString,
	}

	ctx.JSON(http.StatusOK, response)
}
