package model

import "gorm.io/gorm"

// RoutineCategory model
type RoutineCategory struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(25)"`
}

// TableName Function to change the name of a table. In this case of RoutineCategory model
func (rc *RoutineCategory) TableName() string {
	return "routine_category"
}
