package main

import (
	"github.com/gin-gonic/gin"
	"log"
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
	log.SetPrefix("[LOG] ")
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
}

func main() {
	router := gin.Default()
	// api
	apiGroup := router.Group("/api")
	apiGroup.GET("/quote", service.GetRandomQuote)            // get a random quote
	apiGroup.GET("/picture", service.GetRandomPicture)        // get a random picture
	apiGroup.GET("/backup-picture", service.GetBackupPicture) // get a random picture
	// apiGroup.GET("/test", service.Test) // test

	// web page

	router.LoadHTMLGlob("public/*.html")
	router.Static("/css", "./public/css")
	router.Static("/image", "./public/image")
	router.GET("/", service.IndexPage)

	// listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
