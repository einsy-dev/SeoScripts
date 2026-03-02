package models

import (
	"gorm.io/gorm"
)

type Majestic struct {
	gorm.Model
	DomainID   uint `gorm:"uniqueIndex;not null"`
	TF         *uint
	CF         *uint
	Topic      *string
	RefDomains *uint
	OutDomains *uint
}
