package repository

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-instagram/config"
	"github.com/mahdi-cpp/api-go-instagram/model"
)

func InitPost() {

	var count int64 = 0
	config.DB.Find(&model.Post{}).Count(&count)
	if count > 0 {
		return
	}

	config.DB.Create(&model.Post{UserID: 1, Caption: "dear mahdi... ", Location: "Tehran"})
	config.DB.Create(&model.Post{UserID: 2, Caption: "hello every body, I am in germany!!!", Location: "Berlin"})
	config.DB.Create(&model.Post{UserID: 3, Caption: "پایان روز اول جیتکس 2021", Location: "Dubai"})
	config.DB.Create(&model.Post{UserID: 4, Caption: "صنعت پارچه یکی از قدیمی‌ترین صنایع بشری است، اما نتوانسته به راحتی وارد دنیای تکنولوژی...", Location: "Dubai"})

	config.DB.Create(&model.Media{Thumbnail: "0001.jpg", PostRefer: 1})
	config.DB.Create(&model.Media{Thumbnail: "0002.jpg", PostRefer: 2})
	config.DB.Create(&model.Media{Thumbnail: "0003.jpg", PostRefer: 3})
	config.DB.Create(&model.Media{Thumbnail: "0004.jpg", PostRefer: 4})

	config.DB.Create(&model.Tag{Name: "@mahdiabdolmaleki", MediaRefer: 1, X: 100, Y: 150})
	config.DB.Create(&model.Tag{Name: "@ronaldo", MediaRefer: 1})

	config.DB.Create(&model.Tag{Name: "bitbod", MediaRefer: 3, X: 100, Y: 150})
}

func GetAllPosts() []model.Post {

	var posts []model.Post
	//config.DB.Preload("Medias.Hashtags").Preload("Medias").Find(&posts)

	config.DB.
		Preload("Medias.Tags").
		Preload("Medias").
		Preload("User.Follower").
		Preload("User").
		Preload("Like", "user_id = ?", 1).
		Preload("Bookmark", "user_id = ?", 1).
		Offset(0).
		Limit(200).
		Order("id ASC").
		Find(&posts)

	//for i, post := range posts {
	//	fmt.Println(i+1, post.Caption)
	//}
	return posts
}

func GetPostById(id uint) model.Post {

	var post model.Post

	config.DB.
		Preload("Medias.Tags").
		Preload("Medias").
		Preload("User.Follower").
		Preload("User").
		Preload("Like", "user_id = ?", 1).
		Preload("Bookmark", "user_id = ?", 1).
		Where("id", 431).
		Find(&post)

	fmt.Println(post.ID)

	return post
}

func GetAllReals() []model.Post {

	var posts []model.Post
	//config.DB.Preload("Medias.Hashtags").Preload("Medias").Find(&posts)

	config.DB.
		Preload("Medias.Tags").
		Preload("Medias").
		Preload("User.Follower").
		Preload("User").
		Preload("Like", "user_id = ?", 1).
		Preload("Bookmark", "user_id = ?", 1).
		Offset(0).
		Limit(200).
		Order("id ASC").
		Where("is_igtv", "true").
		Find(&posts)

	//for i, post := range posts {
	//	fmt.Println(i+1, post.Caption)
	//}
	return posts
}

func GetAllCategory() []model.Category {

	var categories []model.Category

	var ali model.Category
	ali.Categoryid = 1
	ali.Name = "Ali"
	ali.Photo = ""

	var mahdi model.Category
	mahdi.Categoryid = 2
	mahdi.Name = "Mahdi"
	mahdi.Photo = ""

	categories = append(categories, ali)

	return categories
}

func GetAllMedias() []model.Media {

	var medias []model.Media
	config.DB.Find(&medias)

	return medias
}

func GetPostsFeed(page int, size int) []model.Post {
	var currentUser = 1
	var posts []model.Post

	config.DB.
		Preload("Medias.Tags").
		Preload("Medias").
		Preload("User.Follower").
		Preload("User").
		Preload("Like", "user_id = ?", currentUser).
		Preload("Bookmark", "user_id = ?", currentUser).
		Offset(0).
		Limit(200).
		Order("id ASC").
		Find(&posts)

	return posts
}
