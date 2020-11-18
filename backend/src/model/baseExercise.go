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
