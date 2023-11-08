package services

import (
	"errors"
	"go-locate/model"
	"go-locate/repository"
)

type Business struct {
	businessRepo *repository.Business
}

func NewBusiness(businessRepo *repository.Business) *Business {
	return &Business{businessRepo: businessRepo}
}

func (b *Business) Create(name string, description string, email string, location []model.BusinessLocation, user *model.User, contact []model.Contact, category []model.BusinessCategory) (*model.Business, error) {
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

	if businessExist := b.businessRepo.GetByName(name); businessExist != nil {
		return nil, errors.New("business name already taken")
	}

	err := b.businessRepo.Create(business)
	if err != nil {
		return nil, err
	}

	return business, nil
}

func (b *Business) Find(location string, category int) ([]model.Business, error) {
	params := model.BusinessSearch{
		Location: location,
		Category: category,
	}
	business, err := b.businessRepo.GetBusinessByCategoryOrLocation(params)
	if err != nil {
		return nil, err
	}
	return business, err
}

func (b *Business) GetByID(ID uint) (*model.Business, error) {
	return b.businessRepo.GetByID(ID)
}
