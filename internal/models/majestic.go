package models

import "gorm.io/gorm"

type Majestic struct {
	gorm.Model
	DomainID   uint
	TF         *uint
	CF         *uint
	Topic      *string
	RefDomains *uint
	OutDomains *uint
}
