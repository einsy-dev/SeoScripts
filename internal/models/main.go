package models

import (
	"domains/internal/app"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Startup() {
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Domain{}, &Ahrefs{})
	app.DB = db
}
