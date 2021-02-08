package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"tiny-blog/configs"
	"tiny-blog/models"
	"tiny-blog/utils"
)

func PostShow(c *gin.Context) {
	pid := c.Param("id")
	id, _ := strconv.Atoi(pid)
	post := models.GetPost(id)

	if post.ID <= 0 {
		c.String(http.StatusNotFound, "404 page not found!")
		return
	}

	c.HTML(http.StatusOK, "post.tmpl", gin.H{
		"user":  utils.GetLoginUserInfo(*c),
		"blog":  configs.Conf.Blog,
		"title": configs.Conf.Blog.Name + "-" + post.Title,
		"post":  post,
	})
}

// AboutPage about page
func AboutPage(c *gin.Context) {
	post := models.GetAboutPost()

	c.HTML(http.StatusOK, "about.tmpl", gin.H{
		"user":  utils.GetLoginUserInfo(*c),
		"blog":  configs.Conf.Blog,
		"title": configs.Conf.Blog.Name + "-about",
		"post":  post,
	})
}

func PostDelete(c *gin.Context) {

}

func PostCreate(c *gin.Context) {

}

func PostEdit(c *gin.Context) {
	pid := c.Param("id")
	id, _ := strconv.Atoi(pid)
	post := models.GetPost(id)

	user := utils.GetLoginUserInfo(*c)

	if user.ID > 0 {
		c.HTML(http.StatusOK, "edit.tmpl", gin.H{
			"user":  utils.GetLoginUserInfo(*c),
			"blog":  configs.Conf.Blog,
			"title": configs.Conf.Blog.Name + "-" + post.Title + "-edit",
			"post":  post,
		})
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

func PostAdd(c *gin.Context) {
	user := utils.GetLoginUserInfo(*c)

	if user.ID > 0 {
		c.HTML(http.StatusOK, "add.tmpl", gin.H{
			"user":  utils.GetLoginUserInfo(*c),
			"blog":  configs.Conf.Blog,
			"title": configs.Conf.Blog.Name + "-add",
		})
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

func ApiPostEdit(c *gin.Context) {
	pid := c.Param("id")
	id, _ := strconv.Atoi(pid)
	post := models.GetPost(id)

	if post.ID <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "post nof found!",
		})
	} else {
		var p models.Post
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		p.ID = id
		uid, _ := c.Get("userid")
		p.Author = uid.(int)
		models.UpdatePost(p)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
		})
	}

}

func ApiPostAdd(c *gin.Context) {
	var p models.Post
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	uid, _ := c.Get("userid")
	p.Author = uid.(int)
	p.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	id := models.AddPost(p)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"id":   id,
	})
}

func ApiPostDelete(c *gin.Context) {
	pid := c.Param("id")
	id, _ := strconv.Atoi(pid)
	post := models.GetPost(id)

	if post.ID <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "post nof found!",
		})
	} else {
		models.DeletePost(id)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
		})
	}
}
