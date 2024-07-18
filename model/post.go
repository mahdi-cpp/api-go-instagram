package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Location string
	Caption  string
	Likes    uint
	IsIgtv   bool
	Medias   []Media `gorm:"foreignKey:PostRefer"`

	//A belongs to association sets up a one-to-one connection with another model
	UserID uint
	User   User

	Like     Like     `gorm:"foreignKey:PostID"`
	Bookmark Bookmark `gorm:"foreignKey:PostID"`
}

type Media struct {
	gorm.Model
	PostRefer     uint
	Thumbnail     string
	Photo         string
	Video         string
	VideoDuration int
	Width         int
	Height        int
	Tags          []Tag `gorm:"foreignKey:MediaRefer"`
}

type Category struct {
	Categoryid uint
	Name       string
	Photo      string
}

type Tag struct {
	gorm.Model
	MediaRefer uint
	Name       string
	X          int `gorm:"default:0"`
	Y          int `gorm:"default:0"`
}

type Like struct {
	gorm.Model
	UserID uint `gorm:"references:UserID"`
	PostID uint `gorm:"references:PostID"`
	Mahdi  string
}

type Bookmark struct {
	gorm.Model
	UserID uint `gorm:"references:UserID"`
	PostID uint `gorm:"references:PostID"`
}
