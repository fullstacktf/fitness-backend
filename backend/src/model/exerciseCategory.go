package model

// ExerciseCategory model
type ExerciseCategory struct {
	ID   uint8  `gorm:"column:id;type:mediumint unsigned;autoIncrement;primaryKey"`
	Name string `gorm:"column:name;type:varchar(25)"`
}

// TableName Function to change the name of a table. In this case of ExerciseCategory model
func (ec *ExerciseCategory) TableName() string {
	return "exercise_category"
}
