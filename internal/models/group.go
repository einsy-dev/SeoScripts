package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model `json:"-"`
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"not null"`
	Domains    []Domain
}
