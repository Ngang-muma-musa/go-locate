package model

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	IsAdmin  bool   `json:"isAdmin"`
}

// Create User
func CreateUser(user *User) error {
	var existingUser User
	newUser := db.Where("email = ?", user.Email).Limit(1).Find(&existingUser)
	if newUser.RowsAffected > 0 {
		return errors.New("Email already exist")
	}
	newUser = db.Create(user)
	return newUser.Error
}

// Update user
func Updateuser(user *User) error {
	updatedUser := db.Save(user)
	return updatedUser.Error
}

// Get user by Id
func GetUserByID(ID uint) *User {
	var user User
	res := db.First(&user, ID)
	if res.RowsAffected == 0 {
		return nil
	}
	return &user
}

// Get user by email
func GetUserByEmail(email string) *User {
	var user User
	res := db.First(&user, "email = ?", email)
	if res.RowsAffected == 0 {
		return nil
	}
	return &user
}
