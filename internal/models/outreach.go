package models

import (
	"domains/internal/app"
	f "domains/internal/models/format"

	"gorm.io/gorm"
)

type Outreach struct {
	gorm.Model
	Comment string
	Context *string
	Note    *string
	Article *string
	Image   *string

	Contact *Contact

	DomainID *uint `gorm:"uniqueIndex"`
}

func (o *Outreach) ToMap() map[string]any {
	var res = make(map[string]any)

	res["Comment"] = o.Comment
	res["Context"] = o.Context
	res["Note"] = o.Note
	res["Article"] = o.Article
	res["Image"] = o.Image

	if o.Contact != nil {
		res["Contact"] = o.Contact.ToMap()
	} else {
		res["Contact"] = nil
	}

	return res
}

func (o *Outreach) ToStruct(m map[string]any) {

	f.Format(&o.Comment, m["Comment"])
	f.Format(&o.Context, m["Context"])
	f.Format(&o.Note, m["Note"])
	f.Format(&o.Article, m["Article"])
	f.Format(&o.Image, m["Image"])

	if s, ok := m["Contact"].(map[string]any); ok {
		if o.Contact == nil {
			o.Contact = handleContact(s)
		}
		o.Contact.ToStruct(s)
	}
}

func handleContact(m map[string]any) *Contact {
	c := &Contact{}
	query := app.DB.Model(&Contact{})

	hasIdentifier := false

	if val, ok := m["Email"].(string); ok && val != "" {
		query = query.Or("email = ?", val)
		hasIdentifier = true
	}
	if val, ok := m["WhatsApp"].(string); ok && val != "" {
		query = query.Or("whats_app = ?", val)
		hasIdentifier = true
	}
	if val, ok := m["Telegram"].(string); ok && val != "" {
		query = query.Or("telegram = ?", val)
		hasIdentifier = true
	}
	if val, ok := m["Phone"].(string); ok && val != "" {
		query = query.Or("phone = ?", val)
		hasIdentifier = true
	}
	if !hasIdentifier {
		return c
	}

	query.FirstOrCreate(c)

	return c
}
