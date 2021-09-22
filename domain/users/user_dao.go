package users

import (
	"fmt"
	"github.com/vijaykum74/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)
func (user *User)Get() (*errors.RestErr) {
	userId := user.Id
	result := userDB[userId]
	if result==nil{
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", userId))
	}
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User)Save() *errors.RestErr{
	if userDB[user.Id]!=nil{
		return errors.NewBadRequestError(fmt.Sprintf("user %d is already exist", user.Id))
	}
	userDB[user.Id]=user
	return nil

}