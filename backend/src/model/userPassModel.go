package model

// UserPass model
type UserPass struct {
	UserID   uint8  `gorm:"column:user_id;type:mediumint unsigned"`
	Password string `gorm:"column:password;type:varchar(50)"`
	User     User   `gorm:"foreignKey:UserID;references:ID"`
}

// TableName for UserPass model
func (up *UserPass) TableName() string {
	return "user_pass"
}

// UpdateUserPass Updates the users pass
func (up *UserPass) UpdateUserPass() string {
	return "UpdateUserPass"
}
