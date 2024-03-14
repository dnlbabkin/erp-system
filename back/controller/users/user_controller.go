package users

import (
	"back/domain/users"
	usrs "back/services"
	"back/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"time"
)

const (
	SecretKey = "zhopa"
)

func Register(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	result, err := usrs.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Login(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invalid json")
		c.JSON(err.Status, err)
		return
	}

	result, err := usrs.GetUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	standardClaims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		Issuer:    strconv.Itoa(int(result.UID)),
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, standardClaims)

	token, claimErr := claims.SignedString([]byte(SecretKey))
	if claimErr != nil {
		c.JSON(err.Status, claimErr)
		return
	}

	result.Token = token

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, result)
}

func Get(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		getErr := errors.NewInternalServerError("could not retrieve cookie")
		c.JSON(getErr.Status, getErr)
		return
	}

	standardClaims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(cookie, standardClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		restErr := errors.NewInternalServerError("error parsing cookie")
		c.JSON(restErr.Status, restErr)
		return
	}

	claims := token.Claims.(*jwt.RegisteredClaims)
	issuer, err := strconv.ParseInt(claims.Issuer, 10, 64)
	if err != nil {
		restErr := errors.NewInternalServerError("user id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, restErr := usrs.GetUserByID(issuer)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
