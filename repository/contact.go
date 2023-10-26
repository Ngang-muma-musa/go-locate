package repository

import (
	"go-locate/model"
	"gorm.io/gorm"
)

type Contact struct {
	db *gorm.DB
}

func NewContact(db *gorm.DB) *Contact {
	return &Contact{db: db}
}

func (c *Contact) Create(contact *model.Contact) error {
	return c.db.Create(contact).Error
}

func (c *Contact) Update(contact *model.Contact) error {
	return c.db.Save(contact).Error
}

func (c *Contact) Delete(contact *model.Contact) error {
	return c.db.Delete(contact).Error
}

func (c *Contact) GetByID(ID uint) (*model.Contact, error) {
	var contact model.Contact
	err := c.db.Limit(1).Find(&contact, ID).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (c *Contact) ListByBusinessID(id uint) []model.Contact {
	var contacts []model.Contact
	c.db.Where("business_id = ?", id).Find(&contacts)
	return contacts
}
