package controllers

import (
	"github.com/mahdi-cpp/api-go-instagram/model"
)

// Import the userModel from the model
var userModel = new(model.UserModel)

// UserController defines the user controller methods
type UserController struct{}

//
//// Signup controller handles registering a user
//func (u *UserController) Signup(c *gin.Context) {
//
//	var data forms.SignupUserCommand
//
//	// Bind the data from the request body to the SignupUserCommand Struct
//	// Also check if all fields are provided
//	if c.BindJSON(&data) != nil {
//		// specified response
//		c.JSON(406, gin.H{"message": "Provide relevant fields"})
//		// abort the request
//		c.Abort()
//		// return nothing
//		return
//	}
//	result, _ := userModel.GetUserByEmail(data.Email)
//
//	// If there happens to be a result respond with a
//	// descriptive mesage
//	if result.Email != "" {
//		c.JSON(403, gin.H{"message": "Email is already in use"})
//		c.Abort()
//		return
//	}
//
//	err := userModel.Signup(data)
//
//	resetToken, _ := services.GenerateNonAuthToken(data.Email)
//
//	link := "http://localhost:5000/api/v1/verify-account?verify_token=" + resetToken
//	body := "Here is your reset <a href='" + link + "'>link</a>"
//	html := "<strong>" + body + "</strong>"
//
//	email := services.SendMail("Verify Account", body, data.Email, html, data.Name)
//
//	// If email fails while sending
//	if !email {
//		c.JSON(500, gin.H{"message": "An issue occured sending you an email"})
//		c.Abort()
//		return
//	}
//
//	// Check if there was an error when saving user
//	if err != nil {
//		c.JSON(400, gin.H{"message": "Problem creating an account"})
//		c.Abort()
//		return
//	}
//
//	c.JSON(201, gin.H{"message": "New user account registered"})
//}
