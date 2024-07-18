package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-instagram/repository"
	"strconv"
)

func AddPostRoutes(rg *gin.RouterGroup) {

	posts := rg.Group("/post")

	posts.GET("/post", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Query("id"))
		post := repository.GetPostById(uint(id))
		context.JSON(210, gin.H{"post": post})
	})

	posts.GET("/posts", func(context *gin.Context) {
		context.JSON(210, repository.GetAllPosts())
	})

	posts.GET("/profile_reels", func(context *gin.Context) {

		posts := repository.GetAllReals()
		context.JSON(210, gin.H{"posts": posts})
	})

	posts.GET("/profile_igtv", func(context *gin.Context) {

		posts := repository.GetAllReals()
		context.JSON(210, gin.H{"posts": posts})
	})

	posts.GET("/explore", func(context *gin.Context) {

		posts := repository.GetAllPosts()
		categories := repository.GetAllCategory()

		context.JSON(210, gin.H{"posts": posts, "categories": categories})
	})

	posts.GET("/medias", func(context *gin.Context) {
		context.JSON(210, repository.GetAllMedias())
	})

	posts.GET("/feed", func(context *gin.Context) {
		page, _ := strconv.Atoi(context.Query("page"))
		size, _ := strconv.Atoi(context.Query("size"))

		posts := repository.GetPostsFeed(page, size)
		users := repository.GetUsers()

		//0xc9849e6fdb743d08faee3e34dd2d1bc69ea11a51
		//0x8bdd8dbcbdf0c066ca5f3286d33673aa7a553c10

		context.JSON(210, gin.H{"posts": posts, "stories": users})
	})

	posts.GET("/profile", func(context *gin.Context) {
		posts := repository.GetPostsFeed(0, 0)
		user := repository.UserByID(9)
		//fmt.Println(user)
		context.JSON(210, gin.H{"posts": posts, "user": user, "postsCount": 12, "followersCount": 13, "followingCount": 14})
	})
}
