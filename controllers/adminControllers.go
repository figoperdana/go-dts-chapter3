package controllers

import (
	"go-jwt/database"
	"go-jwt/helpers"
	"go-jwt/models"
)

func CreateAdminUser() {
	// Connect to the database
	db := database.GetDB()

	// Create a new user with the role "admin"
	user := models.User{
		FullName: "Admin User",
		Email:    "admin@example.com",
		Password: "adminpassword",
		Role:     "admin",
	}

	// Hash the user's password before storing it in the database
	user.Password = helpers.HashPass(user.Password)

	// Save the user in the database
	db.Create(&user)
}
