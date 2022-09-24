package users

import (
	"bookstore_users_api/domain/users"
	"bookstore_users_api/services"
	"bookstore_users_api/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "implement me!")

}

func CreateUser(c *gin.Context){
	var user users.User
	
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}
	c.JSON(http.StatusCreated,result )
}

func SearchUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "implement me!")
}