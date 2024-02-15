package users

import (
	"back/datasource/postgres/users_db"
	"back/utils/errors"
)

var (
	//queryInsertUser        = "INSERT INTO users (first_name, last_name, third_name, password) VALUES (?, ?, ?, ?);"
	queryInsertUser        = "insert into users (first_name, last_name, third_name, password) values ($1, $2, $3, $4)"
	queryGetUserByLastName = "SELECT id, first_name, last_name, third_name, password FROM users WHERE last_name=?;"
	queryGetUserByID       = "SELECT id, first_name, last_name, third_name FROM users WHERE id=?;"
)

func (user *User) Save() *errors.RestErr {
	//stmt, err := users_db.Client.Prepare(queryInsertUser)
	//if err != nil {
	//	return errors.NewInternalServerError("database error")
	//}

	stmt := users_db.Client.QueryRow(queryInsertUser, user.FirstName, user.LastName, user.ThirdName, user.Password)
	err := stmt.Scan(&user.ID, &user.FirstName, &user.LastName, &user.ThirdName, &user.Password)
	if err != nil {
		panic(err)
	}
	//defer stmt.Close()
	//
	//insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.ThirdName, user.Password)
	//if saveErr != nil {
	//	return errors.NewInternalServerError("database error")
	//}

	//userID, err := insertResult.LastInsertId()
	//if err != nil {
	//	return errors.NewInternalServerError("database error")
	//}

	//user.ID = userID

	return nil
}

func (user *User) GetByLastName() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserByLastName)
	if err != nil {
		return errors.NewInternalServerError("invalid last name")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.LastName)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.ThirdName, &user.Password); err != nil {
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) GetByID() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserByID)
	if err != nil {
		return errors.NewInternalServerError("invalid id")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.LastName)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.ThirdName); err != nil {
		return errors.NewInternalServerError("database error")
	}

	return nil
}
