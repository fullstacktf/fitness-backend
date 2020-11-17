package model

// RoutineCategory model
type RoutineCategory struct {
	ID   uint8  `gorm:"column:id;type:mediumint unsigned;autoIncrement;primaryKey"`
	Name string `gorm:"column:name;type:varchar(25)"`
}

// TableName for RoutineCategory model
func (rc *RoutineCategory) TableName() string {
	return "routine_category"
}

// CreateRoutineCategory function
func (rc *RoutineCategory) CreateRoutineCategory() string {
	return "Create a Routinecategory"
}

//GetRoutineCategory function
func (rc *RoutineCategory) GetRoutineCategory() string {
	return "Get a Routinecategory"
}

// GetRoutineCategories function
func (rc *RoutineCategory) GetRoutineCategories() string {
	return "Get all Routinecategory"
}

// UpdateRoutineCategory function
func (rc *RoutineCategory) UpdateRoutineCategory() string {
	return "Update a Routinecategory"
}

// DeleteRoutineCategory function
func (rc *RoutineCategory) DeleteRoutineCategory() string {
	return "Delete a Routinecategory"
}
