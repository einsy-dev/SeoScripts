package models

import (
	u "domains/internal/utils"
	"regexp"

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

func MapToAhrefs(m map[string]any, target *Ahrefs) {
	drVal := m["DR"]
	drValF := u.ToUint(&drVal)
	if drValF != nil {
		target.DR = u.ToUint(&drVal)
	}

	ageVal := m["Age"]
	ageValF := u.ToUint(&ageVal)
	if ageValF != nil {
		target.Age = ageValF
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

	geoVal := m["Geo"]
	geoValF := u.ToString(&geoVal)
	if geoValF != nil {
		target.Geo = formatGeo(geoValF)
	}
}

var rGeo = regexp.MustCompile(`[a-zA-Z]+`)

func formatGeo(geo *string) *string {
	var res = rGeo.FindString(*geo)
	if res == "" {
		return nil
	}
	return &res
}
