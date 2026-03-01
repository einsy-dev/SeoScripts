package models

import "gorm.io/gorm"

type Domain struct {
	gorm.Model
	Domain   string `gorm:"uniqueIndex"`
	Type     *string
	Rel      *string
	Comment  *string
	Ahrefs   Ahrefs
	Semrush  Semrush
	Majestic Majestic
}
