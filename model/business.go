package model

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Location    string             `json:"loaction"`
	Contact     []string           `json:"contact"`
	Category    []BusinessCategory `json:"category"`
	Verified    bool               `json:"verified"`
	UserID      uint               `json:"userId"`
}

type BusinessSearch struct {
	Location string
	Category int
}

// Create a new business
func CreateBusiness(business *Business) error {
	var existingBusiness Business
	newBusiness := db.Where("name = ?", business.Name).Limit(1).Find(&existingBusiness)

	if newBusiness.RowsAffected > 0 {
		return errors.New("Business Name already taken")
	}

	newBusiness = db.Create(business)
	return newBusiness.Error
}

// Update a business
func UpdateBusiness(business *Business) error {
	updatedBusiness := db.Save(business)
	return updatedBusiness.Error
}

// Delete Business
func DeleteBusiness(business *Business) {
	db.Delete(business)
}

// List a users businesses
func ListUserBusiness(ID uint) []Business {
	var business []Business
	db.Where("userId = ?", ID).Find(&business)
	return business
}

// Get business by ID
func GetBusinessByID(ID uint) *Business {
	var busines Business
	res := db.First(&busines, ID)

	if res.RowsAffected == 0 {
		return nil
	}
	return &busines
}

// get business by name
func GetBusinessByName(name string) *Business {
	var busines Business
	res := db.First(&busines, strings.ToLower(name))

	if res.RowsAffected == 0 {
		return nil
	}
	return &busines
}

// Find business by Location and Category

func GetBusinessByCategoryOrBusiness(options BusinessSearch) ([]Business, error) {
	var business []Business
	res := db.Where("category = ?", options.Category).Find(&business)
	if options.Location != "" {
		res.Where("location = ?", options.Location)
	}
	if business == nil {
		return nil, errors.New("Error gettin business")
	}
	return business, nil
}
