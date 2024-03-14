package transactions

import (
	"back/domain/transactions"
	transactions2 "back/services"
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

	result, err := transactions2.CreateTransaction(transaction, c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
