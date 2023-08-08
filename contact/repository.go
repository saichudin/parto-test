package contact

import (
	"saichudin/parto-test/config"
)

func GetAllContactRepo() *[]Contact {
	var contacts []Contact
	config.Db.Model(&Contact{}).Find(&contacts)

	return &contacts
}

func GetDetailContactRepo(id int) (*Contact, error) {
	var contact Contact
	err := config.Db.Model(&Contact{}).Where("id = ?", id).First(&contact).Error
	
	return &contact, err
}

func AddContactRepo(contact *Contact) error {
	err := config.Db.Create(&contact).Error

	return err
}

func UpdateContactRepo(id int, contact *Contact) error {
	_, errGet := GetDetailContactRepo(id)
	if errGet != nil {
		return errGet
	}
	contact.Id = id
	err := config.Db.Save(&contact).Error

	return err
}

func DeleteContactRepo(id int) error {
	contact, errGet := GetDetailContactRepo(id)
	if errGet != nil {
		return errGet
	}

	config.Db.Delete(&contact)
	return nil
}