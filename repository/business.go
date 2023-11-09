package repository

import (
	"errors"
	"go-locate/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Business struct {
	db *gorm.DB
}

func NewBusiness(db *gorm.DB) *Business {
	return &Business{db: db}
}

func (b *Business) Create(business *model.Business) error {
	return b.db.Create(business).Error
}

func (b *Business) Update(business *model.Business) error {
	return b.db.Save(business).Error
}

func (b *Business) GetByID(ID uint) (*model.Business, error) {
	var business model.Business
	err := b.db.Limit(1).Find(&business, ID).Error
	if err != nil {
		return nil, err
	}
	return &business, nil
}

func (b *Business) GetByName(name string) *model.Business {
	var business model.Business
	res := b.db.First(&business, "name = ?", name)
	if res.RowsAffected == 0 {
		return nil
	}
	return &business
}

func (b *Business) ListUnverified() []model.Business {
	var businesses []model.Business
	if err := b.db.Where("verified = ?", false).Find(&businesses).Error; err != nil {
		return nil
	}
	return businesses
}

func (b *Business) ListByUserID(ID uint) []model.Business {
	var business []model.Business
	if err := b.db.Where("userId = ?", ID).Find(&business).Error; err != nil {
		return nil
	}
	return business
}

func (b *Business) Delete(business *model.Business) error {
	return b.db.Delete(business).Error
}

func (b *Business) ListByVerificationStatus(status bool) []model.Business {
	var business []model.Business
	if err := b.db.Where("verified = ?", status).Find(&business).Error; err != nil {
		return nil
	}
	return business
}

func (b *Business) GetBusinessByCategoryOrLocation(options model.BusinessSearch) ([]model.Business, error) {
	var business []model.Business
	if options.Category == 0 && options.Location != "" {
		b.db.Joins("JOIN business_locations ON businesses.id = busniness_locations.business_id").Preload(clause.Associations).Find(&business, "businesses.verified = ?", true)
	} else if options.Category != 0 && options.Location == "" {
		b.db.Joins("JOIN business_categories ON businesses.id = business_categories.business_id").Preload(clause.Associations).Find(&business, "businesses.verified = ?", true)
	} else if options.Category != 0 && options.Location != "" {
		b.db.Joins("JOIN business_locations ON businesses.id = busniness_locations.business_id").Joins("JOIN business_categories ON businesses.id = business_categories.business_id").Preload(clause.Associations).Find(&business, "businesses.verified = ?", true)
	} else {
		return nil, errors.New("at least one query param needed")
	}
	return business, nil
}
