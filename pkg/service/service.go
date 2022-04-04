package service

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	jwtToken "github.com/Tayduro/registration-web-server/pkg/jwt"
	"github.com/Tayduro/registration-web-server/pkg/models"
	"github.com/Tayduro/registration-web-server/pkg/repository/database"
	"github.com/Tayduro/registration-web-server/pkg/validate"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"strings"
	"time"
)

func (ss *SignupService) CheckError(user *models.User) ([]validate.ValidationErr, error) {
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
		email, err := database.NewUsersRepository(*ss.db).GetEmailIfAvailable(user.Email)
		if err != nil {
			if err != sql.ErrNoRows{
				return nil, err
			}
		}

		if email == user.Email {
			Errors = append(Errors, validate.ValidationErr{
				FieldValue: "Email",
				ErrMassage: "this email is already in use",
			})
		}


	return Errors, nil

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

type SignupService struct {
	db  *sqlx.DB
	key string
}

func NewSignupService(db *sqlx.DB, key string ) *SignupService {
	return &SignupService{
		db: db,
		key: key,
	}
}


func (ss *SignupService)SignUp(u *models.User) error {
	salt := RandomString()
	hash := CreatingHash(salt, u.Password)

    userId ,err := database.NewUsersRepository(*ss.db).DataBaseRegistration(u)
	if err != nil {
		return err
	}

	err = database.NewUsersRepository(*ss.db).InsertCredentials(userId, salt, hash)
	return nil

}

func (ss *SignupService)SignIn(u *models.User) (string, error) {
	c, err :=database.NewUsersRepository(*ss.db).GetCredentialsByEmail(u.Email)
	if err != nil {
		return "", err
	}
	hash := CreatingHash(c.Salt, u.Password)


	var token string

	if c.Hash == hash{

		var hmacSampleSecret = []byte(ss.key)

		token = jwtToken.NewJWT(c.Id, hmacSampleSecret)

		err = database.NewUsersRepository(*ss.db).InsertToken(c.Id, token)
		if err != nil {
			return "", err
		}


	}

	return token, nil
}

func (ss *SignupService)GettingUserInformationHandler( token string) ([]models.UserData, error) {

	var hmacSampleSecret = []byte(ss.key)

	 err := jwtToken.ParseHmac(token, hmacSampleSecret)
	if err != nil {
		return nil, err
	}

	userData := make([]models.UserData, 0, 0)
	if len(token) == 0 {
		return userData, nil
	}
	p, err := database.NewUsersRepository(*ss.db).GetUserByToken(token)
	if err != nil {
		return nil, err
	}

	userData = append(userData, models.UserData{
		Field:      "FirstName",
		FieldValue: p.FirstName,
	})
	userData = append(userData, models.UserData{
		Field:      "LastName",
		FieldValue: p.LastName,
	})

	return userData, nil

}

func (ss *SignupService)DeleteToken( token string)  error {
	err := database.NewUsersRepository(*ss.db).DeleteToken(token)
	if err != nil {
		return err
	}
	return nil
}



