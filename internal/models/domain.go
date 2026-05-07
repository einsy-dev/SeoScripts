package models

import (
	u "domains/internal/utils"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model `json:"-"`
	ID         uint       `gorm:"primaryKey" json:"id"`
	Domain     string     `gorm:"uniqueIndex;not null" json:"domain"`
	Type       *string    `json:"type,omitempty"`
	Rel        *string    `json:"rel,omitempty"`
	Comment    string     `json:"comment"`
	Ahrefs     *Ahrefs    `json:"ahrefs,omitempty"`
	Semrush    *Semrush   `json:"semrush,omitempty"`
	Majestic   *Majestic  `json:"majestic,omitempty"`
	Moz        *Moz       `json:"moz,omitempty"`
	Accounts   *[]Account `json:"accounts,omitempty"`
	Links      *[]Link    `json:"links,omitempty"`
	Group      *Group     `json:"group,omitempty"`
	GroupID    *uint      `json:"-"`
}

func MapToDomain(m map[string]any, target *Domain) {
	commentVal := m["Comment"]
	commentValF := u.ToString(&commentVal)
	if commentValF != nil {
		target.Comment = *commentValF
	}

	typeVal := m["Type"]
	typeValF := u.ToString(&typeVal)
	if typeValF != nil {
		target.Type = typeValF
	}

	relVal := m["Rel"]
	relValF := u.ToString(&relVal)
	if relValF != nil {
		target.Rel = relValF
	}

	if a, ok := m["Ahrefs"].(map[string]any); ok {
		if target.Ahrefs == nil {
			target.Ahrefs = &Ahrefs{}
		}
		MapToAhrefs(a, target.Ahrefs)
	}
	if s, ok := m["Semrush"].(map[string]any); ok {
		if target.Semrush == nil {
			target.Semrush = &Semrush{}
		}
		MapToSemrush(s, target.Semrush)
	}
	if maj, ok := m["Majestic"].(map[string]any); ok {
		if target.Majestic == nil {
			target.Majestic = &Majestic{}
		}
		MapToMajestic(maj, target.Majestic)
	}
	if moz, ok := m["Moz"].(map[string]any); ok {
		if target.Moz == nil {
			target.Moz = &Moz{}
		}
		MapToMoz(moz, target.Moz)
	}
}

func DomainToMap(domain *Domain) map[string]any {
	return nil
}

func DomainToCsv(d *Domain) [][]any {
	return [][]any{
		{
			"Domain",
			"Type",
			"Rel",
			"Comment",
			"Ahrefs.DR",
			"Ahrefs.Traffic",
			"Ahrefs.Age",
			"Ahrefs.Geo",
			"Ahrefs.RefDomains",
			"Ahrefs.OutDomains",
			"Semrush.AS",
			"Semrush.Traffic",
			"Semrush.RefDomains",
			"Semrush.OutDomains",
			"Semrush.LinkFarm",
			"Majestic.TF",
			"Majestic.CF",
			"Majestic.Topic",
		},
		{
			d.Domain,
			checkNil(d.Type),
			checkNil(d.Rel),
			d.Comment,
			checkNil(d.Ahrefs.DR),
			checkNil(d.Ahrefs.Traffic),
			checkNil(d.Ahrefs.Age),
			checkNil(d.Ahrefs.Geo),
			checkNil(d.Ahrefs.RefDomains),
			checkNil(d.Ahrefs.OutDomains),
			checkNil(d.Semrush.AS),
			checkNil(d.Semrush.Traffic),
			checkNil(d.Semrush.RefDomains),
			checkNil(d.Semrush.OutDomains),
			checkNil(d.Semrush.LinkFarm),
			checkNil(d.Majestic.TF),
			checkNil(d.Majestic.CF),
			checkNil(d.Majestic.Topic),
		},
	}
}

func checkNil(P any) any {
	if P == nil {
		return ""
	}
	return P
}
