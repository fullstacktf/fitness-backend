package model

// BaseRoutine model
type BaseRoutine struct {
	ID              uint8           `gorm:"column:id;type:mediumint unsigned;autoIncrement;primaryKey"`
	CategoryID      uint8           `gorm:"column:category_id;type:mediumint unsigned"`
	Name            string          `gorm:"column:name;type:varchar(15)"`
	Description     string          `gorm:"column:description;type:varchar(30)"`
	RoutineCategory RoutineCategory `gorm:"foreignKey:CategoryID;references:ID"`
	BaseExercises   []*BaseExercise `gorm:"many2many:base_routines_base_exercises"`
}

//TableName function
func (br *BaseRoutine) TableName() string {
	return "base_exercises"
}

// CreateBaseRoutine function
func (br *BaseRoutine) CreateBaseRoutine() string {
	return "Create a Base Routine"
}

//GetBaseRoutine function
func (br *BaseRoutine) GetBaseRoutine() string {
	return "Get a BaseRoutine"
}

// GetBaseRoutines function
func (br *BaseRoutine) GetBaseRoutines() string {
	return "Get all BaseExercises"
}

// UpdateBaseRoutine function
func (br *BaseRoutine) UpdateBaseRoutine() string {
	return "Update a Base Routine"
}

// DeleteBaseRoutine function
func (br *BaseRoutine) DeleteBaseRoutine() string {
	return "Delete a Base Routine"
}
