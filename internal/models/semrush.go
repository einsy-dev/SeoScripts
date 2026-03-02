package models

import (
	"gorm.io/gorm"
)

type Semrush struct {
	gorm.Model
	DomainID   uint `gorm:"uniqueIndex;not null"`
	AS         *uint
	Traffic    *uint
	RefDomains *uint
	OutDomains *uint
}
