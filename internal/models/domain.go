package models

import (
	f "domains/internal/models/format"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model `json:"-"`
	ID         uint    `gorm:"primaryKey" json:"id"`
	Domain     string  `gorm:"uniqueIndex;not null" json:"domain"`
	Type       *string `json:"type,omitempty"`
	Comment    string  `json:"comment"`

	Ahrefs   *Ahrefs   `json:"ahrefs,omitempty"`
	Semrush  *Semrush  `json:"semrush,omitempty"`
	Majestic *Majestic `json:"majestic,omitempty"`
	Moz      *Moz      `json:"moz,omitempty"`
	Yandex   *Yandex   `json:"yandex,omitempty"`

	Accounts *[]Account `json:"accounts,omitempty"`
	Links    *[]Link    `json:"links,omitempty"`
	Groups   *[]Group   `gorm:"many2many:domain_groups;" json:"group,omitempty"`

	Outreach   *Outreach `json:"outreach,omitempty"`
	OutreachID *uint
}

func (d *Domain) ToMap() map[string]any {
	var res = make(map[string]any)

	res["Domain"] = d.Domain
	res["Type"] = d.Type
	res["Comment"] = d.Comment

	res["Ahrefs"] = d.Ahrefs.ToMap()
	res["Semrush"] = d.Semrush.ToMap()
	res["Majestic"] = d.Majestic.ToMap()
	res["Moz"] = d.Moz.ToMap()
	res["Yandex"] = d.Yandex.ToMap()

	return res
}

func (d *Domain) ToStruct(m map[string]any) {

	f.Format(d.Domain, m["Domain"])
	f.Format(d.Type, m["Type"])
	f.Format(d.Comment, m["Comment"])

	if a, ok := m["Ahrefs"].(map[string]any); ok {
		if d.Ahrefs == nil {
			d.Ahrefs = &Ahrefs{}
		}
		d.Ahrefs.ToStruct(a)
	}

	if s, ok := m["Semrush"].(map[string]any); ok {
		if d.Semrush == nil {
			d.Semrush = &Semrush{}
		}
		d.Semrush.ToStruct(s)
	}

	if maj, ok := m["Majestic"].(map[string]any); ok {
		if d.Majestic == nil {
			d.Majestic = &Majestic{}
		}
		d.Majestic.ToStruct(maj)
	}

	if moz, ok := m["Moz"].(map[string]any); ok {
		if d.Moz == nil {
			d.Moz = &Moz{}
		}
		d.Moz.ToStruct(moz)
	}
}
