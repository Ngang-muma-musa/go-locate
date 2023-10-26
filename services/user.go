package services

import (
	"go-locate/model"
	"go-locate/repository"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	userRepo *repository.User
}

func NewUser(userRepo *repository.User) *User {
	return &User{userRepo: userRepo}
}

func (u *User) Create(username string, password string, email string) (*model.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Username: username,
		Email:    email,
		Password: string(passwordHash),
		IsAdmin:  false,
	}

	err = u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) GetByEmail(email string) *model.User {
	return u.userRepo.GetByEmail(email)
}

func (u *User) GetByID(ID uint) *model.User {
	return u.userRepo.GetByID(ID)
}
