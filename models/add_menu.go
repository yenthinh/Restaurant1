package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"not null;size:150" form:"name"`
	Ingredients string `gorm:"not null;size:300" form:"inger"`
	Price       string `gorm:"not null;size:100" form:"price"`
	Hide 		bool  `form:"hide"`
}

func CreateProduct(db *gorm.DB, i Product) error {
	return db.Create(&i).Error
}
func Update(db *gorm.DB, i *Product, id int) {
	db.Where("id = ?", id).Updates(i)
}
func UpdateHide(db *gorm.DB, id int, hide bool) {
	db.Select("Hide").Where("id = ?", id).Updates(Product{Hide: hide})
}
func Delete(db *gorm.DB, i Product, id int) {
	db.Where("id = ?", id).Delete(&i)
}
func DeleteHide(db *gorm.DB, id int, hide bool) {
	db.Select("Hide").Where("id = ?", id).Updates(Product{Hide: hide})
}
func GetProduct(db *gorm.DB, i *Product, name, inger, price string) {
	db.Where("name = ? and inger = ? and price = ?", name, inger, price).First(i)
}
func GetProducts(db *gorm.DB, i *[]Product, getHide bool) {
	db.Where("hide= ?", getHide).Find(i)
}
func GetAllProduct(db *gorm.DB, i *[]Product) {
	db.Find(i)
}
