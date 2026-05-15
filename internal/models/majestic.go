package models

import (
	"domains/internal/models/format"

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

func (maj *Majestic) ToMap() map[string]any {
	var res = make(map[string]any)

	res["TF"] = maj.TF
	res["CF"] = maj.CF
	res["Topic"] = maj.Topic

	return res
}
func (maj *Majestic) ToStruct(m map[string]any) {
	format.Format(&maj.TF, m["TF"])
	format.Format(&maj.CF, m["CF"])
	format.Format(&maj.Topic, m["Topic"])
}
