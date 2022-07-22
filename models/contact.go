package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Yname string `gorm:"not null;size:100" form:"yname"`
	Ymail string `gorm:"not null;size:100" form:"ymail"`
	Topic string`gorm:"not null;size:300" form:"topic"`
	Ymess string`gorm:"not null;size:500" form:"ymess"`
}
func CreateContact(db *gorm.DB, i Contact) error {
	return db.Create(&i).Error
}
func UPDate(db *gorm.DB, i *Contact, id int) {
	db.Where("id = ?", id).Updates(i)
}
func DeleTE(db *gorm.DB, i Contact, id int) {
	db.Where("id = ?", id).Delete(&i)
}

func GetContact(db *gorm.DB, i *Contact, yname, ymail, topic, ymess string) {
	db.Where("yname = ? ymail = ? topic = ? ymess = ?").First(i)
}
func GetContacts(db *gorm.DB, i *[]Contact) {
	db.Find(i)
}
