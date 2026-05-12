package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model `json:"-"`
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"not null" json:"title"`
	Comment    string    `gorm:"not null" json:"comment"`
	Domains    *[]Domain `gorm:"many2many:domain_groups;" json:"domains"`
}

func (g *Group) ToMap() map[string]any {
	var res = make(map[string]any)
	return res
}

func (g *Group) ToStruct(m map[string]any) {}
