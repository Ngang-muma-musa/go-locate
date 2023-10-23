package model

import (
	"errors"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Category string `json:"category"`
}

// Create a new business category
func CreateBusinessCategory(category *Category) error {
	var existingCategory Category
	newCat := db.Where("categoty = ? ", category.Category).Limit(1).Find(&existingCategory)
	if newCat.RowsAffected > 0 {
		return errors.New("Category already exist")
	}
	newCat = db.Create(category)
	return newCat.Error
}

// Delete a business category
func DeleteBusinesCategory(category *Category) {
	db.Delete(category)
}

// Get a business category by a specified ID
func GetBusinessCategoryByID(ID uint) *Category {
	var category Category
	res := db.First(&category, ID)
	if res.RowsAffected == 0 {
		return nil
	}
	return &category
}
