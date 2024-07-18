package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-instagram/repository"
	"net/http"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/getFollowers", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.GetUsers())
	})
	users.GET("/getFollowing", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.GetUsers())
	})
}
