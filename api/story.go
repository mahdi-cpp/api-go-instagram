package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-instagram/repository"
)

func AddStoryRoutes(rg *gin.RouterGroup) {

	story := rg.Group("/story")

	story.GET("/stories", func(context *gin.Context) {
		context.JSON(210, repository.GetStories())
	})
	story.GET("/getStoryHighlight", func(context *gin.Context) {
		context.JSON(210, repository.GetStoryHighlight())
	})
}
