package database

import (
	"github.com/Tayduro/registration-web-server/pkg/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UsersRepository struct {
	db sqlx.DB
}

func NewUsersRepository(db sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (u *UsersRepository) DataBaseRegistration(users *models.User) (string, error) {
	currentUser := models.User{
		FirstName: users.FirstName,
		LastName:  users.LastName,
		Email:     users.Email,
		Password:  users.Password,
	}

	var userId string
	err := u.db.QueryRowx("INSERT INTO users (first_name,last_name,email) VALUES($1, $2, $3) RETURNING user_id", &currentUser.FirstName, &currentUser.LastName, &currentUser.Email).Scan(&userId)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func (u *UsersRepository) InsertCredentials(UserId string, salt string, hash string) error {
	insert, err := u.db.Queryx("INSERT INTO credentials (user_id,salt ,hash) VALUES($1, $2, $3)", UserId, salt, hash)
	if err != nil {
		return err
	}

	defer insert.Close()

	return nil
}

func (u *UsersRepository) InsertToken(dbUserId string, token string) error {
	insert, err := u.db.Queryx("INSERT INTO access_token (user_id,token) VALUES($1, $2)", dbUserId, token)
	if err != nil {
		return err
	}
	defer insert.Close()

	return nil
}

func (u *UsersRepository) DeleteToken(token string) error {

	drop, err := u.db.Queryx("delete from access_token where token = $1", token)
	if err != nil {
		return err
	}

	defer drop.Close()

	return nil
}

func (u *UsersRepository) GetEmailIfAvailable(email string) (string, error) {
	var dbEmail string

	err := u.db.QueryRowx("SELECT email FROM users WHERE email= $1", email).Scan(&dbEmail)
	if err != nil {
		return "", err
	}
	return dbEmail, nil
}

func (u *UsersRepository) GetUserByToken(token string) (models.Person, error) {
	p := models.Person{}

	err := u.db.Get(&p, "SELECT users.first_name, last_name FROM users left join access_token a on users.user_id = a.user_id where a.token = $1", token)
	if err != nil {
		return p, nil
	}
	return p, nil

}

func (u *UsersRepository) GetCredentialsByEmail(email string) (models.Credentials, error) {
	p := models.Credentials{}

	err := u.db.Get(&p, "SELECT credentials.user_id, salt, hash FROM credentials left join users a on credentials.user_id = a.user_id where a.email = $1", email)
	if err != nil {
		return p, nil
	}
	return p, nil
}
