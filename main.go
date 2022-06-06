// main.go

package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "webapp/docs"
)

var router *gin.Engine

// @title Project Golang Maktabkhoneh API
// @version 1.0
// @description This is a web application that with the help of gin library and swagger document can verify the national number of people and also return the information of that person in response.
// @contact.name API Support (My GitHub)
// @contact.url https://github.com/saber-khakbiz
// @contact.email saberkhakbiz73@g

// @host 127.0.0.1:8080
// @BasePath
// @query.collection.format multi

func main() {
	// Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	url := ginSwagger.URL("127.0.0.1:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Initialize the routes
	initRoute()

	// Start serving the application
	router.Run(":8080")
}
