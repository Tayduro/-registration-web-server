package service

import (
	"crypto/sha256"
	"github.com/Tayduro/registration-web-server/pkg/databace"
	"github.com/Tayduro/registration-web-server/pkg/models"
	"github.com/Tayduro/registration-web-server/pkg/validate"
	"github.com/jmoiron/sqlx"
	"math/rand"
	//_ "time"
	"fmt"
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

	//if validate.Length(8, 64, user.Password) != "" {
	//	Errors = append(Errors, validate.ValidationErr{
	//		FieldValue: "Password",
	//		ErrMassage: validate.Length(8, 64, user.Password),
	//	})
	//
	//}

	if validate.UniqueEmail(user.Email) != "" {
		Errors = append(Errors, validate.ValidationErr{
			FieldValue: "Email",
			ErrMassage: validate.UniqueEmail(user.Email),
		})

	}

	salt := RandStringRunes(5)
	fmt.Println(salt,"salt")
	fmt.Println(user.Password, "Password")
	newPassword := fmt.Sprintf("%s%s",salt,user.Password)
	fmt.Println(newPassword, "newPassword")
	hashBits := sha256.Sum256([]byte(newPassword))
	fmt.Println(hashBits, "hash")
	//fmt.Printf("%x", hash)
	hash := fmt.Sprintf("%x", hashBits)
	fmt.Println(hash, "hash")



	connstring := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		databace.Host, databace.Port, databace.Dbname, databace.User, databace.Password)

	db, err := sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}

	defer db.Close()



	insert, err := db.Queryx("INSERT INTO credentials (salt ,hash) VALUES($1, $2)", salt, hash)
	if err != nil{
		panic(err)
	}

	defer insert.Close()

	return Errors

}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}