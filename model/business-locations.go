package model

import "gorm.io/gorm"

type BusinessLocation struct {
	gorm.Model
	BusinessId   uint   `json:"businessId"`
	MainLocation string `json:"mainLocation"`
	Description  string `json:"description"`
	Map          string `json:"map"`
}
