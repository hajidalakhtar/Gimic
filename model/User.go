package model

import "github.com/jinzhu/gorm"

// User : users table
type User struct {
	gorm.Model
	Fullname string `form:"fullname" json:"fullname" gorm:"size:100"`
	Email    string `form:"email" json:"email" gorm:"size:100"`
	Phone    string `form:"phone" json:"phone" gorm:"size:13"`
}

