package models

import (
	f "domains/internal/models/format"

	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Email    *string `gorm:"uniqueIndex" json:"email"`
	WhatsApp *string `gorm:"uniqueIndex" json:"whatsapp"`
	Telegram *string `gorm:"uniqueIndex" json:"telegram"`
	Phone    *string `gorm:"uniqueIndex" json:"phone"`
	Pbn      *bool   `gorm:"default:false" json:"pbn"`
	Comment  string  `json:"comment"`

	Domains *[]Domain `json:"domains"`

	OutreachID *uint
}

func (c *Contact) ToMap() map[string]any {
	var res = make(map[string]any)

	res["Comment"] = c.Comment
	res["Email"] = c.Email
	res["WhatsApp"] = c.WhatsApp
	res["Telegram"] = c.Telegram
	res["Phone"] = c.Phone
	res["Pbn"] = c.Pbn

	return res
}

func (c *Contact) ToStruct(m map[string]any) {
	f.Format(&c.Comment, m["Comment"])
	f.Format(&c.Email, m["Email"])
	f.Format(&c.WhatsApp, m["WhatsApp"])
	f.Format(&c.Telegram, m["Telegram"])
	f.Format(&c.Phone, m["Phone"])
	f.Format(&c.Pbn, m["Pbn"])
}

func (c *Contact) ReplaceHeaders(h *[]any) {
	var contactMap = map[string]string{
		"Outreach.Email":    "Outreach.Contact.Email",
		"Outreach.WhatsApp": "Outreach.Contact.WhatsApp",
		"Outreach.Telegram": "Outreach.Contact.Telegram",
		"Outreach.Phone":    "Outreach.Contact.Phone",
		"Outreach.Pbn":      "Outreach.Contact.Pbn",

		"Contact.Comment":  "Outreach.Contact.Comment",
		"Contact.Email":    "Outreach.Contact.Email",
		"Contact.WhatsApp": "Outreach.Contact.WhatsApp",
		"Contact.Telegram": "Outreach.Contact.Telegram",
		"Contact.Phone":    "Outreach.Contact.Phone",
		"Contact.Pbn":      "Outreach.Contact.Pbn",
	}

	for i, v := range *h {
		if _, ok := contactMap[v.(string)]; ok {
			(*h)[i] = contactMap[v.(string)]
		}
	}
}
