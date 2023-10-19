package pkg

import (
	"errors"
	"go-locate/model"
)

type BusinessInfo struct {
	Business model.Business           `json:"business"`
	Contact  []model.Contact          `json:"contact"`
	Category []model.BusinessCategory `json:"category"`
}

func CreateBusiness(name string, description string, email string, location string, phoneNumber string, user model.User, category string) (*model.Business, error) {
	business := &model.Business{
		Name:        name,
		Description: description,
		Location:    location,
		Email:       email,
		UserID:      user.ID,
		PhoneNumber: phoneNumber,
		Category:    category,
		Verified:    false,
	}

	if businessExist := model.GetBusinessByName(name); businessExist != nil {
		return nil, errors.New("Business namae already taken")
	}

	err := model.CreateBusiness(business)
	if err != nil {
		return nil, err
	}

	return business, nil

}

func findBusiness(location string, category string) (*[]model.Business, error) {
	params := &model.BusinessSearch{
		Location: location,
		Category: category,
	}
	business, err := model.GetBusinessByCategoryOrLocation(*params)
	if err != nil {
		return nil, err
	}
	return business, err
}
