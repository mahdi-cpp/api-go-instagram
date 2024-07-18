package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-instagram/repository"
	"net/http"
)

func AddCommentRoutes(rg *gin.RouterGroup) {

	users := rg.Group("/comments")

	users.GET("/getComments", func(context *gin.Context) {
		context.JSON(http.StatusOK, repository.GetCommentsById())
	})
}
