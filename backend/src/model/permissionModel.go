package model

import (
	"gorm.io/gorm"
)

// Permission model
type Permission struct {
	gorm.Model
	Description string `gorm:"column:description;type:varchar(30);unique"`
}

// TableName Function to change the name of a table.
func (p *Permission) TableName() string {
	return "permissions"
}
