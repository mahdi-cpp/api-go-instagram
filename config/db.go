package config

import (
	"github.com/mahdi-cpp/api-go-instagram/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var err error

func Database() {

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=PostgreSQL user=mahdi password=aliali dbname=shopgram port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "app.", // schema name
			SingularTable: false,
		}})

	if err != nil {
		//panic("failed to connect database" + err.Error())
	}

	DB.AutoMigrate(&model.StoryHighlight{}, &model.Story{}, &model.StoryMention{}, &model.StoryQuestionBox{})
	// Migrate the schema
	DB.AutoMigrate(&model.User{}, &model.Follow{}, &model.Post{}, &model.Media{}, &model.Tag{}, &model.Like{}, &model.Bookmark{})
	DB.AutoMigrate(&model.Comment{})
}
