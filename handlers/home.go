package handlers

import (
	"net/http"
	"strconv"
	"tiny-blog/configs"
	"tiny-blog/utils"

	"tiny-blog/models"

	"github.com/gin-gonic/gin"
)

// Index page
func Index(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := configs.Conf.Blog.PageSize
	p, _ := strconv.Atoi(page)

	user := utils.GetLoginUserInfo(*c)

	isAll := false
	if user.ID > 0 {
		isAll = true
	}

	posts, _ := models.GetPosts(pageSize, p, isAll)
	total := models.GetPostTotal(isAll)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"user":     user,
		"blog":     configs.Conf.Blog,
		"title":    configs.Conf.Blog.Name,
		"posts":    posts,
		"total":    total,
		"pageHtml": utils.CreatePaginatorHtml(p, pageSize, total),
	})
}

// Login page
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"user":  utils.GetLoginUserInfo(*c),
		"blog":  configs.Conf.Blog,
		"title": configs.Conf.Blog.Name + "-Login",
	})
}

func LoginFail(c *gin.Context) {
	c.HTML(http.StatusOK, "login-fail.tmpl", gin.H{
		"user":  utils.GetLoginUserInfo(*c),
		"blog":  configs.Conf.Blog,
		"title": configs.Conf.Blog.Name + "-LoginFail",
	})
}
