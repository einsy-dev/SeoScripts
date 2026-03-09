package models

import (
	"gorm.io/gorm"
)

type Ahrefs struct {
	gorm.Model
	DomainID   uint `gorm:"uniqueIndex;not null"`
	SourceID   uint
	DR         *uint
	Traffic    *uint
	Age        *uint
	Geo        *string
	RefDomains *uint
	OutDomains *uint
}
