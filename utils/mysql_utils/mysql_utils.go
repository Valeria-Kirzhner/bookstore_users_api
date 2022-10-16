package mysql_utils

import (
	"bookstore_users_api/utils/errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
	errorDuplicatedEntry = "1062"
)
func ParseError(err error) *errors.RestErr{
	sqlError, ok := err.(*mysql.MySQLError) //truing to make a type cast to a pinter of mysql error. I shoud use that to be able to know what error number it is and switch and act by the mysql err num.
	if !ok {
		fmt.Println("the err is",err.Error())
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlError.Number {
	case 1062:
		return errors.NewInternalServerError("invalid data")
	}
	return errors.NewInternalServerError("error proccessing request")
}