package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	UserType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	//if the type of user is not admin and the id is not of current user
	if UserType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	//otherwise check the type of user
	err = CheckUserType(c, UserType)
	return err
}
