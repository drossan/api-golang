package models

import "github.com/jinzhu/gorm"

// User Model
type Level struct {
	gorm.Model
	Level string `json:"level,omitempty" gorm:"not null;"`
}
