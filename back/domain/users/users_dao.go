package users

import (
	"back/datasource/postgres/users_db"
	"back/utils/errors"
)

var (
	queryInsertUser     = "INSERT INTO users (name, second_name, third_name, password) VALUES ($1, $2, $3, $4)"
	queryGetUserByNames = "SELECT id, first_name, last_name, third_name, password FROM users WHERE (name=?) AND (last_name=?) AND (third_name=?);"
	queryGetUserByID    = "SELECT id, first_name, last_name, third_name FROM users WHERE id=?;"
)

func (user *User) Save() *errors.RestErr {
	client := users_db.InitDB()

	b, err := client.Query(queryInsertUser, user.Name, user.LastName, user.ThirdName, user.Password)
	if err != nil {
		panic(err)
	}

	client.Close()
	b.Close()

	return nil
}

func (user *User) GetByNames() *errors.RestErr {
	client := users_db.InitDB()

	client.QueryRow(queryGetUserByNames, user.Name, user.LastName, user.ThirdName).
		Scan(&user.ID, &user.Name, &user.LastName, &user.ThirdName, &user.Password)

	client.Close()

	return nil
}

func (user *User) GetByID() *errors.RestErr {
	client := users_db.InitDB()

	client.QueryRow(queryGetUserByID, user.ID).
		Scan(&user.ID, &user.Name, &user.LastName, &user.ThirdName)

	client.Close()

	return nil
}
