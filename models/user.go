package models

import "github.com/jinzhu/gorm"

// User Model
type User struct {
	gorm.Model
	Username        string `json:"username,omitempty" gorm:"not null;unique"`
	Email           string `json:"email,omitempty" gorm:"not null;unique"`
	FullName        string `json:"fullname,omitempty" gorm:"not null"`
	Password        string `json:"password,omitempty" gorm:"not null;type:varchar(256)"`
	ConfirmPassword string `json:"confirmPassword,omitempty" gorm:"-"`
	Picture         string `json:"picture,omitempty"`
	Level           uint   `json:"level,omitempty"`
	Token           string `json:"token,omitempty"`
}
