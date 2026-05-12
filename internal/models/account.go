package models

type Account struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	DomainID uint    `gorm:"uniqueIndex;not null" json:"-"`
	Email    *string `gorm:"uniqueIndex;not null" json:"email"`
	Password *string `json:"password"`
	Username *string `json:"username"`
}

func (a *Account) ToMap() map[string]any {
	var res = make(map[string]any)
	return res
}
func (a *Account) ToStruct(m map[string]any) {

}
