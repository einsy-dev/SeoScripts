package models

import (
	f "domains/internal/models/format"

	"gorm.io/gorm"
)

// sungle table slices fields must be separat tables that handle bulk values

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
	Outreach *Outreach `json:"outreach,omitempty"`

	Links    *[]Link    `json:"links,omitempty"`
	Accounts *[]Account `json:"accounts,omitempty"`
	Groups   *[]Group   `gorm:"many2many:domain_groups;" json:"group,omitempty"`

	ContactID *uint
}

func (d *Domain) ToMap() map[string]any {
	if d == nil {
		return nil
	}

	var res = make(map[string]any)

	res["Domain"] = d.Domain
	res["Type"] = d.Type
	res["Comment"] = d.Comment

	res["Ahrefs"] = d.Ahrefs.ToMap()
	res["Semrush"] = d.Semrush.ToMap()
	res["Majestic"] = d.Majestic.ToMap()
	res["Moz"] = d.Moz.ToMap()
	res["Yandex"] = d.Yandex.ToMap()
	res["Outreach"] = d.Outreach.ToMap()

	return res
}

func (d *Domain) ToStruct(m map[string]any) {

	f.Format(&d.Domain, m["Domain"])
	f.Format(&d.Type, m["Type"])
	f.Format(&d.Comment, m["Comment"])

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

	if y, ok := m["Yandex"].(map[string]any); ok {
		if d.Yandex == nil {
			d.Yandex = &Yandex{}
		}
		d.Yandex.ToStruct(y)
	}

	if y, ok := m["Outreach"].(map[string]any); ok {
		if d.Outreach == nil {
			d.Outreach = &Outreach{}
		}
		d.Outreach.ToStruct(y)
	}
}
