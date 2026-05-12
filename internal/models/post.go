package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model `json:"-"`
	Price      string // normal price
	Adult      string // price for adult / casino etc
	Comment    string
}

func (p *Post) ToMap() map[string]any {
	var res = make(map[string]any)
	return res
}
func (p *Post) ToStruct(m map[string]any) {

}
