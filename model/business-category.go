package model

import (
	"errors"

	"gorm.io/gorm"
)

type BusinessCategory struct {
	gorm.Model
	Category string `json:"category"`
}

// Create a new business category
func CreateBusinessCategory(category *BusinessCategory) error {
	var existingCategory BusinessCategory
	newCat := db.Where("categoty = ? ", category.Category).Limit(1).Find(&existingCategory)
	if newCat.RowsAffected > 0 {
		return errors.New("Category already exist")
	}
	newCat = db.Create(category)
	return newCat.Error
}

// Delete a business category
func DeleteBusinesCategory(category *BusinessCategory) {
	db.Delete(category)
}

// Get a business category by a specified ID
func GetBusinessCategoryByID(ID uint) *BusinessCategory {
	var category BusinessCategory
	res := db.First(&category, ID)
	if res.RowsAffected == 0 {
		return nil
	}
	return &category
}
