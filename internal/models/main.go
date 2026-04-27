package models

import (
	"domains/internal/app"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Startup() {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Domain{}, &Ahrefs{}, &Semrush{}, &Majestic{}, &Moz{}, &Account{}, &Link{}, &Group{}, &User{})

	db.Exec("CREATE TYPE status AS ENUM ('active', 'inactive', 'pending')")
	db.Exec("CREATE TYPE rel AS ENUM ('follow', 'nofollow')")

	app.DB = db
}
