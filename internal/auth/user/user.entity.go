package user

import (
	c "form_management/common/type"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Code     string `gorm:"unique;not null"`
	Role     c.Role `gorm:"not null;default:GUEST"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type UserInfo struct {
	ID    uint
	Role  c.Role
	Email string
	Code  string
}
