package models

import (
	u "domains/internal/utils"

	"gorm.io/gorm"
)

type Majestic struct {
	gorm.Model `json:"-"`
	ID         uint    `gorm:"primaryKey" json:"id"`
	DomainID   uint    `gorm:"uniqueIndex;not null" json:"-"`
	TF         *uint   `json:"tf,omitempty"`
	CF         *uint   `json:"cf,omitempty"`
	Topic      *string `json:"topic,omitempty"`
}

func MapToMajestic(m map[string]any, target *Majestic) {
	tfVal := m["TF"]
	tfValF := u.ToUint(&tfVal)
	if tfValF != nil {
		target.TF = tfValF
	}

	cfVal := m["CF"]
	cfValF := u.ToUint(&cfVal)
	if cfValF != nil {
		target.CF = cfValF
	}

	topicVal := m["Topic"]
	topicValF := u.ToString(&topicVal)
	if topicValF != nil {
		target.Topic = topicValF
	}
}
