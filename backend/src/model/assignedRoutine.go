package model

// AssignedRoutine model
type AssignedRoutine struct {
	ID              uint8           `gorm:"column:id;type:mediumint unsigned;autoIncrement;primaryKey"`
	CategoryID      uint8           `gorm:"column:category_id;type:mediumint unsigned"`
	UserID          uint8           `gorm:"column:user_id;type:mediumint unsigned"`
	Name            string          `gorm:"column:name;type:varchar(15)"`
	Description     string          `gorm:"column:description;type:varchar(30)"`
	RoutineCategory RoutineCategory `gorm:"foreignKey:CategoryID;references:ID"`
	User            User            `gorm:"foreignKey:UserID;references:ID"`
}
