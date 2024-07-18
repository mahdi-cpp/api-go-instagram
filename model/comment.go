package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostRefer uint
	Message   string

	//A belongs to association sets up a one-to-one connection with another model
	UserID uint
	User   User
}
