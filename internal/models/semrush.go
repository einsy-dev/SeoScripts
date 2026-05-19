package models

import (
	"domains/internal/models/format"

	"gorm.io/gorm"
)

type Semrush struct {
	gorm.Model `json:"-"`
	ID         uint    `gorm:"primaryKey" json:"id"`
	DomainID   uint    `gorm:"uniqueIndex;not null" json:"-"`
	AS         *uint   `json:"as,omitempty"`
	Traffic    *uint   `json:"traffic,omitempty"`
	RefDomains *uint   `json:"refDomains,omitempty"`
	OutDomains *uint   `json:"outDomains,omitempty"`
	LinkFarm   *string `json:"linkFarm,omitempty"`
}

func (s *Semrush) ToStruct(m map[string]any) {
	format.Format(&s.AS, m["AS"])
	format.Format(&s.Traffic, m["Traffic"])
	format.Format(&s.RefDomains, m["RefDomains"])
	format.Format(&s.OutDomains, m["OutDomains"])
	format.Format(&s.LinkFarm, m["LinkFarm"])
}

func (s *Semrush) ToMap() map[string]any {
	if s == nil {
		return nil
	}
	var res = make(map[string]any)
	res["AS"] = s.AS
	res["Traffic"] = s.Traffic
	res["RefDomains"] = s.RefDomains
	res["OutDomains"] = s.OutDomains
	res["LinkFarm"] = s.LinkFarm
	return res
}
