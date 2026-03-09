package models

import "gorm.io/gorm"

type Source struct {
	gorm.Model
	app    *string
	title  *string
	userID uint
}
