package models

import "gorm.io/gorm"

type Semrush struct {
	gorm.Model
	DomainID   uint
	AS         *uint
	Traffic    *uint
	RefDomains *uint
	OutDomains *uint
}
