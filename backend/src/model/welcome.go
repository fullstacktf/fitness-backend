package model

import "time"

// Welcome struct
type Welcome struct {
	Message string
	Date    time.Time
}

// TableName for Welcome model
func (w *Welcome) TableName() string {
	return "welcome"
}

// Get of the Welcome struct
func (w *Welcome) Get() error {
	welcome := Welcome{
		Message: "Welcome to the youLift API",
		Date:    time.Now(),
	}
	*w = welcome
	return nil
}
