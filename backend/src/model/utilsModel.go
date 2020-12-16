package model

// ErrorVO Error model for sending back an error code to the front app
type ErrorVO struct {
	Code    int
	Message string
}

// Pagedrequest Model for paged requests
type Pagedrequest struct {
	Size   int
	Offset int
}
