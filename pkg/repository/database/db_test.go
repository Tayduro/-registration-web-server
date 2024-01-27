package database

import (
	"database/sql"
	"fmt"
	"github.com/Tayduro/registration-web-server/pkg/models"
	"testing"
)

func Test_DataBaseRegistration(t *testing.T) {

	testCase := models.User{
		Id:        "2",
		FirstName: "Bob",
		LastName:  "Nicolson",
		Email:     "cat2@gmail.com",
		Password:  "12345678",
	}
	id, err := svc.DataBaseRegistration(&testCase)
	if err != nil {
		t.Error("сan not insert data in database:", err)
	}

	type testUser struct {
		Id        string
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		Email     string `db:"email"`
	}

	u := testUser{}

	err = svc.db.Get(&u, "SELECT first_name, last_name, email FROM users WHERE user_id = $1", id)
	if err != nil {
		t.Error("can not get data from database", err)
	}

	if testCase.FirstName != u.FirstName || testCase.LastName != u.LastName || testCase.Email != u.Email {
		t.Error("user data does not match")
	}
}

func Test_GetEmailIfAvailable(t *testing.T) {
	testCase := models.User{
		Id:        "1",
		FirstName: "Bob",
		LastName:  "Nicolson",
		Email:     "cat1@gmail.com",
		Password:  "12345678",
	}
	_, err := svc.GetEmailIfAvailable(testCase.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			t.Error(sql.ErrNoRows)
		}
	}
}

type credentialsByDatabase struct {
	Id   string
	Salt string `db:"salt"`
	Hash string `db:"hash"`
}

type testCredentials struct {
	id   string
	salt string
	hash string
}

func Test_RecordingCredentials(t *testing.T) {
	testCase := testCredentials{
		id:   "2",
		salt: "M2FQhzj8WD",
		hash: "e7f84d27c02d72b80ea5bd4ca31e9a6dcf38145382928e9195c6304bd9e21715",
	}

	err := svc.InsertCredentials(testCase.id, testCase.salt, testCase.hash)
	if err != nil {
		t.Error("сan not insert data in database:", err)
	}

	c := credentialsByDatabase{}

	err = svc.db.Get(&c, "SELECT salt, hash FROM credentials WHERE user_id = $1", testCase.id)
	if err != nil {
		t.Error("can not get data from database", err)
	}

	if testCase.salt != c.Salt || testCase.hash != c.Hash {
		t.Error("salt or hash do not match")
	}
}

func Test_GetCredentialsByEmail(t *testing.T) {
	testCaseUser := models.User{
		Id:        "1",
		FirstName: "Bob",
		LastName:  "Nicolson",
		Email:     "cat1@gmail.com",
		Password:  "12345678",
	}

	testCaseCredentials := testCredentials{
		id:   "1",
		salt: "M2FQhzj8WD",
		hash: "e7f84d27c02d72b80ea5bd4ca31e9a6dcf38145382928e9195c6304bd9e21715",
	}

	c, err := svc.GetCredentialsByEmail(testCaseUser.Email)
	if err != nil {
		t.Error("can not get data from database", err)
	}

	if c.Id != testCaseCredentials.id || c.Salt != testCaseCredentials.salt || c.Hash != testCaseCredentials.hash {
		t.Error("id or salt or hash do not match")
	}

}

type testAccessToken struct {
	id    string
	token string
}

func Test_InsertToken(t *testing.T) {
	testCase := testAccessToken{
		id:    "2",
		token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxIiwiZXhwIjoxNjQ3ODk2OTE1fQ.rg-Z0tbs9Kn2dDsRAlaZ3qF1JMuSyaC8uER55TWO14k",
	}

	err := svc.InsertToken(testCase.id, testCase.token)
	if err != nil {
		t.Error("сan not insert data in database:", err)
	}
	var token string
	err = svc.db.QueryRowx("select token from access_token where token = $1", &testCase.token).Scan(&token)
	if err != nil {
		t.Error("can not get data from database", err)
	}
	if testCase.token != token {
		t.Error("tokens do not match")
	}

}

func Test_GetUserByToken(t *testing.T) {
	testCase := struct {
		id        string
		token     string
		FirstName string
		LastName  string
	}{
		id:        "1",
		FirstName: "Bob",
		LastName:  "Nicolson",
		token:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxIiwiZXhwIjoxNjQ3ODk2OTE1fQ.rg-Z0tbs9Kn2dDsRAlaZ3qF1JMuSyaC8uER55TWO14k",
	}

	p, err := svc.GetUserByToken(testCase.token)
	if err != nil {
		t.Error("can not get data from database", err)
	}

	if testCase.FirstName != p.FirstName || testCase.LastName != p.LastName {
		t.Error("first name or last name do not match")
	}
}

func Test_DeleteToken(t *testing.T) {
	testCase := testAccessToken{
		id:    "1",
		token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxIiwiZXhwIjoxNjQ3ODk2OTE1fQ.rg-Z0tbs9Kn2dDsRAlaZ3qF1JMuSyaC8uER55TWO14k",
	}

	err := svc.DeleteToken(testCase.token)
	if err != nil {
		t.Error("can not delete token from database", err)
	}
	var token string
	err = svc.db.QueryRowx("select token from access_token where token = $1", &testCase.token).Scan(&token)
	if err == nil {
		t.Error("can not get token from database", err)
	}
}

func truncateCustomTable(name string) {
	_, err := svc.db.Queryx(fmt.Sprintf("TRUNCATE TABLE %s CASCADE;", name))
	if err != nil {
		panic(err)
	}
}
