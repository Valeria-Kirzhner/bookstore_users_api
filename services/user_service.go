package services

import (
	"bookstore_users_api/domain/users"
	"bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}