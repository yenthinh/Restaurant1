package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB() *gorm.DB {
	connectionString := "root:Thinhyen2801@tcp(127.0.0.1:3306)/restaurant?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}
func Migrate() {
	db := OpenDB()
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&People{})
	db.AutoMigrate(&Arrival{})
	CloseDB(db)
}
