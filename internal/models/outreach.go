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

func isNew(c *Contact, m map[string]any) bool {
	match := func(mapKey string, ptr *string) bool {
		val, ok := m[mapKey].(string)
		if !ok || ptr == nil {
			return false
		}
		return *ptr == val
	}

	if match("Email", c.Email) ||
		match("WhatsApp", c.WhatsApp) ||
		match("Telegram", c.Telegram) ||
		match("Phone", c.Phone) {
		return false
	}

	return true
}

func (o *Outreach) ToStruct(m map[string]any) {

	f.Format(&o.Comment, m["Comment"])
	f.Format(&o.Context, m["Context"])
	f.Format(&o.Note, m["Note"])
	f.Format(&o.Article, m["Article"])
	f.Format(&o.Image, m["Image"])

	if s, ok := m["Contact"].(map[string]any); ok {
		if o.Contact == nil || isNew(o.Contact, s) {
			o.Contact = handleContact(s)
		}
		if o.Contact != nil {
			o.Contact.ToStruct(s)
			o.ContactID = &o.Contact.ID
		}

	}
}

func handleContact(m map[string]any) *Contact {
	c := &Contact{}

	// Create a structural query to find any existing record matching the identifiers
	var conditions []string
	var values []any

	if val, ok := m["Email"].(string); ok && val != "" {
		conditions = append(conditions, "email = ?")
		values = append(values, val)
	}
	if val, ok := m["WhatsApp"].(string); ok && val != "" {
		conditions = append(conditions, "whats_app = ?")
		values = append(values, val)
	}
	if val, ok := m["Telegram"].(string); ok && val != "" {
		conditions = append(conditions, "telegram = ?")
		values = append(values, val)
	}
	if val, ok := m["Phone"].(string); ok && val != "" {
		conditions = append(conditions, "phone = ?")
		values = append(values, val)
	}

	if len(conditions) == 0 {
		return nil
	}

	queryStr := conditions[0]
	for i := 1; i < len(conditions); i++ {
		queryStr += " OR " + conditions[i]
	}
	err := app.DB.Where(queryStr, values...).First(c).Error
	if err != nil {
		if createErr := app.DB.Create(c).Error; createErr != nil {
			return nil
		}
	}
	return c
}
