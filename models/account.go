package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Email    string `gorm:"not null;size:300" form:"email"`
	Password string `gorm:"not null;size:300" form:"password"`
}

func CreateAccount(db *gorm.DB, a Account) error {
	return db.Create(&a).Error
}

func UPdate(db *gorm.DB, a Account, id int) {
	db.Where("id = ?", id).Updates(&a)
}
func Deletechan(db *gorm.DB, a Account, id int) {
	db.Where("id = ?", id).Delete(&a)
}

func GetAccount(db *gorm.DB, a *Account, email, pass string) {
	db.Where("email = ? and password = ?", email, pass).First(a)
}
