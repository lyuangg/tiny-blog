package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tiny-blog/models"
)

func GetLoginUserInfo(c gin.Context) models.User {
	user := models.User{}
	userName, err := c.Cookie("user_name")
	if err == nil {
		user.Name = userName
	}
	userId, err := c.Cookie("user_id")
	if err == nil {
		user.ID, _ = strconv.Atoi(userId)
	}
	userToken, err := c.Cookie("user_token")
	if err == nil {
		user.Token = userToken
	}
	return user
}
