package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/Tayduro/registration-web-server/pkg/models"
	"github.com/Tayduro/registration-web-server/pkg/validate"
	"math/rand"
	"strings"
	"time"
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

	if validate.Length(models.MinNameLength, models.MaxNameLength, user.Password) != "" {
		Errors = append(Errors, validate.ValidationErr{
			FieldValue: "Password",
			ErrMassage: validate.Length(8, 64, user.Password),
		})

	}
	if validate.Email(user.Email) != "" {
		Errors = append(Errors, validate.ValidationErr{
			FieldValue: "Email",
			ErrMassage: validate.Email(user.Email),
		})

	}

	if validate.UniqueEmail(user.Email) != "" {
		Errors = append(Errors, validate.ValidationErr{
			FieldValue: "Email",
			ErrMassage: validate.UniqueEmail(user.Email),
		})

	}
	return Errors

}

func CreatingHash(dbSalt string, userPassword string) string {
	password := dbSalt + userPassword
	hashBits := sha256.Sum256([]byte(password))
	hash := fmt.Sprintf("%x", hashBits)

	return hash
}

func RandomString() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	length := 10
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()

	return str
}



//func GettingUserData2(token string) []models.UserData {
//
//	userData := make([]models.UserData, 0, 0)
//
//	UserId := databace.GetUserId(token)
//	FirstName := databace.GetUserFirstName(UserId)
//	LastName := databace.GetUserLastName(UserId)
//	hmacSampleSecret := []byte(config.GetKey())
//
//	if jwtToken.ParseHmac(token, hmacSampleSecret) == nil {
//
//
//		userData = append(userData, models.UserData{
//			Field:      "FirstName",
//			FieldValue: FirstName,
//		})
//		userData = append(userData, models.UserData{
//			Field:      "LastName",
//			FieldValue: LastName,
//		})
//		return userData
//	}
//
//	return userData
//}
