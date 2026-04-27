package models

type Account struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	DomainID uint    `gorm:"uniqueIndex;not null" json:"-"`
	Email    *string `gorm:"uniqueIndex;not null" json:"email"`
	Password *string `json:"password"`
	Username *string `json:"username"`
}
