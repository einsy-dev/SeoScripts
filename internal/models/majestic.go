package models

import (
	"gorm.io/gorm"
)

type Majestic struct {
	gorm.Model
	DomainID uint `gorm:"uniqueIndex;not null"`
	SourceID uint
	TF       *uint
	CF       *uint
	Topic    *string
}
