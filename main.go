package main

import (
	"math/rand"
	"time"
	"urlshortener/handlers"
	"urlshortener/middlewares"
	"urlshortener/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.Rand = rand.NewSource(time.Now().UnixNano())
	handlers.Store = make(map[string]string)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.POST("/shorten", middlewares.CheckAuthentication(), handlers.Shorten)
	r.POST("/unshorten", middlewares.CheckAuthentication(), handlers.Unshorten)

	r.POST("/signup", handlers.Signup)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
