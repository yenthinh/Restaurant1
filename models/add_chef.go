package models

import "gorm.io/gorm"

type People struct {
	gorm.Model
	Names       string `gorm:"not null;size:150" form:"names"`
	Position string `gorm:"not null;size:300" form:"posi"`
	Picture       string `gorm:"not null;size:100" form:"pic"`
}

func CreatePeople(db *gorm.DB, i People) error {
	return db.Create(&i).Error
}
func UpDate(db *gorm.DB, i *People, id int) {
	db.Where("id = ?", id).Updates(i)
}
func DElete(db *gorm.DB, i People, id int) {
	db.Where("id = ?", id).Delete(&i)
}

func GetPeople(db *gorm.DB, i *People, names, posi, pic string) {
	db.Where("names = ? and posi = ? and pic = ?", names, posi, pic).First(i)
}
func GetPeoples(db *gorm.DB, i *[]People) {
	db.Find(i)
}
