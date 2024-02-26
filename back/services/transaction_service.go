package services

import (
	"back/domain/transactions"
	"back/utils/errors"
)

func CreateTransaction(transaction transactions.Transaction) (*transactions.Transaction, *errors.RestErr) {
	if err := transaction.Save(); err != nil {
		return nil, err
	}

	return &transaction, nil
}
