package models

import (
	"domains/internal/models/enums"
	"time"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model `json:"-"`
	ID         uint          `gorm:"primaryKey" json:"id"`
	DomainID   uint          `gorm:"not null" json:"-"`
	LinkPage   *string       `json:"linkPage"`
	TargetUrl  *string       `json:"targetUrl"`
	Anchor     *string       `json:"anchor"`
	Rel        *string       `json:"rel"`
	ReviewedAt *time.Time    `json:"reviewedAt"`
	Status     *enums.Status `gorm:"type:status;default:'pending'" json:"status"`
}
