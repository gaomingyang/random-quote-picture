package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"random-quote-picture/config"
	"random-quote-picture/service"
)

func init() {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config.Init(workPath)
}

func main() {
	router := gin.Default()
	// api
	apiGroup := router.Group("/api")
	apiGroup.GET("/quote", service.GetRandomQuote)     // get a random quote
	apiGroup.GET("/picture", service.GetRandomPicture) // get a random picture

	apiGroup.GET("/test", service.Test) // test

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// web page
	router.LoadHTMLGlob("public/*")
	router.GET("/", service.IndexPage)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
