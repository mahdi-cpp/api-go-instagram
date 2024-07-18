package repository

import (
	"github.com/mahdi-cpp/api-go-instagram/config"
	"github.com/mahdi-cpp/api-go-instagram/model"
)

func GetCommentsById() []model.Comment {
	var comments []model.Comment
	config.DB.Preload("User").Find(&comments)
	return comments
}
