//dao = data access object. This is the only point where we actually interact with the db.
package users

import (
	"bookstore_users_api/datasources/mysql/users_db"
	"bookstore_users_api/utils/date_utils"
	"bookstore_users_api/utils/errors"
	"fmt"
	"strings"
)
const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	duplicatedEntry = "1062"
)
var ( usersDB = make(map[int64]*User))

func (user *User) Get()  *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return  nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), duplicatedEntry) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email ))
			
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when truing to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when truing to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}