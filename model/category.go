package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Category string `json:"category"`
}

// Create a new business category
func CreateBusinessCategory(c *Category) error {
	category := db.Create(c)
	return category.Error
}

func GetCategoryByName(categoryName string) *Category {
	var category Category

	if err := db.First(&category, "category = ?", categoryName).Error; err != nil {
		return nil
	}
	return &category
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
