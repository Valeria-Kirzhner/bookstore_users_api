//dao = data access object. This is the only point where we actually interact with the db.
package users

import (
	"bookstore_users_api/datasources/mysql/users_db"
	"bookstore_users_api/utils/date_utils"
	"bookstore_users_api/utils/errors"
	"bookstore_users_api/utils/mysql_utils"
)
const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	queryGetUser = "SELECT id, first_name,last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?;"
	querDeleteUser = "DELETE FROM users WHERE id=?;"

)

func (user *User) Get()  *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser);
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return  nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)// in use of prepare i have a choise of validating the data first.
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, saveError := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveError != nil {
		return mysql_utils.ParseError(saveError)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(querDeleteUser);
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil{
		return mysql_utils.ParseError(err)

	}
	return nil
}