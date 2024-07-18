package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-instagram/repository"
)

func AddProfileRoutes(rg *gin.RouterGroup) {
	posts := rg.Group("/profile")
	posts.GET("/create", func(context *gin.Context) {
		repository.CreateProfile()
		//context.JSON(210, CreateProfile())
	})
}
