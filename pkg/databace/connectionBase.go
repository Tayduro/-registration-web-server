package databace

import (
	"fmt"
	"github.com/Tayduro/registration-web-server/pkg/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	Host = "localhost"
	Port = 6080
	User     = "postgres"
	Password =  "12345"
	Dbname   = "users"
)

func DataB(users * models.User) {
	currentUser := models.User{
		FirstName: users.FirstName,
		LastName: users.LastName,
		Email: users.Email,
		Password: users.Password,
	}

	connstring := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		Host, Port, Dbname, User, Password)
	//psqlconn := fmt.Sprintf(
	//	"postgres://%s:%s@%s:%d/%s?sslmode=disable",
	//	user, password, host, port, dbname )
	//fmt.Println(psqlconn)

	db, err := sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}

	//psqlconn := "postgres://postgres:12345@localhost:6080/users?sslmode=disable"
	//
	//db, err := sqlx.Open("postgres", psqlconn)
	//if err != nil{
	//	panic(err)
	//}

	defer db.Close()

	insert, err := db.Queryx("INSERT INTO users (first_name,last_name,email) VALUES($1, $2, $3)", &currentUser.FirstName, &currentUser.LastName, &currentUser.Email)
	if err != nil{
		panic(err)
	}
	defer insert.Close()
	//fmt.Println(insert)

	fmt.Println("connect to server...")
}
