package repository

import (
	"github.com/mahdi-cpp/api-go-instagram/config"
	"github.com/mahdi-cpp/api-go-instagram/model"
)

func InitStory() {

	var count int64 = 0
	config.DB.Find(&model.Story{}).Count(&count)
	if count > 0 {
		return
	}

	config.DB.Create(&model.StoryHighlight{Name: "Dubai", Cover: "cover-0000.jpg"})
	config.DB.Create(&model.StoryHighlight{Name: "Istanbul", Cover: "cover-0000.jpg"})

	config.DB.Create(&model.Story{UserID: 9, Video: "story/263993021_697272981237727_3647134352488780650_n.mp4"})
	config.DB.Create(&model.Story{UserID: 9, Video: "story/264037395_586844522419937_7253850601971856457_n.mp4"})
	config.DB.Create(&model.Story{UserID: 9, Video: "story/264070342_1253009261863694_4707964809382961830_n.mp4"})
	config.DB.Create(&model.Story{UserID: 9, Video: "story/264193724_1090786088374973_1083205495427038682_n.mp4"})
	config.DB.Create(&model.Story{UserID: 9, Video: "story/265531518_493662441937205_3057512426208477196_n.mp4"})
	config.DB.Create(&model.Story{UserID: 9, Video: "story/266294288_1520072435033015_7797002831637582954_n.mp4"})

	//config.DB.Create(&model.Story{UserID: 9, Photo: "story_0004.jpg"})
	//config.DB.Create(&model.Story{UserID: 9, Photo: "story_0005.jpg"})

	//config.DB.Create(&model.StoryMention{StoryRefer: 1, Username: "@mahdiabdolmaleki", X: 100, Y: 150})
	//config.DB.Create(&model.StoryMention{StoryRefer: 1, Username: "@ronaldo"})

	//config.DB.Create(&model.StoryMention{StoryRefer: 1, Username: "bitbod", X: 100, Y: 150})
}

func GetStories() []model.Story {

	var stories []model.Story

	config.DB.Debug().
		//Preload("StoryMentions").
		Preload("StoryQuestionBox").
		Preload("User").
		Order("id ASC").
		Find(&stories)

	return stories
}

func GetStoryHighlight() []model.StoryHighlight {

	var results []model.StoryHighlight

	config.DB.Debug().
		Preload("Stories").
		Order("id ASC").
		Find(&results)

	return results
}
