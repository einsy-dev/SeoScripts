package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	token *string `gorm:"uniqueIndex;not null"`
	name  *string
	role  string
}
