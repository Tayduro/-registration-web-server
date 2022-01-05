package service

import (
	"github.com/Tayduro/registration-web-server/pkg/models"
	"github.com/Tayduro/registration-web-server/pkg/validate"
)

func Signup(user *models.User) error {
	err := validate.Email(user.Email)
	validate.Length(models.MinNameLength, models.MaxNameLength, user.Name)
	validate.Length(models.MinNameLength, models.MaxNameLength, user.LastName)
}
