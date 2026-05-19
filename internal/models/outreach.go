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

	Contact   *Contact
	ContactID *uint
	DomainID  *uint `gorm:"uniqueIndex"`
}

func (o *Outreach) ToMap() map[string]any {
	if o == nil {
		return nil
	}
	var res = make(map[string]any)

	res["Comment"] = o.Comment
	res["Context"] = o.Context
	res["Note"] = o.Note
	res["Article"] = o.Article
	res["Image"] = o.Image

	res["Contact"] = o.Contact.ToMap()

	return res
}

func (o *Outreach) ToStruct(m map[string]any) {

	f.Format(&o.Comment, m["Comment"])
	f.Format(&o.Context, m["Context"])
	f.Format(&o.Note, m["Note"])
	f.Format(&o.Article, m["Article"])
	f.Format(&o.Image, m["Image"])

	if s, ok := m["Contact"].(map[string]any); ok {
		if o.Contact == nil || check(o.Contact, s) {
			o.Contact = handleContact(s)
		}
		if o.Contact != nil {
			o.Contact.ToStruct(s)
		}
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
		return nil
	}

	query.FirstOrCreate(c)

	return c
}

func check(c *Contact, m map[string]any) bool {
	match := func(mapKey string, ptr *string) bool {
		val, ok := m[mapKey].(string)
		if !ok {
			return ptr == nil // match if both map value doesn't exist and db string is nil
		}
		return ptr != nil && *ptr == val
	}

	if match("Email", c.Email) || match("WhatsApp", c.WhatsApp) || match("Telegram", c.Telegram) || match("Phone", c.Phone) {
		return false
	}

	return true
}
