package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-instagram/repository"
	"log"
	"net/http"
)

func AddUploadRoutes(rg *gin.RouterGroup) {

	posts := rg.Group("/upload")

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	//api.router.MaxMultipartMemory = 28 << 20 // 8 MiB

	posts.POST("/photo", func(c *gin.Context) {

		repository.CreateProfile()
		//context.JSON(210, CreateProfile())

		//single file
		file, _ := c.FormFile("photo")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		err := c.SaveUploadedFile(file, "/home/mahdi/photos/ali/Top.mkv")
		if err != nil {
			c.String(http.StatusOK, fmt.Sprintf("'%s' error!\n", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!\n", file.Filename))
	})

}
