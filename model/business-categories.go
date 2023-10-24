package model

import "gorm.io/gorm"

type BusinessCategory struct {
	gorm.Model
	CategoryId uint `json:"categoryId"`
	BusinessId uint `json:"businessId"`
}

func AddBusinesCategory() {

}
