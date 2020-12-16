package model

import "gorm.io/gorm"

// Muscle model
type Muscle struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(20)"`
}

// TableName Function to change the name of a table.
func (m *Muscle) TableName() string {
	return "muscles"
}
