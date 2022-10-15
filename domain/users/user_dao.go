//dao = data access object. This is the only point where we actually interact with the db.
package users

import (
	"bookstore_users_api/datasources/mysql/users_db"
	"bookstore_users_api/utils/date_utils"
	"bookstore_users_api/utils/errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)
const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	errorDuplicatedEntry = "1062"
	queryGetUser = "SELECT id, first_name,last_name, email, date_created FROM users WHERE id=?;"
	errorNoRows = "no rows in result set"
)

func (user *User) Get()  *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser);
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewBadRequestError(fmt.Sprintf("user %d does not found", user.Id ))
		}
		fmt.Println(err)
		return errors.NewInternalServerError(fmt.Sprintf("error when truing to get user %d %s:", user.Id,err.Error()))
	}
	return  nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, saveError := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveError != nil {

		sqlError, ok := saveError.(*mysql.MySQLError) //truing to make type cast to chek if its a mysql error type. I Shell use that to be able to know what error number it is and switch and act by the mysql err num.
		if !ok{
			return errors.NewInternalServerError(fmt.Sprintf("error when truing to save user: %s", err.Error()))
		}
		fmt.Println(sqlError.Number)
		fmt.Println(sqlError.Message)

		return errors.NewInternalServerError(fmt.Sprintf("error when truing to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when truing to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}