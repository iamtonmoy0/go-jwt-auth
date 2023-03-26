package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// checking user type
func CheckUserType(c *gin.Context, role string) (err error) {

}

// matching  user
func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	user_type := c.GetString("user_type")
	uid := c.GetString("userId")
	err = nil

	if user_type == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err

	}
	err = CheckUserType(c, userType)
	return err
}
