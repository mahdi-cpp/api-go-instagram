package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-instagram/api"
)

var (
	router = gin.Default()
)

// Run will Go-Instagram the server
func Run() {
	router.Use(CORSMiddleware())
	getRoutes()
	router.Run(":8095")
}

// getRoutes will create our api of our entire application
// this way every group of api can be defined in their own file
// so this one won't be so messy
func getRoutes() {

	v1 := router.Group("/v1")

	api.AddUserRoutes(v1)
	api.AddStoryRoutes(v1)
	api.AddPostRoutes(v1)
	api.AddProfileRoutes(v1)
	api.AddCommentRoutes(v1)
	api.AddUploadRoutes(v1)

	//v2 := router.Group("/v2")
	//AddStoryRoutes(v2)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
