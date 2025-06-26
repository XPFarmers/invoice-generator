package database

import (
	"github.com/XPFarmers/invoice-generator/invoices"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewDBConnection() *gorm.DB {
	var err error
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&invoices.Invoice{}, &invoices.LineItem{})

	return db
}
