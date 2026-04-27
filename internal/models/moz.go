package models

import (
	u "domains/internal/utils"

	"gorm.io/gorm"
)

type Moz struct {
	gorm.Model
	ID        uint  `gorm:"primaryKey" json:"id"`
	DomainID  uint  `gorm:"uniqueIndex;not null" json:"-"`
	DA        *uint `json:"DA,omitempty"`
	PA        *uint `json:"PA,omitempty"`
	SpamScore *uint `json:"spamScore,omitempty"`
}

func MapToMoz(m map[string]any, target *Moz) {
	daVal := m["DA"]
	daValF := u.ToUint(&daVal)
	if daValF != nil {
		target.DA = daValF
	}

	paVal := m["PA"]
	paValF := u.ToUint(&paVal)
	if paValF != nil {
		target.PA = paValF
	}

	spamVal := m["SpamScore"]
	spamValF := u.ToUint(&spamVal)
	if spamValF != nil {
		target.SpamScore = spamValF
	}
}
