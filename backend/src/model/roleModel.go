package model

import (
	"gorm.io/gorm"
)

// Role model
type Role struct {
	gorm.Model
	Description string        `gorm:"column:description;type:varchar(20);unique"`
	Permissions []*Permission `gorm:"many2many:roles_permissions;"`
}

// TableName Function to change the name of a table.
func (r *Role) TableName() string {
	return "roles"
}
