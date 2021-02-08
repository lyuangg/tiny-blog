package middlewares

import (
	"github.com/gin-gonic/gin"
	"tiny-blog/models"
	"tiny-blog/utils"
)

func ApiAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		user := utils.GetLoginUserInfo(*c)
		if user.ID <= 0 {
			c.JSON(200, gin.H{
				"code":    -1,
				"message": "not login",
			})
			c.Abort()
		} else {
			user = models.CheckUserToken(user.ID, user.Token)
			if user.ID <= 0 {
				c.JSON(200, gin.H{
					"code":    -2,
					"message": "not login",
				})
				c.Abort()
			} else {
				c.Set("userid", user.ID)
				c.Next()
			}
		}

	}
}
