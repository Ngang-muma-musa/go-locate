package pkg

import (
	"go-locate/model"

	"golang.org/x/crypto/bcrypt"
)

func createUser(username string, password string, email string) (*model.User, error) {
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

	err = model.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
