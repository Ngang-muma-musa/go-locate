package model

import (
	"errors"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Business struct {
	gorm.Model
	Name         string             `json:"name"`
	Email        string             `gorm:"unique" json:"email"`
	PhoneNumbers []Contact          `json:"phonenumbers" `
	Categories   []BusinessCategory `json:"categories"`
	Description  string             `json:"description"`
	Location     string             `json:"loaction"`
	Verified     bool               `json:"verified"`
	UserID       uint               `json:"userId"`
}

type BusinessSearch struct {
	Location string
	Category int
}

// Create a new business
func CreateBusiness(business *Business) error {
	err := db.Create(business)
	if err != nil {
		return nil
	}
	return err.Error
}

// Update a business
func UpdateBusiness(business *Business) error {
	err := db.Save(business).Error
	if err != nil {
		return nil
	}
	return err
}

// Delete Business
func DeleteBusiness(business *Business) {
	db.Delete(business)
}

// List a users businesses
func ListUserBusiness(ID uint) []Business {
	var business []Business
	if err := db.Where("userId = ?", ID).Find(&business).Error; err != nil {
		return nil
	}
	return business
}

// Get business by ID
func GetBusinessByID(ID uint) *Business {
	var busines Business
	if err := db.First(&busines, ID).Error; err != nil {
		return nil
	}
	return &busines
}

// get business by name
func GetBusinessByName(name string) *Business {
	var business Business
	if err := db.First(&business, "name = ?", strings.ToLower(name)).Error; err != nil {
		return nil
	}
	return &business
}

// Find business by Location and Category
func GetBusinessByCategoryOrLocation(options BusinessSearch) (*[]Business, error) {
	var business []Business
	if options.Category == 0 && options.Location != "" {
		db.Where("location = ?", options.Location).Where("verified = ?", true).Find(&business)
	} else if options.Category != 0 && options.Location == "" {
		db.Model(&Business{}).Preload(clause.Associations).Where("verified = ?", true).Find(&business)
	} else if options.Category != 0 && options.Location != "" {
		db.Model(&Business{}).Preload(clause.Associations).Where("location = ?", options.Location).Where("verified = ?", true).Find(&business)
	} else {
		return nil, errors.New("Atleat one query param needed")
	}
	return &business, nil
}

func GetBusinessByVerificationStatus(status bool) (*[]Business, error) {
	var business []Business
	if err := db.Where("verified = ?", status).Find(&business).Error; err != nil {
		return nil, db.Error
	}

	return &business, nil
}
