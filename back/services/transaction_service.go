package services

import (
	"back/domain/transactions"
	"back/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(transaction transactions.Transaction, c *gin.Context) (*transactions.Transaction, *errors.RestErr) {
	if err := transaction.Save(c); err != nil {
		return nil, err
	}

	return &transaction, nil
}
