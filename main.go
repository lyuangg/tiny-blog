package main

import (
	"math/rand"
	"strconv"
	"time"
	"tiny-blog/configs"
	"tiny-blog/middlewares"
	"tiny-blog/models"
	"tiny-blog/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	configs.LoadConfig()
	gin.SetMode(configs.Conf.Mode)

	middlewares.InitLogger()
	models.DbInit()
	r := routes.InitRouter()

	host := configs.Conf.Host + ":" + strconv.Itoa(configs.Conf.Port)

	r.Run(host)
}
