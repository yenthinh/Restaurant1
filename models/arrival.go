package models

import "gorm.io/gorm"

type Arrival struct {
	gorm.Model
	Uname    string `gorm:"not null;size:150" form:"uname"`
	Umail    string `gorm:"not null;size:150" form:"umail"`
	Uphone   string `gorm:"not null;size:15" form:"uphone"`
	Date     string `gorm:"not null;size:15" form:"date"`
	Time     string `gorm:"not null;size:150" form:"time"`
	Ofpeople string `gorm:"not null;size:30" form:"ofpeo"`
	Message  string `gorm:"not null;size:300" form:"mess"`
	Check    bool   `form:"check"`
}

func CreateArrival(db *gorm.DB, i Arrival) error {
	return db.Create(&i).Error
}
func UpdaTe(db *gorm.DB, i *Arrival, id int) {
	db.Where("id = ?", id).Updates(i)
}
func UpdaTeCheck(db *gorm.DB, id int, check bool) {
	db.Select("Check").Where("id = ?", id).Updates(Arrival{Check: check})
}
func DeleTe(db *gorm.DB, i Arrival, id int) {
	db.Where("id = ?", id).Delete(&i)
}
func DeleTeCheck(db *gorm.DB, id int, check bool){
	db.Select("Check").Where("id = ?", id).Updates(Arrival{Check: check})
}

func GetArrival(db *gorm.DB, i *Arrival, uname, umail, uphone, date, time, ofpeo, mess string) {
	db.Where("uname = ? umail = ? uphone = ? date = ? time = ? ofpeo = ? mess = ?").First(i)
}
func GetArrivals(db *gorm.DB, i *[]Arrival, getCheck bool) {
	db.Find(i)
}
