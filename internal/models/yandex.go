package models

import (
	"domains/internal/models/format"

	"gorm.io/gorm"
)

type Yandex struct {
	gorm.Model `json:"-"`
	ID         uint  `gorm:"primaryKey" json:"id"`
	DomainID   uint  `gorm:"uniqueIndex;not null" json:"-"`
	X          *uint `json:"X,omitempty"`
}

func (y *Yandex) ToMap() map[string]any {
	if y == nil {
		return nil
	}

	var res = make(map[string]any)
	res["X"] = y.X
	return res
}

func (y *Yandex) ToStruct(m map[string]any) {
	format.Format(&y.X, m["X"])
}
