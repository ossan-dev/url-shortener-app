package main

import (
	"math/rand"
	"sync"
	"time"
	"urlshortener/handlers"
	"urlshortener/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.Rand = rand.NewSource(time.Now().UnixNano())
	handlers.Store = sync.Map{}
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.POST("/shorten", handlers.Shorten)
	r.POST("/unshorten", handlers.Unshorten)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
