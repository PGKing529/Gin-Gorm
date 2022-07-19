package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(36)" json:"identity"`
	CategoryId string `gorm:"column:category_id;type:varchar(255)" json:"category_id"`
	Title      string `gorm:"column:title;type:varchar(255)" json:"title"`
	Content    string `gorm:"column:content;type:text" json:"content"`
	MaxRuntime int    `gorm:"column:max_runtime;type:int(11)" json:"max_runtime"`
	MaxMemory  int    `gorm:"column:max_mem;type:int(11)" json:"max_memory"`
}

func (table *Problem) TableName() string {
	return "problems"
}
