package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	name    *string
	token   *string
	sources []Source
}
