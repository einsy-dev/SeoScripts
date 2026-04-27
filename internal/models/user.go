package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Token      *string `gorm:"uniqueIndex;not null" json:"token"`
	Name       *string `json:"name"`
	Role       string  `gorm:"default:USER"`
	Status     *string `json:"status"`
}
