package repository

import (
	"errors"
	"fmt"
	"go-locate/model"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

var ErrAlreadyExists = errors.New("already exist")

func (u *User) Create(user *model.User) error {
	var existingUser User
	newUser := u.db.Where("email = ?", user.Email).Limit(1).Find(&existingUser)
	if newUser.RowsAffected > 0 {
		return fmt.Errorf("%w: %s", ErrAlreadyExists, user.Email)
	}
	newUser = u.db.Create(user)
	return newUser.Error
}

func (u *User) Update(user *model.User) error {
	return u.db.Save(user).Error
}

func (u *User) GetByID(ID uint) *model.User {
	var user model.User
	res := u.db.First(&user, ID)
	if res.RowsAffected == 0 {
		return nil
	}
	return &user
}

func (u *User) GetByEmail(email string) *model.User {
	var user model.User
	res := u.db.First(&user, "email = ?", email)
	if res.RowsAffected == 0 {
		return nil
	}
	return &user
}
