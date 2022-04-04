package validate

import (
	"fmt"
	"regexp"
)

type ValidationErr struct {
	FieldValue string
	ErrMassage string
}

func (v *ValidationErr) Error() string {
	return fmt.Sprintf("filed %s, %s", v.FieldValue, v.ErrMassage)
}

func Length(min, max int, str string) string {
	l := len(str)
	if l < min || l > max {
		return fmt.Sprintf("filed should be not less that %v and not greater that %v", min, max)
	}
	return ""
}

func Email(email string) string {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(email) {
		return fmt.Sprintf("incorrect email format")
	}
	return ""
}


