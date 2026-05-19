package models

import (
	"domains/internal/models/format"

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

func (moz *Moz) ToMap() map[string]any {
	if moz == nil {
		return nil
	}
	var res = make(map[string]any)

	res["DA"] = moz.DA
	res["PA"] = moz.PA
	res["SpamScore"] = moz.SpamScore

	return res
}

func (moz *Moz) ToStruct(m map[string]any) {
	format.Format(&moz.DA, m["DA"])
	format.Format(&moz.PA, m["PA"])
	format.Format(&moz.SpamScore, m["SpamScore"])
}
