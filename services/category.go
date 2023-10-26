package services

import (
	"go-locate/model"
	"go-locate/repository"
)

type Category struct {
	categoryRepo *repository.Category
}

func NewCategory(categoryRepo *repository.Category) *Category {
	return &Category{categoryRepo: categoryRepo}
}

func (c *Category) Create(name string) (*model.Category, error) {
	category := &model.Category{Name: name}
	err := c.categoryRepo.Create(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}
