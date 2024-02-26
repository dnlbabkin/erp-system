package transactions

import (
	"back/datasource/postgres/transactions_db"
	"back/domain/users"
	"back/utils/errors"
	"time"
)

var (
	queryInsertTransaction = "INSERT INTO transactions (user_id, type, name_product, quantity, rate, sum_money, date_time) VALUES ($1, $2, $3, $4, $5, $6, $7)"
)

func (transaction *Transaction) Save() *errors.RestErr {
	client := transactions_db.InitDB()
	var user users.User

	t := time.Now().Format("2006-01-02")
	transaction.Rate = int(transaction.SumMoney / transaction.Quantity)

	b, err := client.Query(queryInsertTransaction, user.UID, transaction.TypeTrans, transaction.NameProduct, transaction.Quantity,
		transaction.Rate, transaction.SumMoney, t)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	client.Close()
	b.Close()

	return nil
}
