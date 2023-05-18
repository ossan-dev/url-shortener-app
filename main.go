package main

import (
	"context"
	"math/rand"
	"time"
	"urlshortener/handlers"
	"urlshortener/models"
	"urlshortener/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost port=54322 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Url{})

	utils.Rand = rand.NewSource(time.Now().UnixNano())
	handlers.Store = make(map[string]string)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		timeoutCtx, cancelFunc := context.WithTimeout(ctx.Request.Context(), time.Second*6)
		defer cancelFunc()

		ctx.Set(utils.DbKey, db.WithContext(timeoutCtx))
		ctx.Next()
	})

	r.POST("/shorten", handlers.Shorten)
	r.POST("/unshorten", handlers.Unshorten)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
