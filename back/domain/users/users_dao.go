package users

import (
	"back/datasource/postgres/users_db"
	"back/utils/errors"
	"github.com/google/uuid"
)

var (
	queryInsertUser                = "INSERT INTO users (uid, name, second_name, third_name, password) VALUES ($1, $2, $3, $4, $5)"
	queryGetUserByUIDAndByPassword = "SELECT uid, name, second_name, third_name, password FROM users WHERE uid=$1 AND password=$2"
	queryGetUserByID               = "SELECT uid, name, second_name, third_name FROM users WHERE uid=$1"
	queryGetUid                    = "SELECT uid FROM users"
)

func (user *User) Save() *errors.RestErr {
	client := users_db.InitDB()

	id := uuid.New()

	b, err := client.Query(queryInsertUser, id.ID(), user.Name, user.LastName, user.ThirdName, user.Password)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	user.UID = int64(id.ID())

	client.Close()
	b.Close()

	return nil
}

func (user *User) GetByUIDAndByPassword() *errors.RestErr {
	client := users_db.InitDB()

	res := client.QueryRow(queryGetUserByUIDAndByPassword, user.UID, user.Password)
	if err := res.Scan(&user.UID, &user.Name, &user.LastName, &user.ThirdName, &user.Password); err != nil {
		return errors.NewInternalServerError("database error")
	}

	client.Close()

	return nil
}

func (user *User) GetByID() *errors.RestErr {
	client := users_db.InitDB()

	res := client.QueryRow(queryGetUserByID, user.UID)
	if err := res.Scan(&user.UID, &user.Name, &user.LastName, &user.ThirdName); err != nil {
		return errors.NewInternalServerError("database error")
	}

	client.Close()

	return nil
}
