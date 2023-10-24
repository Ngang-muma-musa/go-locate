package pkg

import (
	"errors"
	"go-locate/model"
)

func VarifyBusiness(ID int) error {
	busines, err := model.GetBusinessByID(ID)
	if err != nil {
		return err
	}
	busines.Verified = true
	err = model.UpdateBusiness(busines)
	return err
}

func CreateCategory(c string) (*model.Category, error) {
	if categoryExist := model.GetCategoryByName(c); categoryExist != nil {
		return nil, errors.New("Category Already Exist")
	}

	category := &model.Category{
		Category: c,
	}

	err := model.CreateBusinessCategory(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}
