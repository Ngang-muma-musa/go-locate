package pkg

import (
	"errors"
	"go-locate/model"
)

func CreateBusiness(name string, description string, location string, contact []string, user model.User, category []model.BusinessCategory) (*model.Business, error) {
	business := &model.Business{
		Name:        name,
		Description: description,
		Location:    location,
		UserID:      user.ID,
		Verified:    false,
		Contact:     contact,
		Category:    category,
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
