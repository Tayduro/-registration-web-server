package service

import (
<<<<<<< HEAD
=======
	"crypto/sha256"
	"fmt"
	"github.com/jmoiron/sqlx"
	//"gopkg.in/yaml.v2"
	//"io/ioutil"
	//"log"

	"github.com/Tayduro/registration-web-server/pkg/config"
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
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
<<<<<<< HEAD

}
=======
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3

}

func Login(user *models.User) string {

<<<<<<< HEAD
=======
	connstring :=config.ConfigServer()
>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3




<<<<<<< HEAD






=======
	var dbUserId string
	err = db.QueryRow("SELECT user_id FROM users WHERE email= $1", user.Email).Scan(&dbUserId)
	if err != nil{
		panic(err)
	}
	var dbSalt string
	err = db.QueryRow("SELECT salt FROM credentials WHERE user_id= $1", dbUserId).Scan(&dbSalt)
	if err != nil{
		panic(err)
	}
	password := dbSalt + user.Password
	hashBits := sha256.Sum256([]byte(password))
	hash := fmt.Sprintf("%x", hashBits)
	var dbHash string
	err = db.QueryRow("SELECT hash FROM credentials WHERE user_id= $1", dbUserId).Scan(&dbHash)
	if err != nil{
		panic(err)
	}
	var dbToken string
	if dbHash == hash {

		err = db.QueryRow("SELECT token FROM access_token WHERE user_id= $1", dbUserId).Scan(&dbToken)
		if err != nil{
			panic(err)
		}

	}
	return dbToken
}







>>>>>>> bcf1b17145918f375e87bb4df57158a9ea8e60a3
