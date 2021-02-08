package routes

import (
	"tiny-blog/handlers"
	"tiny-blog/middlewares"

	"github.com/gin-gonic/gin"
)

// InitRouter init router config
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.LoggerMiddleware())
	r.Use(gin.Recovery())

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.Index)
	r.GET("/loginfail", handlers.LoginFail)
	r.GET("/login", handlers.Login)
	r.GET("/logout", handlers.UserLogout)
	r.POST("/login", handlers.UserLogin)
	r.GET("/post/:id", handlers.PostShow)
	r.GET("/about", handlers.AboutPage)
	r.GET("/profile", handlers.UserProfile)

	r.GET("/edit/:id", handlers.PostEdit)
	r.GET("/add", handlers.PostAdd)

	api := r.Group("/api")
	api.Use(middlewares.ApiAuthMiddleware())
	{
		api.POST("profile", handlers.ApiProfile)
		api.POST("post", handlers.ApiPostAdd)
		api.PUT("post/:id", handlers.ApiPostEdit)
		api.DELETE("post/:id", handlers.ApiPostDelete)
	}

	return r
}
