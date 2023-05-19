package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type SignupReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Signup(c *gin.Context) {
	var signupReq SignupReq
	if err := c.ShouldBind(&signupReq); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(15 * time.Minute).Unix()
	claims["username"] = signupReq.Username

	// TODO: avoid hard-coding the key here
	tokenString, err := token.SignedString([]byte("abcd1234!"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, tokenString)
}
