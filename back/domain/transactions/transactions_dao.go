package transactions

import (
	"back/datasource/postgres/transactions_db"
	"back/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

var (
	queryInsertTransaction = "INSERT INTO transactions (user_id, type, name_product, quantity, rate, sum_money, date_time) VALUES ($1, $2, $3, $4, $5, $6, $7)"
)

func (transaction *Transaction) Save(c *gin.Context) *errors.RestErr {
	client := transactions_db.InitDB()

	t := time.Now().Format("2006-01-02")
	transaction.Rate = int(transaction.SumMoney / transaction.Quantity)

	b, err := client.Query(queryInsertTransaction, isUser(c), transaction.TypeTrans, transaction.NameProduct, transaction.Quantity,
		transaction.Rate, transaction.SumMoney, t)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	client.Close()
	b.Close()

	return nil
}

func isUser(c *gin.Context) int64 {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		getErr := errors.NewInternalServerError("could not retrieve cookie")
		c.JSON(getErr.Status, getErr)
	}

	standardClaims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(cookie, standardClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("zhopa"), nil
	})
	if err != nil {
		restErr := errors.NewInternalServerError("error parsing cookie")
		c.JSON(restErr.Status, restErr)
	}

	claims := token.Claims.(*jwt.RegisteredClaims)
	issuer, err := strconv.ParseInt(claims.Issuer, 10, 64)
	if err != nil {
		restErr := errors.NewInternalServerError("user id should be a number")
		c.JSON(restErr.Status, restErr)
	}

	return issuer
}
