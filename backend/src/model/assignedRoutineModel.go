package model

import (
	"gorm.io/gorm"
)

// AssignedRoutine model
type AssignedRoutine struct {
	gorm.Model
	UserID        uint64      `gorm:"column:user_id;type:bigint(20) unsigned"`
	Name          string      `gorm:"column:name;type:varchar(25)"`
	Description   string      `gorm:"column:description;type:varchar(600)"`
	BaseRoutineID uint64      `gorm:"column:base_routine_id;type:bigint(20) unsigned"`
	BaseRoutine   BaseRoutine `gorm:"foreignKey:BaseRoutineID;"`
	User          User        `gorm:"foreignKey:UserID;"`
}

// TableName Function to change the name of a table.
func (ar *AssignedRoutine) TableName() string {
	return "assigned_routines"
}
