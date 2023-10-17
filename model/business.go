package model

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	Name        string
	description string
	location    string
	contact     []string
	category    []int
	verified    bool
	userID      int
}

// Create a new business
func createBusiness(business *Business) error {
	var existingBusiness Business
	newBusiness := db.Where("name = ?", business.Name).Limit(1).Find(&existingBusiness)

	if newBusiness.RowsAffected > 0 {
		return errors.New("Business Name already taken")
	}

	newBusiness = db.Create(business)
	return newBusiness.Error
}

// Update a business
func updateBusiness(business *Business) error {
	updatedBusiness := db.Save(business)
	return updatedBusiness.Error
}

// Delete Business
func deleteBusiness(business *Business) {
	db.Delete(business)
}

// List a users businesses
func listUserBusiness(ID uint) []Business {
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
