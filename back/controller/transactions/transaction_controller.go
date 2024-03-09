package transactions

import (
	"back/domain/transactions"
	"back/services"
	"back/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func In(c *gin.Context) {
	//TODO разобраться с куками
	var transaction transactions.Transaction
	cookie, err := c.Cookie("jwt")
	if err != nil {
		getErr := errors.NewInternalServerError("could not retrieve cookie")
		c.JSON(getErr.Status, getErr)
		return
	}

	standardClaims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(cookie, standardClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("zhopa"), nil
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

	res, restErr := services.GetUserByID(issuer)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := c.ShouldBindJSON(&transaction); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	result, err := services.CreateTransaction(transaction)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
