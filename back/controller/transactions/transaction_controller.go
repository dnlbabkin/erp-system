package transactions

import (
	"back/domain/transactions"
	"back/services"
	"back/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func In(c *gin.Context) {
	var transaction transactions.Transaction

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
