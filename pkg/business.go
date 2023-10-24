package pkg

import (
	"errors"
	"go-locate/model"
)

type BusinessInfo struct {
	Business model.Business   `json:"business"`
	Contact  []model.Contact  `json:"contact"`
	Category []model.Category `json:"category"`
}

func CreateBusiness(name string, description string, email string, location string, user *model.User, contact []model.Contact, category []model.BusinessCategory) (*model.Business, error) {
	business := &model.Business{
		Name:         name,
		Description:  description,
		Location:     location,
		PhoneNumbers: contact,
		Categories:   category,
		Email:        email,
		UserID:       user.ID,
		Verified:     false,
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

func FindBusiness(location string, category int) (*[]model.Business, error) {
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
