package repository

import (
	"go-locate/model"
	"gorm.io/gorm"
)

type Category struct {
	db *gorm.DB
}

func NewCategory(db *gorm.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(category *model.Category) error {
	return c.db.Create(category).Error
}

func (c *Category) Update(category *model.Category) error {
	return c.db.Save(category).Error
}

func (c *Category) Delete(category *model.Category) error {
	return c.db.Delete(category).Error
}

func (c *Category) GetByName(name string) (*model.Category, error) {
	var category model.Category
	err := c.db.Limit(1).Find(&category, "name = ?", name).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}
