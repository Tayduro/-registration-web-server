package models

const (
	MinNameLength = 2
	MaxNameLength = 255
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}
