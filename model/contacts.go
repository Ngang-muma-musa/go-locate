package model

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Contact    string `json:"contact"`
	BusinessId uint   `json:"businessId"`
}
