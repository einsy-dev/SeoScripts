package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Token      *string `json:"token" gorm:"uniqueIndex;not null"`
	Name       *string `json:"name"`
	Role       string  `json:"rolw" gorm:"default:USER"`
	Status     *string `json:"status"`
}

func (u *User) ToMap() map[string]any {
	var res = make(map[string]any)
	return res
}
func (u *User) ToStruct(m map[string]any) {}
