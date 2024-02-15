package services

import (
	"back/domain/users"
	"back/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.NewBadRequestError("failed to encrypt the password")
	}

	user.Password = string(pwSlice[:])

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(user users.User) (*users.User, *errors.RestErr) {
	result := &users.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		ThirdName: user.ThirdName,
	}

	if err := result.GetByLastName(); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
		return nil, errors.NewBadRequestError("failed to encrypt the password")
	}

	resultWp := &users.User{
		ID:        result.ID,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		ThirdName: result.ThirdName,
	}

	return resultWp, nil
}

func GetUserByID(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userId}

	if err := result.GetByID(); err != nil {
		return nil, err
	}

	return result, nil
}
