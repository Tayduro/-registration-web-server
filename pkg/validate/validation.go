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

//func Validate(u User) []ValidationErr {
//	errors := make([]ValidationErr, 0, 0)
//
//	if len(u.FirstName) < 2 {
//		errors = append(errors, ValidationErr{
//			FieldValue: "FirstName",
//			ErrMassage: fmt.Sprintf("field %s length should be equal or longer than 2", "FirstName"),
//		})
//	}
//	if len(u.FirstName) > 64 {
//		errors = append(errors, ValidationErr{
//			FieldValue: "FirstName",
//			ErrMassage: fmt.Sprintf("field %s length should be less than 64", "FirstName"),
//		})
//	}
//
//	if len(u.LastName) < 2 {
//		errors = append(errors, ValidationErr{
//			FieldValue: "LastName",
//			ErrMassage: fmt.Sprintf("field %s length should be equal or longer than 2", "LastName"),
//		})
//	}
//
//	if len(u.LastName) > 64 {
//		errors = append(errors, ValidationErr{
//			FieldValue: "LastName",
//			ErrMassage: fmt.Sprintf("field %s length should be less than 64", "LastName"),
//		})
//	}
//	if len(u.Password) < 8 {
//		errors = append(errors, ValidationErr{
//			FieldValue: "Password",
//			ErrMassage: fmt.Sprintf("field %s length should be equal or longer than 8", "Password"),
//		})
//	}
//
//	if len(u.Password) > 64 {
//		errors = append(errors, ValidationErr{
//			FieldValue: "Password",
//			ErrMassage: fmt.Sprintf("field %s length should be less than 64", "Password"),
//		})
//	}
//

//	return errors
//}
