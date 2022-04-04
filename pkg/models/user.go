package models

const (
	MinNameLength = 2
	MaxNameLength = 255
)

type User struct {
	Id string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UserData struct {
	Field string
	FieldValue string
}

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

type Credentials struct {
	Id string `db:"user_id"`
	Salt string `db:"salt"`
	Hash  string `db:"hash"`
}

