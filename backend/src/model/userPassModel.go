package model

// UserPass model
type UserPass struct {
	UserID   uint64 `gorm:"column:user_id;type:bigint(20) unsigned; unique"`
	Password string `gorm:"column:password;type:varchar(50)"`
	User     User   `gorm:"foreignKey:UserID;"`
}

// TableName for UserPass model
func (up *UserPass) TableName() string {
	return "user_pass"
}

// UpdateUserPass Updates the users pass
func (up *UserPass) UpdateUserPass() string {
	return "UpdateUserPass"
}
