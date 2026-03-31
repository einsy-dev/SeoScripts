package models

import (
	u "domains/internal/utils"

	"gorm.io/gorm"
)

type Semrush struct {
	gorm.Model `json:"-"`
	ID         uint    `gorm:"primaryKey" json:"id"`
	DomainID   uint    `gorm:"uniqueIndex;not null" json:"-"`
	AS         *uint   `json:"AS,omitempty"`
	Traffic    *uint   `json:"traffic,omitempty"`
	RefDomains *uint   `json:"refDomians,omitempty"`
	OutDomains *uint   `json:"outDomains,omitempty"`
	LinkFarm   *string `json:"linkFarm,omitempty"`
}

func MapToSemrush(m map[string]any, target *Semrush) {
	asVal := m["AS"]
	asValF := u.ToUint(&asVal)
	if asValF != nil {
		target.AS = asValF
	}

	trafficVal := m["Traffic"]
	trafficValF := u.ToUint(&trafficVal)
	if trafficValF != nil {
		target.Traffic = trafficValF
	}

	refDVal := m["RefDomains"]
	refDValF := u.ToUint(&refDVal)
	if refDValF != nil {
		target.RefDomains = refDValF
	}

	outDVal := m["OutDomains"]
	outDValF := u.ToUint(&outDVal)
	if outDValF != nil {
		target.OutDomains = outDValF
	}

	linkFarmVal := m["LinkFarm"]
	linkFarmValF := u.ToString(&linkFarmVal)
	if linkFarmValF != nil {
		target.LinkFarm = linkFarmValF
	}
}
