package model

import (
	"gorm.io/gorm"
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
