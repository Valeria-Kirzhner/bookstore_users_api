package users

import (
	"bookstore_users_api/domain/users"
	"bookstore_users_api/services"
	"bookstore_users_api/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
func getUserId(userIdParam string) (int64, *errors.RestErr){
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
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

func GetUser(c *gin.Context){

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK,user )
}

func UpdateUser(c *gin.Context){
	var user users.User
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}
	user.Id = userId

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context){
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}

	if err := services.DeleteUser(userId); err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}