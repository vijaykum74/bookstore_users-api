package users

import (

	"github.com/gin-gonic/gin"
	"github.com/vijaykum74/bookstore_users-api/domain/users"
	"github.com/vijaykum74/bookstore_users-api/services"
	"github.com/vijaykum74/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context)  {
	var user users.User
	if err := c.ShouldBindJSON(&user); err!=nil{
		restErr := errors.NewBadRequestError("invalid json data")
		c.JSON(restErr.Status,restErr)
		return
	}
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err !=nil{
	//	//TODO: Handle error
	//	return
	//}
	//
	//if err := json.Unmarshal(bytes,&user); err !=nil{
	//	fmt.Println(err.Error())
	//	//TODO: handle json error
	//	return
	//}

	result, saveErr := services.CreateUser(user)
	if saveErr !=nil{
		//TODO: handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}

func getUserId(userIdParam string)(int64, *errors.RestErr)  {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil{
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}
func GetUser(c *gin.Context)  {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr !=nil{
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil{
		c.JSON(getErr.Status,getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
