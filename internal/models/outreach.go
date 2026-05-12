package models

import "gorm.io/gorm"

type Outreach struct {
	gorm.Model
	Email    *string
	WhatsApp *string
	Telegram *string
	Phone    *string
	Domains  *[]Domain
}
