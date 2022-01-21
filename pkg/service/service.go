package service

import (
	"github.com/Tayduro/registration-web-server/pkg/models"
	"github.com/Tayduro/registration-web-server/pkg/validate"
)

func Signup(user *models.User) []validate.ValidationErr {
	Errors := make([]validate.ValidationErr, 0, 0)
	if validate.Length(models.MinNameLength, models.MaxNameLength, user.FirstName) != "" {
		Errors = append(Errors, validate.ValidationErr{
			FieldValue: "FirstName",
			ErrMassage: validate.Length(models.MinNameLength, models.MaxNameLength, user.FirstName),
		})

	}
	if validate.Length(models.MinNameLength, models.MaxNameLength, user.LastName) != "" {
		Errors = append(Errors, validate.ValidationErr{
			FieldValue: "LastName",
			ErrMassage: validate.Length(models.MinNameLength, models.MaxNameLength, user.LastName),
		})

	}
	if validate.Email(user.Email) != "" {
		Errors = append(Errors, validate.ValidationErr{
			FieldValue: "Email",
			ErrMassage: validate.Email(user.Email),
		})

	}

	if validate.Length(8, 64, user.Password) != "" {
		Errors = append(Errors, validate.ValidationErr{
			FieldValue: "Password",
			ErrMassage: validate.Length(8, 64, user.Password),
		})

	}

	if validate.UniqueEmail(user.Email) != "" {
		Errors = append(Errors, validate.ValidationErr{
			FieldValue: "Email2",
			ErrMassage: validate.UniqueEmail(user.Email),
		})

	}

	return Errors
}
