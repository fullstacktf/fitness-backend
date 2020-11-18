package model

import "time"

// Welcome model
type Welcome struct {
	Message string
	Date    time.Time
}

// TableName Function to change the name of a table. In this case of Welcome model
func (w *Welcome) TableName() string {
	return "welcome"
}

// Get of the Welcome model
func (w *Welcome) Get() error {
	welcome := Welcome{
		Message: "Welcome to the youLift API",
		Date:    time.Now(),
	}
	*w = welcome
	return nil
}
