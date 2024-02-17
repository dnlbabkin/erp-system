package services

import (
	"back/domain/users"
	"back/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(user users.User) (*users.User, *errors.RestErr) {
	result := &users.User{
		Name:      user.Name,
		LastName:  user.LastName,
		ThirdName: user.ThirdName,
	}

	if err := result.GetByNames(); err != nil {
		return nil, err
	}

	resultWp := &users.User{
		ID:        result.ID,
		Name:      result.Name,
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
