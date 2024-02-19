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
		UID:      user.UID,
		Password: user.Password,
	}

	if err := result.GetByUIDAndByPassword(); err != nil {
		return nil, err
	}

	resultWp := &users.User{
		UID:       result.UID,
		Name:      result.Name,
		LastName:  result.LastName,
		ThirdName: result.ThirdName,
	}

	return resultWp, nil
}

func GetUserByID(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{UID: userId}

	if err := result.GetByID(); err != nil {
		return nil, err
	}

	return result, nil
}
