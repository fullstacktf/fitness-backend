package model

// BaseExercise model
type BaseExercise struct {
	ID               uint8            `gorm:"column:id;type:mediumint unsigned;autoIncrement;primaryKey"`
	Name             string           `gorm:"column:name;type:varchar(25)"`
	Description      string           `gorm:"column:description;type:varchar(60)"`
	VideoURL         string           `gorm:"column:video_url;type:varchar(50)"`
	CategoryID       uint8            `gorm:"column:category_id;type:mediumint unsigned"`
	Series           uint             `gorm:"column:series;type:tinyint unsigned"`
	Repetitions      uint             `gorm:"column:repetitions;type:tinyint unsigned"`
	ExerciseCategory ExerciseCategory `gorm:"foreignKey:CategoryID;references:ID"`
	Muscles          []*Muscle        `gorm:"many2many:exercises_muscles;"`
}

//TableName function
func (be *BaseExercise) TableName() string {
	return "base_exercises"
}

// CreateBaseExercise function
func (be *BaseExercise) CreateBaseExercise() string {
	return "Create a Base Exercise"
}

//GetBaseExercise function
func (be *BaseExercise) GetBaseExercise() string {
	return "Get a BaseExercise"
}

// GetBaseExercises function
func (be *BaseExercise) GetBaseExercises() string {
	return "Get all BaseExercises"
}

// // GetBaseExercisesByCategoryID function
// func (be *BaseExercise) GetBaseExercisesByCategoryID() string {
// 	return "Get all BaseExercises by category"
// }

// UpdateBaseExercise function
func (be *BaseExercise) UpdateBaseExercise() string {
	return "Update a BaseExxercises"
}

// DeleteBaseExercise function
func (be *BaseExercise) DeleteBaseExercise() string {
	return "Delete a BaseExercises"
}
