package model

// ExerciseCategory model
type ExerciseCategory struct {
	ID   uint8  `gorm:"column:id;type:mediumint unsigned;autoIncrement;primaryKey"`
	Name string `gorm:"column:name;type:varchar(25)"`
}

// TableName for RoutineCategory model
func (ec *ExerciseCategory) TableName() string {
	return "exercise_category"
}
