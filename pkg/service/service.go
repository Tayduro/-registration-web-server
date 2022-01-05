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
	if validate.Length(models.MinNameLength, models.MaxNameLength, user.Email) != "" {
		Errors = append(Errors, validate.ValidationErr{
			FieldValue: "Email",
			ErrMassage: validate.Length(models.MinNameLength, models.MaxNameLength, user.Email),
		})

	}

	if validate.Length(models.MinNameLength, models.MaxNameLength, user.Password) != "" {
		Errors = append(Errors, validate.ValidationErr{
			FieldValue: "Password",
			ErrMassage: validate.Length(models.MinNameLength, models.MaxNameLength, user.Password),
		})

	}
	//validate.Length(models.MinNameLength, models.MaxNameLength, user.FirstName)
	//validate.Length(models.MinNameLength, models.MaxNameLength, user.LastName)
	return Errors
}
