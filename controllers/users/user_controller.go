package users

import (
	"bookstore_users_api/domain/users"
	"bookstore_users_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "implement me!")

}

func CreateUser(c *gin.Context){
	var user users.User
	
	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO: handle json err
		return
	}
	result, saveError := services.CreateUser(user)
	if saveError != nil {
		//TODO: handle user creation err
		return
	}
	c.JSON(http.StatusCreated,result )
}

func SearchUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "implement me!")
}