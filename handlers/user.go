package handlers

import (
	"net/http"
	"strconv"
	"tiny-blog/configs"
	"tiny-blog/models"
	"tiny-blog/utils"

	"github.com/gin-gonic/gin"
)

var (
	maxAge = 86400
)

func UserLogin(c *gin.Context) {
	email := c.PostForm("email")
	pwd := c.PostForm("password")

	user := models.UserLogin(email, pwd)
	domain := configs.Conf.Blog.Domain
	if user.ID > 0 {
		c.SetCookie("user_name", user.Name, maxAge, "/", domain, false, true)
		c.SetCookie("user_id", strconv.Itoa(user.ID), maxAge, "/", domain, false, true)
		c.SetCookie("user_token", user.Token, maxAge, "/", domain, false, true)

		c.Redirect(http.StatusFound, "/")
	} else {
		c.SetCookie("user_name", "", 0, "/", domain, false, true)
		c.SetCookie("user_id", "0", 0, "/", domain, false, true)
		c.SetCookie("user_token", "", 0, "/", domain, false, true)
		c.Redirect(http.StatusFound, "/loginfail")
	}
}

func UserProfile(c *gin.Context) {
	user := utils.GetLoginUserInfo(*c)

	if user.ID > 0 {
		c.HTML(http.StatusOK, "profile.tmpl", gin.H{
			"user":  user,
			"blog":  configs.Conf.Blog,
			"title": configs.Conf.Blog.Name + "-Profile",
		})
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

type Profile struct {
	Id       int    `form:"id" json:"id" xml:"id"  binding:"required"`
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"`
}

func ApiProfile(c *gin.Context) {
	var p Profile
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	if p.Id > 0 && p.UserName != "" {
		models.UpdateUserName(p.Id, p.UserName)
		domain := configs.Conf.Blog.Domain
		c.SetCookie("user_name", p.UserName, maxAge, "/", domain, false, true)
	}
	if p.Id > 0 && p.Password != "" {
		models.UpdateUserPwd(p.Id, p.Password)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func UserLogout(c *gin.Context) {
	user := utils.GetLoginUserInfo(*c)

	if user.ID > 0 {
		models.UpdateUserToken(user.ID, "")
	}
	domain := configs.Conf.Blog.Domain
	c.SetCookie("user_name", "", 0, "/", domain, false, true)
	c.SetCookie("user_id", "0", 0, "/", domain, false, true)
	c.SetCookie("user_token", "", 0, "/", domain, false, true)

	c.Redirect(http.StatusFound, "/")
}
