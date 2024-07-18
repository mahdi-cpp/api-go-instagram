package model

import (
	"gorm.io/gorm"
)

// UserModel defines the model structure
type UserModel struct{}

//// Signup handles registering a user
//func (u *UserModel) Signup(data forms.SignupUserCommand) error {
//	// Connect to the user collection
//	collection := dbConnect.Use(databaseName, "user")
//	// Assign result to error object while saving user
//	err := collection.Insert(bson.M{
//		"name":     data.Name,
//		"email":    data.Email,
//		"password": helpers.GeneratePasswordHash([]byte(data.Password)),
//		// This will come later when adding verification
//		"is_verified": false,
//	})
//
//	return err
//}

type User struct {
	gorm.Model
	Username   string `gorm:"unique"`
	Email      string
	Phone      string
	Avatar     string
	FullName   string
	Biography  string
	Website    string
	Category   string
	IsVerified bool
	//FollowedByViewer bool
}

type Follow struct {
	gorm.Model
	UserID     uint `gorm:"primarykey;references:UserID"`
	FollowerID uint `gorm:"primarykey;references:UserID"`
}

type Follower struct {
	gorm.Model
	FollowerID uint `gorm:"primarykey;references:UserID"`

	//A belongs to association sets up a one-to-one connection with another model
	User User `gorm:"foreignKey:FollowerID"`
}

type Following struct {
	gorm.Model
	UserID uint `gorm:"primarykey;references:UserID"`

	//A belongs to association sets up a one-to-one connection with another model
	User User `gorm:"foreignKey:UserID"`
}
