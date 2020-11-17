package model

// Muscle model
type Muscle struct {
	ID   uint8  `gorm:"column:id;type:mediumint unsigned;autoIncrement;primaryKey"`
	Name string `gorm:"column:name;type:varchar(20)"`
}

// TableName for RoutineCategory model
func (m *Muscle) TableName() string {
	return "muscles"
}

// CreateMuscle function
func (m *Muscle) CreateMuscle() string {
	return "Create a Muscle"
}

// GetMuscle function
func (m *Muscle) GetMuscle() string {
	return "Get a Muscle"
}

// GetMuscles function
func (m *Muscle) GetMuscles() string {
	return "Get all Muscles"
}

// UpdateMuscle function
func (m *Muscle) UpdateMuscle() string {
	return "Update a Muscle"
}

// DeleteMuscle function
func (m *Muscle) DeleteMuscle() string {
	return "Delete a Muscle"
}
