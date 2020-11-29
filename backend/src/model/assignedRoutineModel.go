package model

import "gorm.io/gorm"

// AssignedRoutine model
type AssignedRoutine struct {
	gorm.Model
	CategoryID      uint64          `gorm:"column:category_id;type:bigint(20) unsigned"`
	UserID          uint64          `gorm:"column:user_id;type:bigint(20) unsigned"`
	Name            string          `gorm:"column:name;type:varchar(15)"`
	Description     string          `gorm:"column:description;type:varchar(30)"`
	RoutineCategory RoutineCategory `gorm:"foreignKey:CategoryID;"`
	User            User            `gorm:"foreignKey:UserID;"`
}

// TableName Function to change the name of a table.
func (ar *AssignedRoutine) TableName() string {
	return "assigned_routines"
}

// CreateAssignedRoutine Create a new assigned routine for a user
func (ar *AssignedRoutine) CreateAssignedRoutine() string {
	return "CreateAssignedRoutine"
}

// GetAssignedRoutine Gets an assigned routine of a user
func (ar *AssignedRoutine) GetAssignedRoutine() string {
	return "GetAssignedRoutine"
}

// GetAssignedRoutines Get all assigned routines of a user
func (ar *AssignedRoutine) GetAssignedRoutines() string {
	return "GetAssignedRoutines"
}

// UpdateAssignedRoutine Update specific assigned routine
func (ar *AssignedRoutine) UpdateAssignedRoutine() string {
	return "UpdateAssignedRoutine"
}

// DeleteAssignedRoutine Delete assigned routine by id
func (ar *AssignedRoutine) DeleteAssignedRoutine() string {
	return "DeleteAssignedRoutine"
}
