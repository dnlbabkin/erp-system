package users

import (
	"back/utils/errors"
	"strings"
)

type User struct {
	UID       int64  `json:"uid"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	ThirdName string `json:"third_name"`
	Password  string `json:"password"`
}

func (user *User) Validate() *errors.RestErr {
	user.Name = strings.TrimSpace(user.Name)
	user.LastName = strings.TrimSpace(user.LastName)
	user.ThirdName = strings.TrimSpace(user.ThirdName)

	if user.Name == "" && user.LastName == "" && user.ThirdName == "" {
		return errors.NewBadRequestError("invalid user names")
	}

	user.Password = strings.TrimSpace(user.Password)

	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
