package users

import (
	"back/utils/errors"
	"strings"
)

type User struct {
	UID        int64  `json:"uid"`
	Name       string `json:"name"`
	SecondName string `json:"second_name"`
	ThirdName  string `json:"third_name"`
	Password   string `json:"password"`
	Token      string `json:"token"`
}

func (user *User) Validate() *errors.RestErr {
	user.Name = strings.TrimSpace(user.Name)
	user.SecondName = strings.TrimSpace(user.SecondName)
	user.ThirdName = strings.TrimSpace(user.ThirdName)

	if user.Name == "" && user.SecondName == "" && user.ThirdName == "" {
		return errors.NewBadRequestError("invalid user names")
	}

	user.Password = strings.TrimSpace(user.Password)

	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
