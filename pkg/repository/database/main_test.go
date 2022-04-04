package database

import (
	"github.com/Tayduro/registration-web-server/pkg/models"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
)

var svc *UsersRepository

func TestMain(m *testing.M) {
	conn, err := sqlx.Connect("postgres", "postgres://postgres:12345@localhost:6080/users?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	svc = NewUsersRepository(*conn)

	testCaseUser := models.User{
		Id: "1",
		FirstName: "Bob",
		LastName: "Nicolson",
		Email: "cat1@gmail.com",
		Password: "12345678",
	}

	testCaseCredentials := testCredentials {
		id: "1",
		salt: "M2FQhzj8WD",
		hash: "e7f84d27c02d72b80ea5bd4ca31e9a6dcf38145382928e9195c6304bd9e21715",
	}

	testCaseAccessToken := testAccessToken{
		id:    "1",
		token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxIiwiZXhwIjoxNjQ3ODk2OTE1fQ.rg-Z0tbs9Kn2dDsRAlaZ3qF1JMuSyaC8uER55TWO14k",
	}


	insertUser, err := svc.db.Queryx("INSERT INTO users (first_name,last_name,email) VALUES($1, $2, $3)", &testCaseUser.FirstName, &testCaseUser.LastName, &testCaseUser.Email)
	if err != nil {
		log.Fatal("cannot save data in database:", err)
	}
	defer insertUser.Close()


	insertCredentials, err := svc.db.Queryx("INSERT INTO credentials (user_id,salt ,hash) VALUES($1, $2, $3)", testCaseCredentials.id, testCaseCredentials.salt, testCaseCredentials.hash)
	if err != nil {
		log.Fatal("cannot save data in database:", err)
	}

	defer insertCredentials.Close()

	insertAccessToken, err := svc.db.Queryx("INSERT INTO access_token (user_id,token) VALUES($1, $2)", testCaseAccessToken.id, testCaseAccessToken.token)
	if err != nil {
		log.Fatal("cannot save data in database:", err)
	}
	defer insertAccessToken.Close()

	code := m.Run()
	truncateCustomTable("users")
	truncateCustomTable("credentials")
	truncateCustomTable("access_token")
	os.Exit(code)
}
