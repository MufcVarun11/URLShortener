package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/MufcVarun11/go-url-shortener/handler"
	"github.com/MufcVarun11/go-url-shortener/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortLink(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortLinkRedirect(c)
	})

	// Note that store initialization happens here
	store.InitializeStorageService()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}