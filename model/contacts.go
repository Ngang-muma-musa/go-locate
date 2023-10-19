package model

import (
	"errors"

	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Contact string `json:"contact"`
	UserID  int    `json:"userId"`
}

// create new business contact
func CreateContact(contact *Contact) error {
	var existtingContact Contact
	newContact := db.Where("contact = ?", contact.Contact).Limit(1).Find(&existtingContact)
	if newContact.RowsAffected > 0 {
		return errors.New("Contact already exist")
	}
	newContact = db.Create(contact)
	return newContact.Error
}

// Delete busines contact
func DeleteContact(contact *Contact) {
	db.Delete(contact)
}

// Get List of business contacts
func GetBusinessContactList(ID uint) []Contact {
	var contacts []Contact
	db.Where("userId = ?", ID).Find(&contacts)
	return contacts
}

// update business contact
func UpdateContact(contact Contact) error {
	updatedContact := db.Save(contact)
	return updatedContact.Error
}
