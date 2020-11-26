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

// CreateMuscle Create a muscle group
func (m *Muscle) CreateMuscle() string {
	return "Create a Muscle"
}

// GetMuscle gets a muscle group
func (m *Muscle) GetMuscle() string {
	return "Get a Muscle"
}

// GetMuscles gets all muscle groups
func (m *Muscle) GetMuscles() string {
	return "Get all Muscles"
}

// UpdateMuscle Updates a muscle group
func (m *Muscle) UpdateMuscle() string {
	return "Update a Muscle"
}

// DeleteMuscle Deletes a muscle group
func (m *Muscle) DeleteMuscle() string {
	return "Delete a Muscle"
}
