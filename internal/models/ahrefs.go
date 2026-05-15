package models

import (
	f "domains/internal/models/format"

	"gorm.io/gorm"
)

type Ahrefs struct {
	gorm.Model `json:"-"`
	ID         uint    `gorm:"primaryKey" json:"id"`
	DomainID   uint    `gorm:"uniqueIndex;not null" json:"-"`
	DR         *uint   `json:"dr,omitempty"`
	Traffic    *uint   `json:"traffic,omitempty"`
	Age        *uint   `json:"age,omitempty"`
	Geo        *string `json:"geo,omitempty"`
	RefDomains *uint   `json:"refDomains,omitempty"`
	OutDomains *uint   `json:"outDomains,omitempty"`
}

func (a *Ahrefs) ToStruct(m map[string]any) {
	f.Format(&a.DR, m["DR"])
	f.Format(&a.Traffic, m["Traffic"])
	f.Format(&a.Age, m["Age"])
	f.Format(&a.Geo, f.FormatGeo(m["Geo"])) // formats geo
	f.Format(&a.RefDomains, m["RefDomains"])
	f.Format(&a.OutDomains, m["OutDomains"])
}

func (a *Ahrefs) ToMap() map[string]any {
	var res = make(map[string]any)

	res["DR"] = a.DR
	res["Traffic"] = a.Traffic
	res["Age"] = a.Age
	res["Geo"] = a.Geo
	res["RefDomains"] = a.RefDomains
	res["OutDomains"] = a.OutDomains

	return res
}
