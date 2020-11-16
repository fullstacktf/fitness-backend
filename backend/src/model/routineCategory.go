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
