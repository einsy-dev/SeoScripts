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
