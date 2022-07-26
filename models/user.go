package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36)" json:"identity"`
	Name     string `gorm:"column:name;type:varchar(100)" json:"name"`
	Password string `gorm:"column:password;type:varchar(32)" json:"email"`
	Phone    string `gorm:"column:phone;type:varchar(20)" json:"phone"`
	Mail     string `gorm:"column:mail;type:varchar(100)" json:"mail"`
}

func (table *User) TableName() string { return "users" }
